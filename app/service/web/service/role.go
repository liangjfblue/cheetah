/*
@Time : 2020/5/3 21:15
@Author : liangjiefan
*/
package service

import (
	"context"
	"time"

	"github.com/jinzhu/copier"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jinzhu/gorm"
	"github.com/liangjfblue/cheetah/app/service/web/models"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/pkg/errors"

	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
)

type RoleService struct {
}

//Add 新增角色
func (r *RoleService) Add(ctx context.Context, in *v1.RoleAddRequest, out *v1.RoleAddRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleAdd.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	if _, err = models.GetRole(&models.TBRole{RoleName: in.RoleName}); err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	role := models.TBRole{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		RoleName:    in.RoleName,
		RoleDesc:    in.RoleName,
		IsAvailable: 1,
		IsAdmin:     int8(in.IsAdmin),
		IsBase:      int8(in.IsBase),
		Sequence:    uint(in.Sequence),
		ParentID:    uint(in.ParentID),
	}

	if err = role.Create(); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

//Delete 删除角色
func (r *RoleService) Delete(ctx context.Context, in *v1.RoleDeleteRequest, out *v1.RoleDeleteRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleDelete.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	if len(in.RoleIds) > 0 {
		roleIds := make([]uint, 0, len(in.RoleIds))

		for _, roleId := range in.RoleIds {
			roleIds = append(roleIds, uint(roleId))
		}

		if err != nil {
			return models.DB.Where("id in ?", roleIds).Delete(&models.TBRole{}).Error
		}
	}

	return nil
}

//Get 获取角色信息
func (r *RoleService) Get(ctx context.Context, in *v1.RoleGetRequest, out *v1.RoleGetRespond) error {
	var (
		err  = ctx.Err()
		role *models.TBRole
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleGet.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	role, err = models.GetRole(&models.TBRole{
		Model: gorm.Model{ID: uint(in.ID)}},
	)
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = copier.Copy(out, *role); err != nil {
		logger.Error("service web role get err: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	return nil
}

//List 获取角色列表
func (r *RoleService) List(ctx context.Context, in *v1.RoleListRequest, out *v1.RoleListRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleList.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	in.Page, in.PageSize = models.CheckPageSize(in.Page, in.PageSize)

	query := make(map[string]interface{})
	if in.Search != "" {
		query["role_name LIKE ? "] = "%" + in.Search + "%"
	}

	count, roles, err := models.ListRoles(
		query,
		[]string{"id desc"},
		"",
		(in.Page-1)*in.PageSize,
		in.PageSize,
		in.IsLimit)
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	out.Code = errno.Success.Code
	out.Count = int32(count)
	out.All = make([]*v1.RoleListRespond_RoleOne, 0)
	for _, role := range roles {
		out.All = append(out.All, &v1.RoleListRespond_RoleOne{
			ID:          uint32(role.ID),
			RoleName:    role.RoleName,
			RoleDesc:    role.RoleDesc,
			IsAvailable: uint32(role.IsAvailable),
			IsAdmin:     uint32(role.IsAdmin),
			IsBase:      uint32(role.IsBase),
			Sequence:    uint32(role.Sequence),
			ParentID:    uint32(role.ParentID),
		})
	}

	return nil
}

//Update 更新角色
func (r *RoleService) Update(ctx context.Context, in *v1.RoleUpdateRequest, out *v1.RoleUpdateRespond) error {
	var (
		err  = ctx.Err()
		role *models.TBRole
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleUpdate.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	role, err = models.GetRole(&models.TBRole{Model: gorm.Model{ID: uint(in.ID)}})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	if err = copier.Copy(role, *in); err != nil {
		logger.Error("service web update err: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = role.Update(); err != nil {
		logger.Error("service web update role: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

//SetMenus 设置角色菜单权限
func (r *RoleService) SetMenus(ctx context.Context, in *v1.RoleSetMenusRequest, out *v1.RoleSetMenusRespond) error {
	var (
		err = ctx.Err()
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	tx := models.DB.Begin()

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleSetMenus.Code
		} else {
			out.Code = errno.Success.Code
		}

		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = models.DB.Where("role_id = ?", in.RoleId).Delete(&models.TBRoleMenu{}).Error; err != nil {
		tx.Rollback()
		logger.Error("service web role set menus: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	for _, menuId := range in.MenuIds {
		roleMenu := &models.TBRoleMenu{
			RoleID: uint(in.RoleId),
			MenuID: uint(menuId),
		}
		if err = roleMenu.Create(); err != nil {
			tx.Rollback()
			logger.Error("service web role create menus: %s", err.Error())
			return errors.Wrap(err, " service web")

		}
	}

	tx.Commit()

	//casbin 设置角色权限
	if err = CasbinSetRolePermission(uint(in.RoleId)); err != nil {
		logger.Error("service web role CasbinSetRolePermission: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

//AllMenus 获取角色菜单权限列表
func (r *RoleService) AllMenus(ctx context.Context, in *v1.RoleAllMenusRequest, out *v1.RoleAllMenusRespond) error {
	var (
		err   = ctx.Err()
		menus []*models.TBRoleMenu
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	defer func() {
		if err != nil {
			out.Code = errno.ErrRoleAllMenus.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	_, menus, err = models.ListRoleMenus(
		map[string]interface{}{"role_id = ?": in.RoleId}, []string{"id desc"},
		"", -1, -1, false,
	)
	if err != nil {
		logger.Error("service web role list: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	out.RoleId = in.RoleId

	out.MenuIds = map[uint32]uint32{}
	for k, menu := range menus {
		out.MenuIds[uint32(k)] = uint32(menu.MenuID)
	}

	return nil
}
