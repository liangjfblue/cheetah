/*
@Time : 2020/5/4 17:01
@Author : liangjiefan
*/
package menus

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

func Add(ctx context.Context, req *models.MenuAddRequest) (*models.MenuAddRespond, error) {
	var (
		err            error
		menuAddRequest webV1.MenuAddRequest
	)

	if err = copier.Copy(&menuAddRequest, *req); err != nil {
		return nil, errno.ErrCopy
	}

	result, err := service.MenuSrvClient.Add(ctx, &menuAddRequest)
	if err != nil {
		logger.Error("web web menu Add err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuAdd
		}
		return nil, err
	}

	resp := &models.MenuAddRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web menu Add err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
