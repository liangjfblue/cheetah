package discovery

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNode_Heartbeat(t *testing.T) {
	e := NewEtcd([]string{"192.168.0.112:9001", "192.168.0.112:9003", "192.168.0.112:9005"}, time.Duration(3))
	if err := e.InitEtcd(); err != nil {
		t.Fatal(err)
	}

	//for i := 0; i < 5; i++ {
	//	index := i % 2
	//	n := NewNode(e, 2, &NodeInfo{
	//		Path:     "discovery",
	//		Env:      "dev",
	//		SrvName:  fmt.Sprintf("%s-%d", "web-service", index),
	//		Addr:     "172.16.7.16",
	//		Hostname: fmt.Sprintf("%s-%d", "node", i),
	//		Status:   1,
	//		Color:    "all",
	//	})
	//
	//	if err := n.Register(); err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	go n.Heartbeat(context.Background())
	//	time.Sleep(2 * time.Second)
	//}
	n := NewNode(e, 2, &NodeInfo{
		Path:     "discovery",
		Env:      "dev",
		SrvName:  fmt.Sprintf("%s-%d", "web-service", 1),
		Addr:     "172.16.7.16",
		Hostname: fmt.Sprintf("%s-%d", "node", 1),
		Status:   1,
		Color:    "all",
	})

	if err := n.Register(); err != nil {
		t.Fatal(err)
	}

	go n.Heartbeat(context.Background())

	//time.Sleep(time.Second * 5)
	//if err := n.Revoke(context.Background()); err != nil {
	//	t.Fatal(err)
	//}

	time.Sleep(20 * time.Second)
}
