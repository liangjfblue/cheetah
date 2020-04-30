package etcdv3

import (
	"testing"
)

func TestClient(t *testing.T) {
	//r := NewDiscovery(
	//	discovery.Addrs([]string{"http://172.16.7.16:9002", "http://172.16.7.16:9004", "http://172.16.7.16:9006"}...),
	//	discovery.Timeout(time.Second*5),
	//)
	//
	//service := &discovery.Service{
	//	SrvName: "user",
	//	Version: "1.0.0",
	//	Nodes: []*discovery.Node{
	//		{
	//			Id:      uuid.New().String(),
	//			Address: "172.16.7.16:8899",
	//		},
	//	},
	//}
	//
	//if err := r.Register(service, discovery.RegisterTTL(time.Second*3)); err != nil {
	//	log.Fatal(err)
	//}
	//defer r.Deregister(service)
	//
	//t.Log("had register service")
	//select {}
}
