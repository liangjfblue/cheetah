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

func Delete(ctx context.Context, req *models.MenuDeleteRequest) (*models.MenuDeleteRespond, error) {
	var rpcReq webV1.MenuDeleteRequest
	rpcReq.MenuIds = make([]int32, 0)
	for _, id := range req.Id {
		rpcReq.MenuIds = append(rpcReq.MenuIds, int32(id))
	}

	result, err := service.MenuSrvClient.Delete(ctx, &rpcReq)
	if err != nil {
		logger.Error("web web menu Delete err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrMenuDelete
		}
		return nil, err
	}

	resp := &models.MenuDeleteRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web menu Delete err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
