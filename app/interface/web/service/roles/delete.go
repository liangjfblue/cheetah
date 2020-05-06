/*
@Time : 2020/5/4 17:01
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

func Delete(ctx context.Context, req *models.RoleDeleteRequest) (*models.RoleDeleteRespond, error) {
	var rpcReq webV1.RoleDeleteRequest
	rpcReq.RoleIds = make([]int32, 0)
	for _, id := range req.Id {
		rpcReq.RoleIds = append(rpcReq.RoleIds, int32(id))
	}

	result, err := service.RoleSrvClient.Delete(ctx, &rpcReq)
	if err != nil {
		logger.Error("web web role Delete err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserDelete
		}
		return nil, err
	}

	resp := &models.RoleDeleteRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web role Delete err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
