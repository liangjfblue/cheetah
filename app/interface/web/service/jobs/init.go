package article

import (
	"github.com/liangjfblue/cheetah/app/interface/web/api"
	userv1 "github.com/liangjfblue/cheetah/app/service/user/proto/v1"
)

type Srv struct {
	userSrvClient userv1.UserService
}

func NewSrv() *Srv {
	return &Srv{
		userSrvClient: api.NewUserSrvClient(),
	}
}
