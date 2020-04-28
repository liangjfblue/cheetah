package users

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
)

func Login(ctx context.Context, req *models.LoginRequest) (*models.LoginRespond, error) {
	result, err := service.UserSrvClient.Login(ctx, &v1.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logger.Error("web web Login err: %s", err.Error())
		err = errno.ErrUserLogin
		return nil, err
	}

	fmt.Println(result)
	resp := &models.LoginRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web Info err: %s", err.Error())
		err = errno.ErrCopy
		return nil, err
	}

	return resp, nil
}
