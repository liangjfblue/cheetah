package users

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Get(ctx context.Context, req *models.GetRequest) (*models.GetRespond, error) {
	result, err := service.UserSrvClient.Get(ctx, &v1.GetRequest{
		Uid: req.Uid,
	})
	if err != nil {
		logger.Error("web web Get err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyReqyest
		} else {
			err = errno.ErrUserInfo
		}
		return nil, err
	}

	resp := &models.GetRespond{}
	if err := copier.Copy(resp, result); err != nil {
		logger.Error("web web Get err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
