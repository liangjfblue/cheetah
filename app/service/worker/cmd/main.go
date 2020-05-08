package main

import (
	"github.com/liangjfblue/cheetah/app/service/worker/server"
	"github.com/liangjfblue/cheetah/common/proto"
)

const (
	srvName    = proto.WorkerSrvName
	srvVersion = proto.WorkerSrvVersion
)

func main() {
	srv := server.NewServer(srvName, srvVersion)
	srv.Init()

	srv.Run()
}
