package users

import (
	"context"

	"github.com/jinzhu/copier"
	v1 "github.com/liangjfblue/cheetah/app/service/user/proto/v1"

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
		err = errno.ErrUserInfo
		logger.Error("web user Info err: %s", err.Error())
		return nil, err
	}

	resp := &models.GetRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		err = errno.ErrCopy
		logger.Error("web user Info err: %s", err.Error())
		return nil, err
	}

	return resp, nil
}
