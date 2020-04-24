package users

import (
	"context"

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
		err = errno.ErrUserLogin
		logger.Error("web web Login err: %s", err.Error())
		return nil, err
	}

	resp := &models.LoginRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		err = errno.ErrCopy
		logger.Error("web web Info err: %s", err.Error())
		return nil, err
	}

	return resp, nil
}
