package server

import (
	"log"
	"time"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"

	ot "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"

	authv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/app/service/web/model"

	"github.com/liangjfblue/cheetah/app/service/web/config"

	"github.com/liangjfblue/cheetah/common/tracer"

	authSrv "github.com/liangjfblue/cheetah/app/service/web/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	"github.com/liangjfblue/cheetah/common/logger"
)

type Server struct {
	serviceName    string
	serviceVersion string

	Config  *config.Config
	service micro.Service
	Tracer  *tracer.Tracer
}

func NewServer(serviceName, serviceVersion string) *Server {
	s := new(Server)

	s.serviceName = serviceName
	s.serviceVersion = serviceVersion

	s.Config = config.NewConfig()

	s.Tracer = tracer.New(configs.TraceAddr, s.serviceName)

	return s
}

func (s *Server) Init() {
	logger.Init(
		gglog.Name("web-arv"),
		gglog.Level(1),
		gglog.LogDir("./logs"),
		gglog.OpenAccessLog(true),
	)

	model.Init(s.Config.MysqlConf)

	s.Tracer.Init()

	//TODO etcd初始化

	//TODO 服务注册

	s.service = micro.NewService(
		micro.Name(s.serviceName),
		micro.Version(s.serviceVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(ot.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	s.service.Init()

	s.initRegisterHandler()
	log.Println("service web server init")
}

func (s *Server) initRegisterHandler() {
	srv := &authSrv.UserService{}
	if err := authv1.RegisterUserHandler(s.service.Server(), srv, server.InternalHandler(true)); err != nil {
		logger.Error("service web err: %s", err.Error())
		return
	}
}

func (s *Server) Run() {
	defer func() {
		logger.Info("srv web close, clean and close something")
		s.Tracer.Close()
	}()

	logger.Debug("service web server run")
	if err := s.service.Run(); err != nil {
		logger.Error("service web err: %s", err.Error())
		return
	}
}
