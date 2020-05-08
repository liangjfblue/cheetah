package server

import (
	"time"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"

	"github.com/liangjfblue/cheetah/app/service/worker/config"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/liangjfblue/gglog"

	ot "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	workerv1 "github.com/liangjfblue/cheetah/app/service/worker/proto/v1"

	model "github.com/liangjfblue/cheetah/app/service/worker/models"

	"github.com/liangjfblue/cheetah/common/tracer"

	workerSrv "github.com/liangjfblue/cheetah/app/service/worker/service"
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
	logger.Init(
		gglog.Name(config.GetInstance().LogConf.Name),
		gglog.Level(config.GetInstance().LogConf.Level),
		gglog.LogDir(config.GetInstance().LogConf.LogDir),
		gglog.OpenAccessLog(config.GetInstance().LogConf.OpenAccessLog),
	)

	model.Init()

	s.Tracer.Init()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.GetInstance().EtcdConf.Addrs
		op.Timeout = time.Duration(config.GetInstance().EtcdConf.Timeout) * time.Second
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
	srv := &workerSrv.WorkerService{}
	if err := workerv1.RegisterWorkerHandler(s.service.Server(), srv, server.InternalHandler(true)); err != nil {
		logger.Error("service worker err: %s", err.Error())
		return
	}
}

func (s *Server) Run() {
	defer func() {
		logger.Info("srv worker close, clean and close something")
		s.Tracer.Close()
	}()

	logger.Debug("service worker server run")
	if err := s.service.Run(); err != nil {
		logger.Error("service worker err: %s", err.Error())
		return
	}
}
