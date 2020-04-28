package users

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/liangjfblue/cheetah/app/interface/web/service"

	userV1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/interface/web/models"
)

func Register(ctx context.Context, req *models.RegisterRequest) (*models.RegisterRespond, error) {
	fmt.Println(req)
	result, err := service.UserSrvClient.Register(ctx, &userV1.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
		Addr:     req.Addr,
	})
	if err != nil {
		logger.Error("web web Register err: %s", err.Error())
		err = errno.ErrUserRegister
		return nil, err
	}

	fmt.Println(result)

	resp := &models.RegisterRespond{}
	if err := copier.Copy(&resp, result); err != nil {
		err = errno.ErrCopy
		logger.Error("web web Info err: %s", err.Error())
		return nil, err
	}

	return resp, nil
}
