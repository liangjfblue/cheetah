package discovery

import (
	"context"
	"errors"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Etcd struct {
	Client *clientv3.Client

	endpoints   []string
	dialTimeout time.Duration
}

func NewEtcd(client *clientv3.Client) *Etcd {
	return &Etcd{
		Client: client,
	}
}

func (e *Etcd) Put(ctx context.Context, key string, val string, ops ...clientv3.OpOption) error {
	if _, err := e.Client.Put(ctx, key, val, ops...); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (e *Etcd) Get(ctx context.Context, key string, ops ...clientv3.OpOption) error {
	if _, err := e.Client.Get(ctx, key, ops...); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (e *Etcd) Watch(ctx context.Context, key string, ops ...clientv3.OpOption) clientv3.WatchChan {
	return e.Client.Watch(ctx, key, ops...)
}
