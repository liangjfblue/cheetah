/*
@Time : 2020/5/4 17:02
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

func SetMenus(ctx context.Context, req *models.RoleSetMenusRequest) (*models.RoleSetMenusRespond, error) {
	in := webV1.RoleSetMenusRequest{
		RoleId: req.RoleId,
	}
	if in.MenuIds == nil {
		in.MenuIds = make(map[uint32]uint32)
	}
	for k, id := range req.MenuIds {
		in.MenuIds[uint32(k)] = uint32(id)
	}

	result, err := service.RoleSrvClient.SetMenus(ctx, &in)
	if err != nil {
		logger.Error("web web role SetMenus err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrRoleSetMenus
		}
		return nil, err
	}

	resp := &models.RoleSetMenusRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web role SetMenus err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
