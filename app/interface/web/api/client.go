package api

import (
	"github.com/liangjfblue/cheetah/common/proto"

	userv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/micro/go-micro/client"
)

func NewUserSrvClient(cli client.Client) userv1.UserService {
	return userv1.NewUserService(proto.UserSrvName, cli)
}
