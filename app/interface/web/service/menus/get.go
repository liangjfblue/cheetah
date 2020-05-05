/*
@Time : 2020/5/4 22:26
@Author : liangjiefan
*/
package roles

import (
	"context"
	"strings"

	webv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func Get(ctx context.Context, req *models.MenuGetRequest) (*models.MenuGetRespond, error) {
	result, err := service.MenuSrvClient.Get(ctx, &webv1.MenuGetRequest{
		ID: uint32(req.Id),
	})
	if err != nil {
		logger.Error("web web menu Get err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuGet
		}
		return nil, err
	}

	resp := &models.MenuGetRespond{}
	if err := copier.Copy(resp, result); err != nil {
		logger.Error("web web menu Get err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
