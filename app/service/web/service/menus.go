/*
@Time : 2020/5/3 21:15
@Author : liangjiefan
*/
package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liangjfblue/cheetah/app/service/web/models"
	v1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
)

type MenuService struct {
}

//Add 新增菜单
func (m *MenuService) Add(ctx context.Context, in *v1.MenuAddRequest, out *v1.MenuAddRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuAdd.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	if _, err = models.GetMenu(&models.TBMenu{MenuCode: in.MenuCode}); err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	var menu models.TBMenu
	if err = copier.Copy(&menu, *in); err != nil {
		logger.Error("service web: %s", err.Error())
		return errno.ErrCopy
	}
	menu.IsAvailable = 1

	if err = menu.Create(); err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

//Delete 删除菜单
func (m *MenuService) Delete(ctx context.Context, in *v1.MenuDeleteRequest, out *v1.MenuDeleteRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuDelete.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	if len(in.MenuIds) > 0 {
		menuIds := make([]uint, 0, len(in.MenuIds))

		for _, menu := range in.MenuIds {
			menuIds = append(menuIds, uint(menu))
		}

		if err = models.TBMenuDelete(menuIds); err != nil {
			logger.Error("service web: %s", err.Error())
			return errors.Wrap(err, "service web")
		}

		if err = InitCasBin(models.DB); err != nil {
			logger.Error("service web: %s", err.Error())
			return errors.Wrap(err, "service web")
		}
	}

	return nil
}

//Get 获取单个菜单信息
func (m *MenuService) Get(ctx context.Context, in *v1.MenuGetRequest, out *v1.MenuGetRespond) error {
	var (
		err  = ctx.Err()
		menu *models.TBMenu
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuGet.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	menu, err = models.GetMenu(&models.TBMenu{
		Model: gorm.Model{ID: uint(in.ID)}},
	)
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = copier.Copy(out, *menu); err != nil {
		logger.Error("service web menu get err: %s", err.Error())
		return errors.Wrap(err, "service web")
	}
	out.CreateTime = menu.Model.CreatedAt.String()
	out.UpdateTime = menu.Model.UpdatedAt.String()

	return nil
}

//List 所有菜单
func (m *MenuService) List(ctx context.Context, in *v1.MenuListRequest, out *v1.MenuListRespond) error {
	var (
		err = ctx.Err()
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuList.Code
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
		query["name LIKE ? "] = "%" + in.Search + "%"
	}

	count, menus, err := models.ListMenus(
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
	out.All = make([]*v1.MenuListRespond_MenuOne, 0)
	for _, menu := range menus {
		out.All = append(out.All, &v1.MenuListRespond_MenuOne{
			ID:          uint32(menu.ID),
			URL:         menu.URL,
			Name:        menu.Name,
			ParentID:    uint32(menu.ParentID),
			Sequence:    uint32(menu.Sequence),
			MenuType:    uint32(menu.MenuType),
			MenuCode:    menu.MenuCode,
			Icon:        menu.Icon,
			OperateType: menu.OperateType,
			IsAvailable: uint32(menu.IsAvailable),
			Remark:      menu.Remark,
			CreateTime:  menu.Model.CreatedAt.String(),
			UpdateTime:  menu.Model.UpdatedAt.String(),
		})
	}

	return nil
}

//Update 更新菜单
func (m *MenuService) Update(ctx context.Context, in *v1.MenuUpdateRequest, out *v1.MenuUpdateRespond) error {
	var (
		err  = ctx.Err()
		menu *models.TBMenu
	)

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning.").Err(), "service web")
	}

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuUpdate.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	menu, err = models.GetMenu(&models.TBMenu{Model: gorm.Model{ID: uint(in.ID)}})
	if err != nil {
		logger.Error("service web: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	if err = copier.Copy(menu, *in); err != nil {
		logger.Error("service web update err: %s", err.Error())
		return errors.Wrap(err, "service web")
	}

	if err = menu.Update(); err != nil {
		logger.Error("service web update menu: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	return nil
}

//SetMenus 获取角色下的菜单ID列表
func (m *MenuService) MenuButtons(ctx context.Context, in *v1.MenuButtonsRequest, out *v1.MenuButtonsRespond) error {
	var (
		err     = ctx.Err()
		buttons []string
	)

	defer func() {
		if err != nil {
			out.Code = errno.ErrMenuMenuButtons.Code
		} else {
			out.Code = errno.Success.Code
		}
	}()

	if ctx.Err() == context.Canceled {
		logger.Error("service web: %s", ctx.Err().Error())
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service web")
	}

	user, err := models.GetUser(&models.TBUser{Model: gorm.Model{
		ID: uint(in.UserId),
	}})
	if err != nil {
		logger.Error("service web menu buttons: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	//若是管理员
	if user.Username == "admin" {
		out.OperateType = append(out.OperateType, []string{"GET", "DELETE", "POST", "PUT"}...)
		return nil
	}

	buttons, err = models.GetMenuButton(uint64(in.UserId), in.MenuCode)
	if err != nil {
		logger.Error("service web menu buttons: %s", err.Error())
		return errors.Wrap(err, " service web")
	}

	for _, button := range buttons {
		out.OperateType = append(out.OperateType, button)
	}

	return nil
}
