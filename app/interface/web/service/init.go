package service

import (
	"github.com/liangjfblue/cheetah/app/interface/web/api"
	userv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/http/middleware"
	"github.com/micro/go-micro/client"
)

var (
	AuthMid       *middleware.Auth
	UserSrvClient userv1.UserService
)

func InitSrvRpc(cli client.Client) {
	AuthMid = middleware.New(cli)
	UserSrvClient = api.NewUserSrvClient(cli)
}
