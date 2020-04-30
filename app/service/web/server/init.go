package server

import (
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	"github.com/liangjfblue/cheetah/app/service/web/config"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"

	ot "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"

	authv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/app/service/web/model"

	"github.com/liangjfblue/cheetah/common/tracer"

	authSrv "github.com/liangjfblue/cheetah/app/service/web/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	"github.com/liangjfblue/cheetah/common/logger"
)

type Server struct {
	serviceName    string
	serviceVersion string

	service micro.Service
	Tracer  *tracer.Tracer
}

func NewServer(serviceName, serviceVersion string) *Server {
	s := new(Server)

	s.serviceName = serviceName
	s.serviceVersion = serviceVersion

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

	model.Init()

	s.Tracer.Init()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.ConfigInstance().EtcdConf.Addrs
		op.Timeout = time.Duration(config.ConfigInstance().EtcdConf.Timeout) * time.Second
	})

	s.service = micro.NewService(
		micro.Name(s.serviceName),
		micro.Version(s.serviceVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(ot.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(reg),
	)

	s.service.Init()

	s.initRegisterHandler()
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
