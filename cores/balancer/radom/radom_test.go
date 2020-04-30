package radom

import (
	"testing"

	"github.com/liangjfblue/cheetah/cores/balancer"
)

func TestRoundBalancer(t *testing.T) {
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
	for i := 0; i < 18; i++ {
		t.Log(b.DoBalance(ins))
	}
}
