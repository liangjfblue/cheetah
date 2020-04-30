package balancer

type IBalancer interface {
	Init(...Option)
	Options() Options
	DoBalance([]*Instance, ...Option) (*Instance, error)
	String() string
}

type Option func(o *Options)
