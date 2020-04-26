package etcdv3

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"

	discovery "github.com/liangjfblue/gdiscovery"

	hash "github.com/mitchellh/hashstructure"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

type etcdDiscovery struct {
	client  *clientv3.Client
	options discovery.Options
	sync.RWMutex
	srvNodes  map[string][]discovery.Service
	registers map[string]uint64
	leases    map[string]clientv3.LeaseID
}

func newEtcdDiscovery(option ...discovery.Option) *etcdDiscovery {
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
		config.Endpoints = append(config.Endpoints, e.options.Addrs...)
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

func (e *etcdDiscovery) Register(ctx context.Context, service *discovery.Service, opts ...discovery.RegisterOption) error {
	if len(service.Nodes) == 0 {
		return errors.New("nodes is empty")
	}

	var err error
	for _, node := range service.Nodes {
		go func(err error, node *discovery.Node) {
			err = e.keepAliveV2(ctx, service, node, opts...)
		}(err, node)
	}
	return err
}

func (e *etcdDiscovery) keepAlive(ctx context.Context, service *discovery.Service, node *discovery.Node, opts ...discovery.RegisterOption) error {
	var (
		err error
		ro  discovery.RegisterOptions
		lgr *clientv3.LeaseGrantResponse
	)

	e.RLock()
	leaseID, ok := e.leases[service.SrvName+node.Id]
	e.RUnlock()

	//若本节点的服务已注册过,更新节点信息
	if !ok {
		ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
		defer cancel()

		rsp, err := e.client.Get(ctx, discovery.NodePath(service.SrvName, node.Id), clientv3.WithSerializable())
		if err != nil {
			return err
		}

		for _, kv := range rsp.Kvs {
			if kv.Lease > 0 {
				leaseID = clientv3.LeaseID(kv.Lease)

				srv := discovery.Decode(kv.Value)

				//if no node then continue
				if srv == nil || len(srv.Nodes) <= 0 {
					continue
				}

				h, err := hash.Hash(srv.Nodes[0], nil)
				if err != nil {
					continue
				}

				e.Lock()
				e.registers[service.SrvName+node.Id] = h
				e.leases[service.SrvName+node.Id] = leaseID
				e.Unlock()
			}
		}
	}

	leaseNotFound := false

	if leaseID > 0 {
		//契约已存在,更新契约
		if _, err := e.client.KeepAliveOnce(ctx, leaseID); err != nil {
			if err != rpctypes.ErrLeaseNotFound {
				return err
			}

			leaseNotFound = true
		}
	}

	//判断节点服务是否有变化
	h, err := hash.Hash(node, nil)
	if err != nil {
		return err
	}

	e.RLock()
	v, ok := e.registers[service.SrvName+node.Id]
	e.RUnlock()

	if ok && h == v && !leaseNotFound {
		log.Fatal(fmt.Sprintf("Service %s node %s unchanged skipping registration", service.SrvName, node.Id))
		return nil
	}

	srv := &discovery.Service{
		SrvName:   service.SrvName,
		Version:   service.Version,
		Metadata:  service.Metadata,
		Endpoints: service.Endpoints,
		Nodes:     []*discovery.Node{node},
	}

	for _, o := range opts {
		o(&ro)
	}

	ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
	defer cancel()

	if ro.TTL.Seconds() > 0 {
		lgr, err = e.client.Grant(ctx, int64(ro.TTL.Seconds()))
		if err != nil {
			return err
		}
	}

	if lgr != nil {
		_, err = e.client.Put(ctx,
			discovery.NodePath(service.SrvName, node.Id), discovery.Encode(srv), clientv3.WithLease(leaseID))
	} else {
		_, err = e.client.Put(ctx, discovery.NodePath(service.SrvName, node.Id), discovery.Encode(srv))
	}
	if err != nil {
		return err
	}

	e.Lock()
	e.registers[service.SrvName+node.Id] = h
	if lgr.ID > 0 {
		e.leases[service.SrvName+node.Id] = lgr.ID
	}
	e.Unlock()

	return nil
}

func (e *etcdDiscovery) keepAliveV2(ctx context.Context, service *discovery.Service, node *discovery.Node, opts ...discovery.RegisterOption) error {
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

	ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
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

	ticker := time.NewTicker(ro.TTL / 2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
			if _, err := e.client.KeepAlive(context.TODO(), lgr.ID); err != nil {
				return err
			}
		case <-ctx.Done():
			return errors.New("ctx done")
		}
	}
}

func (e *etcdDiscovery) Deregister(ctx context.Context, service *discovery.Service) error {
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

func (e *etcdDiscovery) Watch(ctx context.Context, opts ...discovery.WatchOption) {

}

func (e *etcdDiscovery) Get(context.Context, string) ([]*discovery.Service, error) {

	return nil, nil
}

func (e *etcdDiscovery) All(context.Context, string) ([]*discovery.Service, error) {
	return nil, nil
}

func (e *etcdDiscovery) putEvent(ev *clientv3.Event) {
	e.add(discovery.Decode(ev.Kv.Value))
}

func (e *etcdDiscovery) delEvent(ev *clientv3.Event) {
	e.del(discovery.Decode(ev.Kv.Value))
}

func (e *etcdDiscovery) add(service *discovery.Service) {
	e.Lock()
	defer e.Unlock()
	e.srvNodes[service.SrvName+service.Version] = append(e.srvNodes[service.SrvName+service.Version], *service)
}

func (e *etcdDiscovery) del(service *discovery.Service) {
	e.Lock()
	defer e.Unlock()
	if _, ok := e.srvNodes[service.SrvName+service.Version]; ok {
		delete(e.srvNodes, service.SrvName+service.Version)
	}
}
