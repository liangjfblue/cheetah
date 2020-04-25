package discovery

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func TestMaster_WatchEvent(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.112:9002", "192.168.0.112:9004", "192.168.0.112:9006"},
		DialTimeout: time.Second * time.Duration(3),
	})
	if err != nil {
		t.Fatal(err)
	}

	e := NewEtcd(client)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*60)
	defer cancel()

	m := NewMaster(e)

	go m.Watch(ctx, "discovery")

	ticker := time.NewTicker(time.Duration(3) * time.Second)
	for {
		select {
		case <-ticker.C:
			t.Log(m.All(ctx, fmt.Sprintf("%s-%d", "web-user", 1)))
		case <-ctx.Done():
			return
		}
	}
}
