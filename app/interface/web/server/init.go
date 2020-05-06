package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"

	ratelimit2 "github.com/juju/ratelimit"

	"github.com/liangjfblue/cheetah/app/interface/web/service"

	"github.com/liangjfblue/cheetah/app/interface/web/config"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"

	"github.com/liangjfblue/cheetah/common/tracer"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/router"

	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2"
)

type Server struct {
	serviceName    string
	serviceVersion string

	Service micro.Service
	Router  *router.Router

	Tracer *tracer.Tracer
}

func NewServer(serviceName, serviceVersion string) *Server {
	s := new(Server)

	s.serviceName = serviceName
	s.serviceVersion = serviceVersion

	s.Router = router.NewRouter()
	s.Tracer = tracer.New(configs.TraceAddr, s.serviceName)

	return s
}

func (s *Server) Init() {
	config.Init()

	logger.Init(
		gglog.Name(config.ConfigInstance().LogConf.Name),
		gglog.Level(config.ConfigInstance().LogConf.Level),
		gglog.LogDir(config.ConfigInstance().LogConf.LogDir),
		gglog.OpenAccessLog(config.ConfigInstance().LogConf.OpenAccessLog),
	)

	s.Tracer.Init()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.ConfigInstance().EtcdConf.Addrs
		op.Timeout = time.Duration(config.ConfigInstance().EtcdConf.Timeout) * time.Second
	})

	//令牌限流 初始化令牌桶的容量为10, 每1秒往桶放1个令牌
	//即限流 10 request/s
	bRate := ratelimit2.NewBucketWithRate(1, 1000)
	s.Service = micro.NewService(
		micro.Name(s.serviceName),
		micro.Version(s.serviceVersion),
		micro.Registry(reg),
		micro.WrapClient(ratelimit.NewClientWrapper(bRate, false)), //加入限流功能, false为不等待(超限即返回请求失败)
		micro.WrapClient(hystrix.NewClientWrapper()),               // 加入熔断功能, 处理rpc调用失败的情况
	)

	s.Service.Init()

	//init rpc services client
	cli := s.Service.Client()
	service.InitSrvRpc(cli)

	s.Router.Init()
}

func (s *Server) Run() {
	defer func() {
		logger.Info("web close, clean and close something")
		s.Tracer.Close()
	}()

	logger.Debug("web server run")
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", config.ConfigInstance().HttpConf.Port), s.Router.G).Error())
}
