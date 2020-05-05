/*
@Time : 2020/5/2 13:41
@Author : liangjiefan
*/
package users

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	userV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
)

func Update(ctx context.Context, req *models.UserUpdateRequest) (*models.UserUpdateRespond, error) {
	result, err := service.UserSrvClient.Update(ctx, &userV1.UserUpdateRequest{
		ID:       uint32(req.Id),
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
		Addr:     req.Addr,
	})
	if err != nil {
		logger.Error("web web Update err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserAdd
		}
		return nil, err
	}

	resp := &models.UserUpdateRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		logger.Error("web web Update err: %s", err.Error())
		return nil, errno.ErrCopy
	}

	return resp, nil
}
