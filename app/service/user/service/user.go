package service

import (
	"context"
	"time"

	"github.com/liangjfblue/cheetah/common/logger"

	v1 "github.com/liangjfblue/cheetah/app/service/user/proto/v1"

	"github.com/liangjfblue/cheetah/common/auth"
	"github.com/liangjfblue/cheetah/common/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liangjfblue/cheetah/common/uuid"

	"github.com/jinzhu/gorm"
	"github.com/liangjfblue/cheetah/app/service/user/model"
	"github.com/pkg/errors"

	"github.com/liangjfblue/cheetah/common/errno"
)

func (s *Service) Register(ctx context.Context, in *v1.RegisterRequest, out *v1.RegisterRespond) error {
	if ctx.Err() == context.Canceled {
		return ctx.Err()
	}

	if _, err := model.GetUser(&model.TBUser{Username: in.Username}); err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, " service user")
	}

	user := model.TBUser{
		Uid:         uuid.UUID(),
		Username:    in.Username,
		Password:    in.Password,
		Age:         in.Age,
		Address:     in.Addr,
		IsAvailable: 1,
		LastLogin:   time.Now(),
	}

	if err := user.Validate(); err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, " service user")
	}

	if err := user.Encrypt(); err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, " service user")
	}

	if err := user.Create(); err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, " service user")
	}

	out.Code = errno.Success.Code
	out.Uid = user.Uid

	return nil
}

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest, out *v1.LoginRespond) error {
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service user")
	}

	var (
		err      error
		user     *model.TBUser
		tokenStr string
	)

	user, err = model.GetUser(&model.TBUser{Username: in.Username})
	if err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	if err = auth.Compare(user.Password, in.Password); err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	if user.IsAvailable != 1 {
		logger.Error("user unavailable")
		return errors.Wrap(errors.New("user unavailable"), "service user")
	}

	user.LastLogin = time.Now()
	if err = user.Update(); err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	tokenStr, err = s.Token.SignToken(token.Context{Uid: user.Uid})
	if err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	out.Code = errno.Success.Code
	out.Token = tokenStr

	return nil
}

func (s *Service) Get(ctx context.Context, in *v1.GetRequest, out *v1.GetRespond) error {
	var (
		err  error
		user *model.TBUser
	)

	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service user")
	}

	user, err = model.GetUser(&model.TBUser{Uid: in.Uid})
	if err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	out = &v1.GetRespond{
		Code:     errno.Success.Code,
		Username: user.Username,
		Age:      user.Age,
		Addr:     user.Address,
	}

	return nil
}

func (s *Service) List(ctx context.Context, in *v1.ListRequest, out *v1.ListRespond) error {
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service user")
	}

	count, users, err := model.ListUsers(in.Username, in.Page, in.PageSize)
	if err != nil {
		logger.Error("service user: %s", err.Error())
		return errors.Wrap(err, "service user")
	}

	out = &v1.ListRespond{
		Code:  errno.Success.Code,
		Count: int32(count),
	}

	for k, user := range users {
		out.All[int32(k)] = &v1.One{
			Username: user.Username,
			Age:      user.Age,
			Addr:     user.Address,
		}
	}

	return nil
}
