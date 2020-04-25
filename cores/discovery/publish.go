package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Publish struct {
	etcd      *Etcd
	lease     clientv3.Lease
	leaseId   clientv3.LeaseID
	leaseTime int64
}

//NewPublish new node
func NewPublish(etcd *Etcd, leaseTime int64) *Publish {
	n := new(Publish)
	n.etcd = etcd
	n.leaseTime = leaseTime
	return n
}

func (n *Publish) Register(ctx context.Context, nodeInfo NodeInfo) error {
	if n.leaseTime <= 0 {
		return errors.Unwrap(errors.New("lease time must greater zero"))
	}

	n.lease = clientv3.NewLease(n.etcd.Client)

	grantResp, err := n.lease.Grant(context.TODO(), n.leaseTime)
	if err != nil {
		return errors.Unwrap(err)
	}
	n.leaseId = grantResp.ID

	key := fmt.Sprintf("%s/%s", nodeInfo.Path, nodeInfo.SrvName)

	value, err := json.Marshal(nodeInfo)
	if err != nil {
		return errors.Unwrap(err)
	}

	if err = n.etcd.Put(ctx, key, string(value), clientv3.WithLease(grantResp.ID)); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

//Heartbeat renew key for keep self to service discovery
func (n *Publish) Heartbeat(ctx context.Context) {
	var (
		ticker *time.Ticker
	)

	if n.leaseTime == 1 {
		// if leaseTime less then 1 second, so set 500 millisecond heartbeat
		ticker = time.NewTicker(time.Duration(500) * time.Millisecond)
	} else {
		ticker = time.NewTicker(time.Duration(n.leaseTime) / 2 * time.Second)
	}

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
			if _, err := n.lease.KeepAlive(context.TODO(), n.leaseId); err != nil {
				log.Fatal(err)
				return
			}
		case <-ctx.Done():
			log.Println("context done")
			return
		}
	}
}

//Revoke revoke the leaseId
func (n *Publish) Revoke(ctx context.Context) error {
	if _, err := n.etcd.Client.Revoke(ctx, n.leaseId); err != nil {
		return errors.Unwrap(err)
	}
	return nil
}
