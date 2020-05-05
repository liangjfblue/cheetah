package service

import (
	"github.com/liangjfblue/cheetah/app/interface/web/api"
	webv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/http/middleware"
	"github.com/micro/go-micro/client"
)

var (
	AuthMid       *middleware.Auth
	CasBinMid     *middleware.CasBin
	UserSrvClient webv1.UserService
	RoleSrvClient webv1.RoleService
	MenuSrvClient webv1.MenuService
)

func InitSrvRpc(cli client.Client) {
	AuthMid = middleware.New(cli)
	CasBinMid = middleware.NewCasBin(cli)
	UserSrvClient = api.NewUserSrvClient(cli)
	RoleSrvClient = api.NewRoleSrvClient(cli)
	MenuSrvClient = api.NewMenuSrvClient(cli)
}
