/*
@Time : 2020/5/4 17:01
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	webV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Update(ctx context.Context, req *models.RoleUpdateRequest) (*models.RoleUpdateRespond, error) {
	result, err := service.RoleSrvClient.Update(ctx, &webV1.RoleUpdateRequest{
		ID:          uint32(req.Id),
		RoleName:    req.RoleName,
		RoleDesc:    req.RoleDesc,
		IsAvailable: req.IsAvailable,
		IsAdmin:     req.IsAdmin,
	})
	if err != nil {
		logger.Error("web web role Update err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrRoleUpdate
		}
		return nil, err
	}

	resp := &models.RoleUpdateRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web role Update err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
