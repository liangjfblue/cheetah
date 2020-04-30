package cache

import (
	"testing"
)

func TestCache_GetService(t *testing.T) {
	//d := etcdv3.NewDiscovery(
	//	discovery.Addrs([]string{"172.16.7.16:9002", "172.16.7.16:9004", "172.16.7.16:9006"}...),
	//	discovery.Timeout(time.Second*5),
	//)
	//
	////30秒拉取一次服务列表
	//c := New(d, WithTTL(30*time.Second))
	//
	//for i := 0; i < 10; i++ {
	//	start := time.Now().UnixNano()
	//	services, err := c.GetService("user")
	//	if err != nil {
	//		//TODO 从本地配置获取服务的备用地址
	//		t.Log(err)
	//	}
	//	for _, service := range services {
	//		t.Log(service.SrvName)
	//		for _, node := range service.Nodes {
	//			t.Log(node)
	//		}
	//	}
	//	t.Log(fmt.Sprintf("cost ms:%d", (time.Now().UnixNano()-start)/1e6))
	//	time.Sleep(time.Second * 5)
	//}
}
