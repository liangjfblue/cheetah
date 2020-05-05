/*
@Time : 2020/5/3 11:43
@Author : liangjiefan
*/
package users

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	userV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func SetRole(ctx context.Context, req *models.UserSetRoleRequest) (*models.UserSetRoleRespond, error) {
	result, err := service.UserSrvClient.SetRole(ctx, &userV1.UserSetRoleRequest{
		UserId: int32(req.UserId),
		RoleId: int32(req.RoleId),
	})
	if err != nil {
		logger.Error("web web SetRole err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserSetRole
		}
		return nil, err
	}

	resp := &models.UserSetRoleRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web SetRole err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
