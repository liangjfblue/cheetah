package users

import (
	"context"

	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/user/proto/v1"
)

func Login() (interface{}, error) {
	resp, err := service.UserSrvClient.Login(context.Background(), &v1.LoginRequest{
		Username: "",
		Password: "",
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
