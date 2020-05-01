package service

import (
	"context"
	"time"

	"github.com/liangjfblue/cheetah/common/utils"

	"github.com/liangjfblue/cheetah/common/logger"

	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"

	"github.com/liangjfblue/cheetah/common/auth"
	"github.com/liangjfblue/cheetah/common/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liangjfblue/cheetah/common/uuid"

	"github.com/jinzhu/gorm"
	"github.com/liangjfblue/cheetah/app/service/web/models"
	"github.com/pkg/errors"

	"github.com/liangjfblue/cheetah/common/errno"
)

type UserService struct {
}

func (s *UserService) Register(ctx context.Context, in *v1.RegisterRequest, out *v1.RegisterRespond) error {
	if ctx.Err() == context.Canceled {
		return ctx.Err()
	}

	if _, err := models.GetUser(&models.TBUser{Username: in.Username}); err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	user := models.TBUser{
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
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	var (
		err      error
		user     *models.TBUser
		tokenStr string
	)

	user, err = models.GetUser(&models.TBUser{Username: in.Username})
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

	role, err := models.GetRole(&models.TBRole{
		Model: gorm.Model{ID: uint(user.RoleId)},
	})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	tokenStr, err = token.SignToken(token.Context{
		Uid:      user.Uid,
		Username: user.Username,
		RoleId:   user.RoleId,
		RoleName: role.RoleName,
		IsAdmin:  utils.Int2Bool(role.IsAdmin),
	})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Code = errno.Success.Code
	out.Token = tokenStr
	return nil
}

func (s *UserService) Get(ctx context.Context, in *v1.GetRequest, out *v1.GetRespond) error {
	var (
		err  error
		user *models.TBUser
	)

	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	user, err = models.GetUser(&models.TBUser{Uid: in.Uid})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Code = errno.Success.Code
	out.Username = user.Username
	out.Age = user.Age
	out.Addr = user.Address

	return nil
}

func (s *UserService) List(ctx context.Context, in *v1.ListRequest, out *v1.ListRespond) error {
	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	in.Page, in.PageSize = models.CheckPageSize(in.Page, in.PageSize)

	query := make(map[string]interface{})
	if in.Username != "" {
		query["username LIKE ? "] = "%" + in.Username + "%"
	}

	count, users, err := models.ListUsers(
		query,
		nil,
		"",
		(in.Page-1)*in.PageSize,
		in.PageSize,
		true)
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out = &v1.ListRespond{
		Code:  errno.Success.Code,
		Count: int32(count),
	}

	out.All = make(map[int32]*v1.One, 0)
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
	var (
		err  error
		t    *token.Context
		user *models.TBUser
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

	user, err = models.GetUser(&models.TBUser{Uid: t.Uid})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Uid = user.Uid
	out.Username = user.Username

	return nil
}

func (s *UserService) PrivilegeMid(ctx context.Context, in *v1.PrivilegeMidRequest, out *v1.PrivilegeMidRespond) error {
	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	b, err := models.CasBinInstance().EnforceSafe(in.Sub, in.Obj, in.Act)
	if err == nil && b {
		out.Code = errno.Success.Code
	} else {
		out.Code = errno.ErrPrivilege.Code
	}

	return nil
}
