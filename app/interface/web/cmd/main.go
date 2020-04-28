package main

import (
	"github.com/liangjfblue/cheetah/app/interface/web/server"
)

const (
	srvName    = "cheetah.web.web"
	srvVersion = "v1.0.0"
)

func main() {
	srv := server.NewServer(srvName, srvVersion)
	srv.Init()

	srv.Run()
}
