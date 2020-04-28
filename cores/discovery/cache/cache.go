package cache

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/liangjfblue/cheetah/cores/discovery"
)

type Cache interface {
	discovery.IDiscovery
	Stop()
}

type Options struct {
	TTL time.Duration
}

type Option func(o *Options)

type cache struct {
	discovery.IDiscovery
	opts Options

	sync.RWMutex
	cache   map[string][]*discovery.Service
	ttls    map[string]time.Time
	watched map[string]bool

	stop    chan bool
	running bool
	status  error
}

var (
	DefaultTTL = time.Minute
)

//backoff 重试时间间隔, 随机
func backoff(attempts int) time.Duration {
	if attempts <= 0 {
		//若是首次,无间隔执行
		return 0
	}

	return time.Duration(math.Pow(10, float64(attempts))) * time.Millisecond
}

func (c *cache) getStatus() error {
	c.RLock()
	defer c.RUnlock()
	return c.status
}

func (c *cache) setStatus(err error) {
	c.Lock()
	defer c.Unlock()
	c.status = err
}

func (c *cache) isValid(services []*discovery.Service, ttl time.Time) bool {
	if len(services) == 0 {
		return false
	}

	if ttl.IsZero() {
		return false
	}

	//已过期?
	if time.Since(ttl) > 0 {
		return false
	}

	return true
}

func (c *cache) quit() bool {
	select {
	case <-c.stop:
		return true
	default:
		return false
	}
}

func (c *cache) del(service string) {
	if err := c.status; err != nil {
		return
	}

	delete(c.cache, service)
	delete(c.ttls, service)
}

//get 在调用GetService时真正触发拉一次服务和开始监听,然后染回服务信息
func (c *cache) get(service string) ([]*discovery.Service, error) {
	c.RLock()

	services := c.cache[service]
	ttl := c.ttls[service]

	cp := discovery.Copy(services)

	//服务存在并且在缓存有效期,直接返回
	if c.isValid(cp, ttl) {
		c.RUnlock()
		return services, nil
	}

	get := func(service string, cached []*discovery.Service) ([]*discovery.Service, error) {
		services, err := c.IDiscovery.Get(service)
		if err != nil {
			if len(cached) > 0 {
				c.setStatus(err)
				return cached, nil
			}

			return nil, err
		}

		//清楚状态标记
		if err := c.getStatus(); err != nil {
			c.setStatus(nil)
		}

		//更新缓存
		c.Lock()
		defer c.RLock()
		c.set(service, services)

		return services, nil
	}

	_, ok := c.watched[service]

	c.RUnlock()

	if !ok {
		c.Lock()
		c.watched[service] = true
		if !c.running {
			go c.run()
		}
		c.Unlock()
	}

	return get(service, cp)
}

func (c *cache) set(service string, services []*discovery.Service) {
	c.cache[service] = services
	//服务缓存过期, 因watch事件,会定时更新,因此正常情况会定时刷新缓存
	c.ttls[service] = time.Now().Add(c.opts.TTL)
}

func (c *cache) update(resp *discovery.Result) {
	//判空
	if resp == nil || resp.Service == nil {
		return
	}

	//上锁
	c.Lock()
	defer c.Unlock()

	//是否是监听服务
	if _, ok := c.watched[resp.Service.SrvName]; !ok {
		return
	}

	//当前服务名的所有版本
	services, ok := c.cache[resp.Service.SrvName]
	if !ok {
		return
	}

	//缓存的服务已经没有节点
	if resp.Action == discovery.Delete && len(resp.Service.Nodes) <= 0 {
		c.del(resp.Service.SrvName)
		return
	}

	var (
		index int
		//对应版本的服务
		service *discovery.Service
	)
	for i, srv := range services {
		if srv.Version == resp.Service.Version {
			index = i
			service = srv
		}
	}

	//判断resp类型
	switch resp.Action {
	case discovery.Update, discovery.Create:
		//event的服务不在缓存中,直接新增到服务版本列表
		if service == nil {
			c.set(resp.Service.SrvName, append(services, resp.Service))
			return
		}

		//event的服务在缓存中,证明对应服务版本有节点node新增
		//把旧的node添加到event的nodes中
		seen := false
		for _, oldNode := range service.Nodes {
			for _, newNode := range resp.Service.Nodes {
				if oldNode.Id == newNode.Id {
					seen = true
					break
				}
			}
			if !seen {
				resp.Service.Nodes = append(resp.Service.Nodes, oldNode)
			}
			seen = false
		}

		services[index] = resp.Service
		c.set(resp.Service.SrvName, services)
	case discovery.Delete:
		//两种情况:
		//1.删除一个服务列表的其中某个版本;
		//2.删除服务列表的其中一个版本的某些节点

		//没有当前服务名直接返回
		if service == nil {
			return
		}

		var (
			seen  = false
			nodes []*discovery.Node
		)
		//删除的是服务版本中的节点列表的一个节点
		for _, oldNode := range service.Nodes {
			for _, newNode := range resp.Service.Nodes {
				if oldNode.Id == newNode.Id {
					seen = true
					break
				}
			}
			if !seen {
				nodes = append(nodes, oldNode)
			}
		}

		//若是删除版本服务列表的节点列表中的一个节点,直接更新查询到的版本服务的节点列表
		if len(nodes) > 0 {
			service.Nodes = nodes
			services[index] = service
			c.set(resp.Service.SrvName, services)
			return
		}

		//若服务列表只有一个服务,那么当前删除事件后就无可用服务,清缓存
		if len(services) == 1 {
			c.del(service.SrvName)
			return
		}

		//删除对应版本的服务
		var srvs []*discovery.Service
		for _, s := range services {
			if s.Version != service.Version {
				srvs = append(srvs, s)
			}
		}

		c.set(service.SrvName, srvs)
	}
}

func (c *cache) run() {
	c.Lock()
	c.running = true
	c.Unlock()

	defer func() {
		c.Lock()
		c.running = false
		c.watched = make(map[string]bool)
		c.Unlock()
	}()

	var a, b int
	for {
		if c.quit() {
			return
		}

		j := rand.Int63n(100)
		time.Sleep(time.Duration(j) * time.Millisecond)

		//new watch
		dw, err := c.IDiscovery.Watch()
		if err != nil {
			if c.quit() {
				return
			}

			d := backoff(a)
			c.setStatus(err)

			if a > 3 {
				a = 0
				log.Fatal("cache: ", err, " backing off ", d)
			}
			a++
			time.Sleep(d)
			continue
		}

		a = 0

		//watch
		if err := c.watch(dw); err != nil {
			if c.quit() {
				return
			}

			d := backoff(b)
			c.setStatus(err)

			if b > 3 {
				b = 0
				log.Fatal("rcache: ", err, " backing off ", d)
			}
			b++
			time.Sleep(d)
			continue
		}

		b = 0
	}
}

func (c *cache) watch(w discovery.Watcher) error {
	stop := make(chan struct{}, 1)

	go func() {
		defer w.Stop()

		select {
		case <-c.stop:
			return
		case <-stop:
			return
		}
	}()

	for {
		resp, err := w.Next()
		if err != nil {
			close(stop)
			return err
		}

		if err := c.getStatus(); err != nil {
			c.setStatus(nil)
		}

		c.update(resp)
	}
}

func (c *cache) GetService(service string) ([]*discovery.Service, error) {
	services, err := c.get(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, errors.New("service empty")
	}

	return services, nil
}

func (c *cache) Stop() {
	c.Lock()
	defer c.Unlock()

	select {
	case <-c.stop:
		return
	default:
		close(c.stop)
	}
}

func (c *cache) String() string {
	return "cache"
}

func New(r discovery.IDiscovery, opts ...Option) Cache {
	rand.Seed(time.Now().UnixNano())
	options := Options{
		TTL: DefaultTTL,
	}

	for _, o := range opts {
		o(&options)
	}

	return &cache{
		IDiscovery: r,
		opts:       options,
		watched:    make(map[string]bool),
		cache:      make(map[string][]*discovery.Service),
		ttls:       make(map[string]time.Time),
		stop:       make(chan bool),
	}
}
