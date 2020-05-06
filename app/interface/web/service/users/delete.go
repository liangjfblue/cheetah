/*
@Time : 2020/5/3 11:56
@Author : liangjiefan
*/
package users

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

func Delete(ctx context.Context, req *models.UserDeleteRequest) (*models.UserDeleteRespond, error) {
	var rpcReq webV1.UserDeleteRequest
	rpcReq.UserIds = make([]int32, 0)
	for _, id := range req.Id {
		rpcReq.UserIds = append(rpcReq.UserIds, int32(id))
	}

	result, err := service.UserSrvClient.Delete(ctx, &rpcReq)
	if err != nil {
		logger.Error("web web Delete err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserAdd
		}
		return nil, err
	}

	resp := &models.UserDeleteRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web Delete err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
