package etcdv3

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/coreos/etcd/clientv3"

	"github.com/liangjfblue/cheetah/cores/discovery"
)

type etcdWatcher struct {
	w       clientv3.WatchChan
	client  *clientv3.Client
	timeout time.Duration
	stop    chan struct{}
}

func newEtcdWatcher(d *etcdDiscovery, timeout time.Duration, opts ...discovery.WatchOption) (discovery.Watcher, error) {
	var wp discovery.WatchOptions
	for _, o := range opts {
		o(&wp)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	stop := make(chan bool, 1)

	go func() {
		<-stop
		cancel()
	}()

	watchPath := discovery.ServicePrefixPath()
	if len(wp.Service) > 0 {
		watchPath = discovery.ServicePath(wp.Service) + "/"
	}

	return &etcdWatcher{
		w:       d.client.Watch(ctx, watchPath, clientv3.WithPrefix(), clientv3.WithPrevKV()),
		client:  d.client,
		timeout: timeout,
		stop:    make(chan struct{}, 1),
	}, nil
}

func (e *etcdWatcher) Next() (*discovery.Result, error) {
	for v := range e.w {
		if v.Err() != nil {
			return nil, v.Err()
		}
		for _, event := range v.Events {
			var action discovery.EventType
			service := discovery.Decode(event.Kv.Value)

			switch event.Type {
			case clientv3.EventTypePut:
				if event.IsCreate() {
					action = discovery.Create
				} else if event.IsModify() {
					action = discovery.Update
				}
			case clientv3.EventTypeDelete:
				action = discovery.Delete
				service = discovery.Decode(event.PrevKv.Value)
			}

			if service == nil {
				continue
			}

			return &discovery.Result{
				Action:  action,
				Service: service,
			}, nil
		}
	}
	return nil, errors.New("could not get next")
}

func (e *etcdWatcher) Stop() {
	select {
	case <-e.stop:
		return
	default:
		close(e.stop)
	}
}
