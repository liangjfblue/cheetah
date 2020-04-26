package etcdv3

import (
	"context"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func TestConn(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.112:9002", "192.168.0.112:9004", "192.168.0.112:9006"},
		DialTimeout: time.Second * time.Duration(3),
	})

	if err != nil {
		t.Fatal(err)
	}

	e := NewEtcd(client)
	t.Log(e)
}

func TestPutGet(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.112:9002", "192.168.0.112:9004", "192.168.0.112:9006"},
		DialTimeout: time.Second * time.Duration(3),
	})
	if err != nil {
		t.Fatal(err)
	}

	e := NewEtcd(client)

	putResp, err := e.Client.Put(context.TODO(), "name", "liangjf123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(putResp)

	getResp, err := e.Client.Get(context.TODO(), "name")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(getResp)
}
