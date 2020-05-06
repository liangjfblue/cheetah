package server

import (
	"time"

	"github.com/micro/go-micro/v2/registry"

	webv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/micro/go-plugins/registry/etcdv3/v2"

	"github.com/liangjfblue/cheetah/app/service/web/config"
	"github.com/liangjfblue/cheetah/app/service/web/service"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"

	ot "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	"github.com/liangjfblue/cheetah/app/service/web/models"

	"github.com/liangjfblue/cheetah/common/tracer"

	webSrv "github.com/liangjfblue/cheetah/app/service/web/service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"

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

	models.Init()

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

	//init casbin
	if err := service.InitCasBin(models.DB); err != nil {
		panic(err)
	}

	s.initRegisterHandler()
}

func (s *Server) initRegisterHandler() {
	srv1 := &webSrv.UserService{}
	if err := webv1.RegisterUserHandler(s.service.Server(), srv1, server.InternalHandler(true)); err != nil {
		logger.Error("service web RegisterUserHandler err: %s", err.Error())
		return
	}

	srv2 := &webSrv.RoleService{}
	if err := webv1.RegisterRoleHandler(s.service.Server(), srv2, server.InternalHandler(true)); err != nil {
		logger.Error("service web RegisterRoleHandler err: %s", err.Error())
		return
	}

	srv3 := &webSrv.MenuService{}
	if err := webv1.RegisterMenuHandler(s.service.Server(), srv3, server.InternalHandler(true)); err != nil {
		logger.Error("service web RegisterMenuHandler err: %s", err.Error())
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
