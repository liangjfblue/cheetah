package radom

import (
	"math/rand"

	"github.com/liangjfblue/cheetah/cores/balancer"
)

type randomBalancer struct {
	opts balancer.Options
}

func (b *randomBalancer) Init(opts ...balancer.Option) {
	for _, o := range opts {
		o(&b.opts)
	}
}

func (b *randomBalancer) Options() balancer.Options {
	return b.opts
}

func (b *randomBalancer) DoBalance(ins []*balancer.Instance, opts ...balancer.Option) (*balancer.Instance, error) {
	if len(ins) <= 0 {
		return nil, balancer.ErrEmptyInstance
	}

	idx := rand.Intn(len(ins))
	return ins[idx], nil
}
func (b *randomBalancer) String() string {
	return "randomBalancer"
}

func New(opts ...balancer.Option) balancer.IBalancer {
	b := new(randomBalancer)
	b.opts = balancer.DefaultOptions

	for _, o := range opts {
		o(&b.opts)
	}

	return b
}
