package cache

type Options struct {
	Name string
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
