package radom

import (
	"hash/crc32"

	"github.com/liangjfblue/cheetah/cores/balancer"
)

type hashBalancer struct {
	opts balancer.Options
}

func (b *hashBalancer) Init(opts ...balancer.Option) {
	for _, o := range opts {
		o(&b.opts)
	}
}

func (b *hashBalancer) Options() balancer.Options {
	return b.opts
}

func (b *hashBalancer) DoBalance(ins []*balancer.Instance, opts ...balancer.Option) (*balancer.Instance, error) {
	max := uint32(len(ins))
	if max <= 0 {
		return nil, balancer.ErrEmptyInstance
	}

	for _, o := range opts {
		o(&b.opts)
	}

	idxSum := crc32.ChecksumIEEE([]byte(b.opts.Key))
	idx := idxSum % max

	return ins[idx], nil
}
func (b *hashBalancer) String() string {
	return "hashBalancer"
}

func New(opts ...balancer.Option) balancer.IBalancer {
	b := new(hashBalancer)
	b.opts = balancer.DefaultOptions

	for _, o := range opts {
		o(&b.opts)
	}

	return b
}
