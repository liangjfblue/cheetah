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

type Node struct {
	etcd      *Etcd
	lease     clientv3.Lease
	leaseId   clientv3.LeaseID
	stopChan  chan struct{}
	leaseTime int64
	NodeInfo  *NodeInfo
}

//NewNode new node
func NewNode(etcd *Etcd, leaseTime int64, nodeInfo *NodeInfo) *Node {
	n := new(Node)
	n.etcd = etcd
	n.stopChan = make(chan struct{}, 1)
	n.leaseTime = leaseTime
	n.NodeInfo = nodeInfo
	return n
}

func (n *Node) Register() error {
	if n.leaseTime <= 0 {
		return errors.Unwrap(errors.New("lease time must greater zero"))
	}

	n.lease = clientv3.NewLease(n.etcd.Client)

	grantResp, err := n.lease.Grant(context.TODO(), n.leaseTime)
	if err != nil {
		return errors.Unwrap(err)
	}
	n.leaseId = grantResp.ID

	kv := clientv3.NewKV(n.etcd.Client)

	key := fmt.Sprintf("%s/%s", n.NodeInfo.Path, n.NodeInfo.SrvName)

	value, err := json.Marshal(n.NodeInfo)
	if err != nil {
		return errors.Unwrap(err)
	}

	log.Println(value)
	if _, err = kv.Put(context.TODO(), key, string(value), clientv3.WithLease(grantResp.ID)); err != nil {
		return errors.Unwrap(err)
	}
	return nil
}

//Heartbeat renew key for keep self to service discovery
func (n *Node) Heartbeat(ctx context.Context) {
	var (
		ticker *time.Ticker
	)

	if n.leaseTime == 1 {
		// if leaseTime less then 1 second, so set 500 millisecond heartbeat
		ticker = time.NewTicker(time.Duration(500) * time.Millisecond)
	} else {
		ticker = time.NewTicker(time.Duration(n.leaseTime-1) * time.Second)
	}

	for {
		select {
		case <-ticker.C:
			if _, err := n.lease.KeepAlive(context.TODO(), n.leaseId); err != nil {
				log.Fatal(err)
				return
			}
		case <-n.stopChan:
			log.Println("revoke lease, stop heartbeat")
			return
		case <-ctx.Done():
			log.Println("context done")
			return
		}
	}
}

//Revoke revoke the leaseId
func (n *Node) Revoke(ctx context.Context) error {
	if _, err := n.etcd.Client.Revoke(ctx, n.leaseId); err != nil {
		return errors.Unwrap(err)
	}

	n.stopChan <- struct{}{}

	return nil
}
