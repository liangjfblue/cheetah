package users

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Auth(ctx context.Context, req *models.AuthRequest) (*models.AuthResponse, error) {
	result, err := service.UserSrvClient.Auth(ctx, &v1.AuthRequest{
		Token: req.Token,
	})
	if err != nil {
		err = errno.ErrUserAuthMid
		logger.Error("web web Info err: %s", err.Error())
		return nil, err
	}

	resp := &models.AuthResponse{}
	if err := copier.Copy(&resp, result); err != nil {
		err = errno.ErrCopy
		logger.Error("web web Info err: %s", err.Error())
		return nil, err
	}

	return resp, nil
}
