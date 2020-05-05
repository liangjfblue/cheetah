package service

import (
	"context"
	"time"

	"github.com/jinzhu/copier"

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

func (s *UserService) Add(ctx context.Context, in *v1.UserAddRequest, out *v1.UserAddRespond) error {
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

func (s *UserService) Login(ctx context.Context, in *v1.UserLoginRequest, out *v1.UserLoginRespond) error {
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

func (s *UserService) Get(ctx context.Context, in *v1.UserGetRequest, out *v1.UserGetRespond) error {
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

func (s *UserService) List(ctx context.Context, in *v1.UserListRequest, out *v1.UserListRespond) error {
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

	out = &v1.UserListRespond{
		Code:  errno.Success.Code,
		Count: int32(count),
	}

	out.All = make(map[int32]*v1.UserOne, 0)
	for k, user := range users {
		out.All[int32(k)] = &v1.UserOne{
			Username: user.Username,
			Age:      user.Age,
			Addr:     user.Address,
		}
	}

	return nil
}

func (s *UserService) Auth(ctx context.Context, in *v1.UserAuthRequest, out *v1.UserAuthRespond) error {
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

func (s *UserService) PrivilegeMiddleware(ctx context.Context, in *v1.UserPrivilegeMiddlewareRequest, out *v1.UserPrivilegeMiddlewareRespond) error {
	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	b, err := CasbinCheckPermission(in.Sub, in.Obj, in.Act)
	if err == nil && b {
		out.Code = errno.Success.Code
	} else {
		out.Code = errno.ErrPrivilege.Code
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, in *v1.UserDeleteRequest, out *v1.UserDeleteRespond) error {
	var (
		err error
	)
	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	defer func() {
		if err != nil {
			out.Code = errno.ErrUserDelete.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if len(in.UserIds) > 0 {
		userIds := make([]uint, 0, len(in.UserIds))

		for _, userId := range in.UserIds {
			userIds = append(userIds, uint(userId))
		}

		if err != nil {
			return models.DB.Where("id in ?", userIds).Delete(&models.TBUser{}).Error
		}
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, in *v1.UserUpdateRequest, out *v1.UserUpdateRespond) error {
	var (
		err  = ctx.Err()
		user *models.TBUser
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	defer func() {
		if err != nil {
			out.Code = errno.ErrUserUpdate.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	user, err = models.GetUser(&models.TBUser{Model: gorm.Model{ID: uint(in.ID)}})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	if err = copier.Copy(user, *in); err != nil {
		logger.Error("service web update err: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = user.Update(); err != nil {
		logger.Error("service web update user: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

func (s *UserService) SetRole(ctx context.Context, in *v1.UserSetRoleRequest, out *v1.UserSetRoleRespond) error {
	var (
		err  = ctx.Err()
		user *models.TBUser
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrUserSetRole.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if err == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	user, err = models.GetUser(&models.TBUser{Model: gorm.Model{ID: uint(in.UserId)}})
	if err != nil {
		logger.Error("service web get user: %s", ctx.Err().Error())
		return errors.Wrap(err, "service web")
	}

	user.RoleId = uint(in.RoleId)
	if err = user.Update(); err != nil {
		logger.Error("service web update user: %s", ctx.Err().Error())
		return errors.Wrap(err, "service web")
	}

	roleIds := make([]uint, 0)
	roleIds = append(roleIds, uint(in.RoleId))

	if err = CasbinAddRoleForUser(uint(in.UserId), roleIds); err != nil {
		logger.Error("service web casbin user set role: %s", ctx.Err().Error())
		return errors.Wrap(err, "service web")
	}

	return nil
}
