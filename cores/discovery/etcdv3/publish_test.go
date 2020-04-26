package etcdv3

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func TestNode_Heartbeat(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.112:9002", "192.168.0.112:9004", "192.168.0.112:9006"},
		DialTimeout: time.Second * time.Duration(3),
	})
	if err != nil {
		t.Fatal(err)
	}

	e := NewEtcd(client)

	n := NewPublish(e, 2)

	if err := n.Register(context.TODO(), discovery.NodeInfo{
		Path:     "discovery",
		Env:      "dev",
		SrvName:  fmt.Sprintf("%s-%d", "web-user", 1),
		Addr:     "192.168.0.112",
		Hostname: fmt.Sprintf("%s-%d", "node", 1),
		Status:   1,
		Color:    "all",
	}); err != nil {
		t.Fatal(err)
	}

	go n.Heartbeat(context.Background())

	//time.Sleep(time.Second * 5)
	//if err := n.Revoke(context.Background()); err != nil {
	//	t.Fatal(err)
	//}

	time.Sleep(30 * time.Second)
}
