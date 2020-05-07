/*
@Time : 2020/5/4 17:03
@Author : liangjiefan
*/
package menus

import (
	"context"
	"strings"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	webV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func MenuButtons(ctx context.Context, req *models.MenuMenuButtonsRequest) (*models.MenuMenuButtonsRespond, error) {
	result, err := service.MenuSrvClient.MenuButtons(ctx, &webV1.MenuButtonsRequest{
		UserId:   req.UserId,
		MenuCode: req.MenuCode,
	})
	if err != nil {
		logger.Error("web web menu MenuButtons err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuMenuButtons
		}
		return nil, err
	}

	resp := &models.MenuMenuButtonsRespond{}
	resp.Code = result.Code
	for _, op := range result.OperateType {
		resp.OperateTypes = append(resp.OperateTypes, op)
	}

	return resp, nil
}
