package api

import (
	"github.com/liangjfblue/cheetah/common/proto"

	userv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/micro/go-micro/v2/client"
)

func NewUserSrvClient(cli client.Client) userv1.UserService {
	return userv1.NewUserService(proto.WebSrvName, cli)
}

func NewRoleSrvClient(cli client.Client) userv1.RoleService {
	return userv1.NewRoleService(proto.WebSrvName, cli)
}
func NewMenuSrvClient(cli client.Client) userv1.MenuService {
	return userv1.NewMenuService(proto.WebSrvName, cli)
}
