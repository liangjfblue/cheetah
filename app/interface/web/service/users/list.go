package users

import (
	"context"
	"strings"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
)

func List(ctx context.Context, req *models.UserListRequest) (*models.UserListRespond, error) {
	result, err := service.UserSrvClient.List(ctx, &v1.UserListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: req.Username,
	})
	if err != nil {
		logger.Error("web web List err: %s", err.Error())
		if strings.Contains(err.Error(), "too many request") {
			err = errno.ErrTooManyRequest
		} else {
			err = errno.ErrUserList
		}
		return nil, err
	}

	resp := new(models.UserListRespond)
	resp.Users = make([]models.User, 0)

	resp.Code = result.Code
	resp.Count = result.Count
	for _, one := range result.All {
		resp.Users = append(resp.Users, models.User{
			Username: one.Username,
			Age:      one.Age,
			Addr:     one.Addr,
		})
	}

	return resp, nil
}
