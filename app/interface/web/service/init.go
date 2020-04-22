package service

import (
	"github.com/liangjfblue/cheetah/app/interface/web/api"
	"github.com/liangjfblue/cheetah/common/http/middleware"
)

var (
	AuthMid       = middleware.New()
	UserSrvClient = api.NewUserSrvClient()
)
