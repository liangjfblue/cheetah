package balancer

type Options struct {
	Index uint32
	Key   string
}

var (
	DefaultOptions = Options{
		Index: 0,
		Key:   "dead",
	}
)

func WithIndex(index uint32) Option {
	return func(o *Options) {
		o.Index = index
	}
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}
