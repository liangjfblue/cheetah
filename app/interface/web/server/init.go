package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"
	"github.com/micro/go-micro"

	"github.com/liangjfblue/cheetah/common/tracer"

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

	//TODO etcd初始化

	//TODO 拉取服务列表

	s.Router.Init()
}

func (s *Server) Run() {
	defer func() {
		logger.Info("web close, clean and close something")
		s.Tracer.Close()
	}()

	log.Println("service web server run")
	logger.Debug("web server run")
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", s.Config.HttpConf.Port), s.Router.G).Error())
}
