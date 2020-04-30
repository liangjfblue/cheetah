package radom

import (
	"fmt"
	"testing"

	"github.com/liangjfblue/cheetah/cores/balancer"
)

func TestHashBalancer(t *testing.T) {
	ins := []*balancer.Instance{
		{
			Ip:   "127.0.0.1",
			Port: 1,
		},
		{
			Ip:   "127.0.0.2",
			Port: 2,
		},
		{
			Ip:   "127.0.0.3",
			Port: 3,
		},
		{
			Ip:   "127.0.0.4",
			Port: 4,
		},
	}

	b := New()
	for i := 0; i < 10; i++ {
		t.Log(b.DoBalance(ins, balancer.WithKey(fmt.Sprint(i))))
	}
	t.Log("------------------")
	for i := 0; i < 10; i++ {
		t.Log(b.DoBalance(ins, balancer.WithKey("test")))
	}
}
