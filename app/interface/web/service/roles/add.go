/*
@Time : 2020/5/4 17:01
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	webV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Add(ctx context.Context, req *models.RoleAddRequest) (*models.RoleAddRespond, error) {
	var (
		err            error
		roleAddRequest webV1.RoleAddRequest
	)

	if err = copier.Copy(&roleAddRequest, *req); err != nil {
		return nil, errno.ErrCopy
	}

	result, err := service.RoleSrvClient.Add(ctx, &roleAddRequest)
	if err != nil {
		logger.Error("web web role Add err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrRoleAdd
		}
		return nil, err
	}

	resp := &models.RoleAddRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web role Add err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
