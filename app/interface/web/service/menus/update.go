/*
@Time : 2020/5/4 17:01
@Author : liangjiefan
*/
package menus

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

func Update(ctx context.Context, req *models.MenuUpdateRequest) (*models.MenuUpdateRespond, error) {
	result, err := service.MenuSrvClient.Update(ctx, &webV1.MenuUpdateRequest{
		ID:          uint32(req.Id),
		URL:         req.URL,
		Name:        req.Name,
		ParentID:    req.ParentID,
		Sequence:    req.Sequence,
		MenuType:    req.MenuType,
		MenuCode:    req.MenuCode,
		Icon:        req.Icon,
		OperateType: req.OperateType,
		IsAvailable: req.IsAvailable,
		Remark:      req.Remark,
	})
	if err != nil {
		logger.Error("web web menu Update err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuUpdate
		}
		return nil, err
	}

	resp := &models.MenuUpdateRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web menu Update err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
