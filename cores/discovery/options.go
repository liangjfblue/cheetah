package discovery

import (
	"context"
	"time"
)

//注册中心
type Options struct {
	Addrs   []string
	Timeout time.Duration
	Context context.Context
}

//注册参数
type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}

//监听参数
type WatchOptions struct {
	Service string
	Context context.Context
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func RegisterTTL(t time.Duration) RegisterOption {
	return func(o *RegisterOptions) {
		o.TTL = t
	}
}

func WatchService(name string) WatchOption {
	return func(o *WatchOptions) {
		o.Service = name
	}
}
