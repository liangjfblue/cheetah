package discovery

import "context"

type IClient interface {
	//监听服务
	Watch(ctx context.Context, path string) error

	//获取服务
	Get(ctx context.Context, path string) error
}

type IWorker interface {
	//注册服务
	Register(ctx context.Context, info NodeInfo) error

	//销毁服务
	Revoke(ctx context.Context, info NodeInfo) error

	//续租服务
	Heartbeat(ctx context.Context) error
}
