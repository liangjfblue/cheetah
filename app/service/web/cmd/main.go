package main

import (
	"github.com/liangjfblue/cheetah/app/service/web/server"
	"github.com/liangjfblue/cheetah/common/proto"
)

const (
	srvName    = proto.WebSrvName
	srvVersion = proto.WebSrvVersion
)

func main() {
	srv := server.NewServer(srvName, srvVersion)
	srv.Init()

	srv.Run()
}
