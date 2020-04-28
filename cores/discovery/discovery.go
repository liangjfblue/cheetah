package discovery

type IDiscovery interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error

	Watch(...WatchOption) (Watcher, error)
	Get(string) ([]*Service, error)
	All() ([]*Service, error)
	String() string
}

type Option func(*Options)
type RegisterOption func(*RegisterOptions)
type WatchOption func(*WatchOptions)
