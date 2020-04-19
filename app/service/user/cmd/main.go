package main

import (
	"github.com/liangjfblue/cheetah/app/service/user/server"
)

const (
	srvName    = "micro.srv.user"
	srvVersion = "v1.0.0"
)

func main() {
	srv := server.NewServer(srvName, srvVersion)
	srv.Init()

	srv.Run()
}
