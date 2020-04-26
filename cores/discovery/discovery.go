package discovery

import (
	"context"
)

type IDiscovery interface {
	Init(...Option) error
	Options() Options
	Register(context.Context, *Service, ...RegisterOption) error
	Deregister(context.Context, *Service) error

	Watch(context.Context, ...WatchOption)
	Get(context.Context, string) ([]*Service, error)
	All(context.Context, string) ([]*Service, error)
}

type Option func(*Options)
type RegisterOption func(*RegisterOptions)
type WatchOption func(*WatchOptions)
