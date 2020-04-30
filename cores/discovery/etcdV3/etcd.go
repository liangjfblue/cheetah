package etcdV3

import (
	"context"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/liangjfblue/cheetah/cores/discovery"

	"github.com/pkg/errors"

	hash "github.com/mitchellh/hashstructure"
	"go.etcd.io/etcd/clientv3"
)

type etcdDiscovery struct {
	client  *clientv3.Client
	options discovery.Options
	sync.RWMutex
	srvNodes  map[string][]discovery.Service
	registers map[string]uint64
	leases    map[string]clientv3.LeaseID
}

func NewDiscovery(option ...discovery.Option) *etcdDiscovery {
	e := &etcdDiscovery{
		options:   discovery.Options{},
		srvNodes:  make(map[string][]discovery.Service, 0),
		registers: make(map[string]uint64, 0),
		leases:    make(map[string]clientv3.LeaseID, 0),
	}

	if err := e.configure(option...); err != nil {
		log.Fatal(err)
	}

	return e
}

func (e *etcdDiscovery) configure(opts ...discovery.Option) error {
	for _, o := range opts {
		o(&e.options)
	}

	if e.options.Timeout == 0 {
		e.options.Timeout = time.Second * 5
	}

	config := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}
	if e.options.Addrs != nil {
		config.Endpoints = e.options.Addrs
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return err
	}
	e.client = cli

	return nil
}

func (e *etcdDiscovery) Init(opts ...discovery.Option) error {
	return e.configure(opts...)
}

func (e *etcdDiscovery) Options() discovery.Options {
	return e.options
}

func (e *etcdDiscovery) Register(service *discovery.Service, opts ...discovery.RegisterOption) error {
	if len(service.Nodes) == 0 {
		return errors.New("discovery nodes is empty")
	}

	var err error
	for _, node := range service.Nodes {
		go func(err error, node *discovery.Node) {
			err = e.keepAlive(service, node, opts...)
		}(err, node)
	}
	return err
}

func (e *etcdDiscovery) keepAlive(service *discovery.Service, node *discovery.Node, opts ...discovery.RegisterOption) error {
	var (
		err error
		ro  discovery.RegisterOptions
		lgr *clientv3.LeaseGrantResponse
	)

	for _, o := range opts {
		o(&ro)
	}
	if ro.TTL.Seconds() <= 0 {
		ro.TTL = time.Second * 5
	}

	ctx, cancel := context.WithTimeout(context.TODO(), e.options.Timeout)
	defer cancel()

	if ro.TTL.Seconds() > 0 {
		lgr, err = e.client.Grant(ctx, int64(ro.TTL.Seconds()))
		if err != nil {
			return err
		}
	}

	h, err := hash.Hash(node, nil)
	if err != nil {
		return err
	}

	e.Lock()
	e.registers[service.SrvName+node.Id] = h
	e.leases[service.SrvName+node.Id] = lgr.ID
	e.Unlock()

	if _, err := e.client.Put(ctx,
		discovery.NodePath(service.SrvName, node.Id),
		discovery.Encode(service),
		clientv3.WithLease(lgr.ID)); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Duration(int(ro.TTL.Seconds())/2) * time.Second)
	for {
		select {
		case <-ticker.C:
			if _, err := e.client.KeepAlive(context.TODO(), lgr.ID); err != nil {
				return err
			}
		}
	}
}

func (e *etcdDiscovery) Deregister(service *discovery.Service) error {
	ctx, cancel := context.WithTimeout(context.TODO(), e.options.Timeout)
	defer cancel()

	if len(service.Nodes) == 0 {
		return errors.New("Require at least one node")
	}

	for _, node := range service.Nodes {
		e.Lock()
		leaseId, ok := e.leases[service.SrvName+node.Id]
		delete(e.leases, service.SrvName+node.Id)
		delete(e.registers, service.SrvName+node.Id)
		e.Unlock()

		if ok {
			if _, err := e.client.Revoke(ctx, leaseId); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *etcdDiscovery) GetService(srvName string) ([]*discovery.Service, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), e.options.Timeout)
	defer cancel()

	resp, err := e.client.Get(ctx, discovery.ServicePath(srvName)+"/", clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) <= 0 {
		return nil, errors.New("service not found node")
	}

	srvMap := make(map[string]*discovery.Service)

	for _, kv := range resp.Kvs {
		if srvNode := discovery.Decode(kv.Value); srvNode != nil {
			s, ok := srvMap[srvNode.Version]
			if !ok {
				s = &discovery.Service{
					SrvName:   srvNode.SrvName,
					Version:   srvNode.Version,
					Metadata:  srvNode.Metadata,
					Endpoints: srvNode.Endpoints,
					Nodes:     srvNode.Nodes,
				}
				srvMap[srvNode.Version] = s
			}

			for _, node := range srvNode.Nodes {
				s.Nodes = append(s.Nodes, node)
			}
		}
	}

	srvList := make([]*discovery.Service, 0, len(srvMap))
	for _, srv := range srvMap {
		srvList = append(srvList, srv)
	}

	return srvList, nil
}

func (e *etcdDiscovery) ListServices() ([]*discovery.Service, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), e.options.Timeout)
	defer cancel()

	resp, err := e.client.Get(ctx, discovery.ServicePrefixPath()+"/", clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) <= 0 {
		return nil, errors.New("service not found node")
	}

	versions := make(map[string]*discovery.Service)

	for _, kv := range resp.Kvs {
		if srvNode := discovery.Decode(kv.Value); srvNode != nil {
			//去重
			v, ok := versions[srvNode.SrvName+srvNode.Version]
			if !ok {
				versions[srvNode.SrvName+srvNode.Version] = srvNode
				continue
			}

			v.Nodes = append(v.Nodes, srvNode.Nodes...)
		}
	}

	srvList := make([]*discovery.Service, 0, len(versions))
	for _, srv := range versions {
		srvList = append(srvList, srv)
	}

	sort.Slice(srvList, func(i, j int) bool { return srvList[i].SrvName < srvList[j].SrvName })

	return srvList, nil
}

func (e *etcdDiscovery) Watch(opts ...discovery.WatchOption) (discovery.Watcher, error) {
	return newEtcdWatcher(e, e.options.Timeout, opts...)
}

func (e *etcdDiscovery) String() string {
	return "etcd discovery"
}
