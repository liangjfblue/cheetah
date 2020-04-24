package api

import (
	"time"

	"github.com/liangjfblue/cheetah/common/proto"

	userv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/micro/go-micro/client"
)

func NewUserSrvClient() userv1.UserService {
	c := client.NewClient(
		client.Retries(2),
		client.DialTimeout(time.Minute*2),
	)

	return userv1.NewUserService(proto.UserSrvName, c)
}
