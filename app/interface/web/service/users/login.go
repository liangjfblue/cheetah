package users

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
)

func Login(ctx context.Context, req *models.UserLoginRequest) (*models.UserLoginRespond, error) {
	result, err := service.UserSrvClient.Login(ctx, &v1.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logger.Error("web web Login err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserLogin
		}
		return nil, err
	}

	resp := &models.UserLoginRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web Info err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
