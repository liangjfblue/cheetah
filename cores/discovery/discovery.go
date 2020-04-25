package discovery

import "context"

type ISubscribe interface {
	//监听服务
	Watch(ctx context.Context, path string)

	//获取服务
	Get(ctx context.Context, srvName string) (*NodeInfo, error)

	//获取服务
	All(ctx context.Context, srvName string) ([]NodeInfo, error)
}

type IPublish interface {
	//注册服务
	Register(ctx context.Context, info NodeInfo) error

	//销毁服务
	Revoke(ctx context.Context) error

	//续租服务
	Heartbeat(ctx context.Context)
}
