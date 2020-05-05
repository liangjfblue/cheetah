/*
@Time : 2020/5/4 17:19
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Get(ctx context.Context, req *models.RoleGetRequest) (*models.RoleGetRespond, error) {
	result, err := service.RoleSrvClient.Get(ctx, &v1.RoleGetRequest{
		ID: uint32(req.Id),
	})
	if err != nil {
		logger.Error("web web role Get err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrRoleGet
		}
		return nil, err
	}

	resp := &models.RoleGetRespond{}
	if err := copier.Copy(resp, result); err != nil {
		logger.Error("web web role Get err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
