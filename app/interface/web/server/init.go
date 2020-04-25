package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	"github.com/liangjfblue/cheetah/app/interface/web/config"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"
	"github.com/micro/go-micro"

	"github.com/liangjfblue/cheetah/common/tracer"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/router"
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
		gglog.Name(config.ConfigInstance.LogConf.Name),
		gglog.Level(config.ConfigInstance.LogConf.Level),
		gglog.LogDir(config.ConfigInstance.LogConf.LogDir),
		gglog.OpenAccessLog(config.ConfigInstance.LogConf.OpenAccessLog),
	)

	s.Tracer.Init()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.0.112:9002", "http://192.168.0.112:9004", "http://192.168.0.112:9006",
		}
		op.Timeout = 5 * time.Second //5秒超时
	})

	s.Service = micro.NewService(
		micro.Name(s.serviceName),
		micro.Version(s.serviceVersion),
		micro.Registry(reg),
	)

	s.Service.Init()

	s.Router.Init()
}

func (s *Server) Run() {
	defer func() {
		logger.Info("web close, clean and close something")
		s.Tracer.Close()
	}()

	log.Println("service web server run")
	logger.Debug("web server run")
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", config.ConfigInstance.HttpConf.Port), s.Router.G).Error())
}
