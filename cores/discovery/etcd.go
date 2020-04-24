package discovery

import (
	"errors"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Etcd struct {
	Client *clientv3.Client

	endpoints   []string
	dialTimeout time.Duration
}

func NewEtcd(endpoints []string, dialTimeout time.Duration) *Etcd {
	return &Etcd{
		endpoints:   endpoints,
		dialTimeout: dialTimeout,
	}
}

func (e *Etcd) InitEtcd() error {
	var (
		err error
	)
	e.Client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.endpoints,
		DialTimeout: time.Second * e.dialTimeout,
	})
	if err != nil {
		return errors.Unwrap(err)
	}

	return nil
}
