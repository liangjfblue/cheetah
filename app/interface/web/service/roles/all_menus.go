/*
@Time : 2020/5/4 17:02
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	webV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func AllMenus(ctx context.Context, req *models.RoleAllMenusRequest) (*models.RoleAllMenusRespond, error) {
	result, err := service.RoleSrvClient.AllMenus(ctx, &webV1.RoleAllMenusRequest{
		RoleId: req.RoleId,
	})
	if err != nil {
		logger.Error("web web role AllMenus err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuUpdate
		}
		return nil, err
	}

	resp := &models.RoleAllMenusRespond{}
	resp.Code = result.Code
	for _, id := range result.MenuIds {
		resp.MenuIds = append(resp.MenuIds, uint(id))
	}

	return resp, nil
}
