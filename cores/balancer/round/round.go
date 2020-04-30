package round

import (
	"github.com/liangjfblue/cheetah/cores/balancer"
)

type roundBalancer struct {
	opts balancer.Options
}

func (b *roundBalancer) Init(opts ...balancer.Option) {
	for _, o := range opts {
		o(&b.opts)
	}
}

func (b *roundBalancer) Options() balancer.Options {
	return b.opts
}

func (b *roundBalancer) DoBalance(ins []*balancer.Instance, opts ...balancer.Option) (*balancer.Instance, error) {
	max := uint32(len(ins))
	if max <= 0 {
		return nil, balancer.ErrEmptyInstance
	}

	defer func() {
		b.opts.Index++
	}()

	for _, o := range opts {
		o(&b.opts)
	}

	if b.opts.Index >= max {
		b.opts.Index = 0
	}

	return ins[b.opts.Index], nil
}

func (b *roundBalancer) String() string {
	return "round"
}

func New(opts ...balancer.Option) balancer.IBalancer {
	b := new(roundBalancer)
	b.opts = balancer.DefaultOptions

	for _, o := range opts {
		o(&b.opts)
	}

	return b
}
