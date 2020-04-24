package discovery

import (
	"context"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	e := NewEtcd([]string{"172.16.7.16:9001", "172.16.7.16:9003", "172.16.7.16:9005"}, time.Duration(3))
	if err := e.InitEtcd(); err != nil {
		t.Fatal(err)
	}
}

func TestPutGet(t *testing.T) {
	e := NewEtcd([]string{"172.16.7.16:9001", "172.16.7.16:9003", "172.16.7.16:9005"}, time.Duration(3))
	if err := e.InitEtcd(); err != nil {
		t.Fatal(err)
	}

	putResp, err := e.Client.Put(context.TODO(), "name", "liangjf")
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
