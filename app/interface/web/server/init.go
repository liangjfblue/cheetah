package server

import (
	"fmt"
	"net/http"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"
	"github.com/micro/go-micro"

	ratelimit2 "github.com/juju/ratelimit"
	"github.com/liangjfblue/cheetah/common/tracer"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/config"
	"github.com/liangjfblue/cheetah/app/interface/web/router"
)

type Server struct {
	serviceName    string
	serviceVersion string

	Config *config.Config

	Service micro.Service
	Router  *router.Router

	Tracer *tracer.Tracer
}

func NewServer(serviceName, serviceVersion string) *Server {
	s := new(Server)

	s.serviceName = serviceName
	s.serviceVersion = serviceVersion

	s.Config = config.NewConfig()
	s.Router = router.NewRouter()
	s.Tracer = tracer.New(configs.TraceAddr, s.serviceName)

	return s
}

func (s *Server) Init() {
	logger.Init(
		gglog.Name("srv-web"),
		gglog.Level(1),
		gglog.LogDir("./logs"),
		gglog.OpenAccessLog(true),
	)

	s.Tracer.Init()

	//register := etcdv3.NewRegistry(
	//	registry.Addrs("172.16.7.16:9002", "172.16.7.16:9004", "172.16.7.16:9006"),
	//)

	//配置请求容量及qps
	bRate := ratelimit2.NewBucketWithRate(100, 1000)
	s.Service = micro.NewService(
		micro.Name(s.serviceName),
		micro.Version(s.serviceVersion),
		//web.Registry(register),
		micro.WrapClient(ratelimit.NewClientWrapper(bRate, false)), //加入限流功能, false为不等待(超限即返回请求失败)
		micro.WrapClient(hystrix.NewClientWrapper()),               // 加入熔断功能, 处理rpc调用失败的情况(cirucuit breaker)
	)

	s.Service.Init()

	s.Router.Init()
}

func (s *Server) Run() {
	defer func() {
		logger.Info("web close, clean and close something")
		s.Tracer.Close()
	}()

	logger.Debug("web server run")
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", s.Config.HttpConf.Port), s.Router.G).Error())
}
