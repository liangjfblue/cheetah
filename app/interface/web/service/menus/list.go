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

func List(ctx context.Context, req *models.MenuListRequest) (*models.MenuListRespond, error) {
	isLimit := true
	if req.Page <= 0 || req.PageSize <= 0 {
		isLimit = false
	}
	result, err := service.MenuSrvClient.List(ctx, &v1.MenuListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		IsLimit:  isLimit,
		Search:   req.Name,
	})
	if err != nil {
		logger.Error("web web menu List err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuList
		}
		return nil, err
	}

	resp := new(models.MenuListRespond)
	resp.Menus = make([]models.Menu, 0)

	resp.Code = result.Code
	resp.Count = result.Count
	for _, one := range result.All {
		resp.Menus = append(resp.Menus, models.Menu{
			URL:         one.URL,
			Name:        one.URL,
			ParentID:    one.ParentID,
			Sequence:    one.Sequence,
			MenuType:    one.MenuType,
			MenuCode:    one.MenuCode,
			Icon:        one.Icon,
			OperateType: one.OperateType,
			IsAvailable: one.IsAvailable,
			Remark:      one.Remark,
		})
	}

	return resp, nil
}
