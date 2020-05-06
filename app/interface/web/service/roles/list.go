/*
@Time : 2020/5/4 17:01
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func List(ctx context.Context, req *models.RoleListRequest) (*models.RoleListRespond, error) {
	isLimit := true
	if req.Page <= 0 || req.PageSize <= 0 {
		isLimit = false
	}
	result, err := service.RoleSrvClient.List(ctx, &v1.RoleListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		IsLimit:  isLimit,
		Search:   req.Name,
	})
	if err != nil {
		logger.Error("web web role List err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrRoleList
		}
		return nil, err
	}

	resp := new(models.RoleListRespond)
	resp.Roles = make([]models.Role, 0)

	resp.Code = result.Code
	resp.Count = result.Count
	for _, one := range result.All {
		resp.Roles = append(resp.Roles, models.Role{
			RoleName:    one.RoleName,
			RoleDesc:    one.RoleDesc,
			IsAvailable: one.IsAvailable,
			IsAdmin:     one.IsAdmin,
			IsBase:      one.IsBase,
			Sequence:    one.Sequence,
			ParentID:    one.ParentID,
		})
	}

	return resp, nil
}
