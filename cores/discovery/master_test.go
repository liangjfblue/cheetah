package discovery

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMaster_WatchEvent(t *testing.T) {
	e := NewEtcd([]string{"172.16.7.16:9001", "172.16.7.16:9003", "172.16.7.16:9005"}, time.Duration(3))
	if err := e.InitEtcd(); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	m := NewMaster(e, NewRandSelect())
	go m.WatchEvent(ctx, "discovery")

	select {
	case <-ctx.Done():
		fmt.Println(m.SrvAll("web-service1"))
		return
	}
}
