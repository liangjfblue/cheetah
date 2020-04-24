package service

import (
	"context"
	"fmt"
	"time"

	"github.com/liangjfblue/cheetah/common/logger"

	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/common/auth"
	"github.com/liangjfblue/cheetah/common/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liangjfblue/cheetah/common/uuid"

	"github.com/jinzhu/gorm"
	"github.com/liangjfblue/cheetah/app/service/web/model"
	"github.com/pkg/errors"

	"github.com/liangjfblue/cheetah/common/errno"
)

type UserService struct {
}

func (s *UserService) Register(ctx context.Context, in *v1.RegisterRequest, out *v1.RegisterRespond) error {
	fmt.Println("UserService Register")
	if ctx.Err() == context.Canceled {
		return ctx.Err()
	}

	if _, err := model.GetUser(&model.TBUser{Username: in.Username}); err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
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
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	if err := user.Encrypt(); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	if err := user.Create(); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	out.Code = errno.Success.Code
	out.Uid = user.Uid

	return nil
}

func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest, out *v1.LoginRespond) error {
	fmt.Println("UserService Login")
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	var (
		err      error
		user     *model.TBUser
		tokenStr string
	)

	user, err = model.GetUser(&model.TBUser{Username: in.Username})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = auth.Compare(user.Password, in.Password); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if user.IsAvailable != 1 {
		logger.Error("web unavailable")
		return errors.Wrap(errors.New("web unavailable"), "service web")
	}

	user.LastLogin = time.Now()
	if err = user.Update(); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	tokenStr, err = token.SignToken(token.Context{Uid: user.Uid})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Code = errno.Success.Code
	out.Token = tokenStr

	return nil
}

func (s *UserService) Get(ctx context.Context, in *v1.GetRequest, out *v1.GetRespond) error {
	fmt.Println("UserService Get")
	var (
		err  error
		user *model.TBUser
	)

	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	user, err = model.GetUser(&model.TBUser{Uid: in.Uid})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out = &v1.GetRespond{
		Code:     errno.Success.Code,
		Username: user.Username,
		Age:      user.Age,
		Addr:     user.Address,
	}

	return nil
}

func (s *UserService) List(ctx context.Context, in *v1.ListRequest, out *v1.ListRespond) error {
	fmt.Println("UserService List")
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	count, users, err := model.ListUsers(in.Username, in.Page, in.PageSize)
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
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

func (s *UserService) Auth(ctx context.Context, in *v1.AuthRequest, out *v1.AuthRespond) error {
	fmt.Println("UserService Auth")
	var (
		err  error
		t    *token.Context
		user *model.TBUser
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	if t, err = token.ParseRequest(in.Token); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if t.Uid == "" {
		logger.Error("service web: uid empty")
		return errors.Wrap(errors.New("token uid is empty"), "service web")
	}

	user, err = model.GetUser(&model.TBUser{Uid: t.Uid})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Uid = user.Uid

	return nil
}
