package etcdv3

import (
	"testing"
)

func TestServer(t *testing.T) {
	//r := NewDiscovery(
	//	discovery.Addrs([]string{"172.16.7.16:9002", "172.16.7.16:9004", "172.16.7.16:9006"}...),
	//	discovery.Timeout(time.Second*5),
	//)
	//
	//w, err := r.Watch(discovery.WatchService("user"))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	//
	//t.Log("now to watch event")
	//defer w.Stop()
	//for {
	//	select {
	//	case <-ch:
	//		t.Log("get a signal, return")
	//		return
	//	default:
	//		resp, err := w.Next()
	//		if err != nil {
	//			t.Fatal(err)
	//			return
	//		}
	//
	//		//TODO you can update cache
	//		t.Log(resp.Action)
	//		t.Log(resp.Service)
	//		for _, node := range resp.Service.Nodes {
	//			t.Log(*node)
	//		}
	//	}
	//}
}
