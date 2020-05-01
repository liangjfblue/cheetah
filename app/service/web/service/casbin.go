/*
@Time : 2020/5/2 0:28
@Author : liangjiefan
*/
package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/liangjfblue/cheetah/app/service/web/models"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
)

var (
	initOnce        sync.Once
	_casBinEnforcer *casbin.Enforcer
)

const (
	PrefixUserID = "u_"
	PrefixRoleID = "r_"
)

func InitCasBin(db *gorm.DB) error {
	initOnce.Do(func() {
		_casBinEnforcer = casbin.NewEnforcer("rbac_model.conf", gormadapter.NewAdapterByDB(db))
		_casBinEnforcer.EnableLog(true)
	})

	return allLoadPolicy()
}

func allLoadPolicy() error {
	if err := _casBinEnforcer.LoadPolicy(); err != nil {
		panic(err)
	}

	if err := roleLoadPolicy(); err != nil {
		return err
	}

	if err := userLoadPolicy(); err != nil {
		return err
	}

	return nil
}

// roleLoadPolicy 加载角色权限策略
func roleLoadPolicy() error {
	_, roles, err := models.ListRoles(nil, nil, "", -1, -1, false)
	if err != nil {
		return err
	}

	for _, role := range roles {
		_casBinEnforcer.DeleteRole(role.RoleName)

		_, menus, err := models.ListRoleMenus(map[string]interface{}{"role_id = ?": role.ID}, nil, "", -1, -1, false)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		for _, m := range menus {
			mm, err := models.GetMenu(&models.TBMenu{
				Model: gorm.Model{ID: m.MenuID},
			})
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			if mm.URL == "" || mm.OperateType == "" {
				continue
			}
			_casBinEnforcer.AddPermissionForUser(PrefixRoleID+fmt.Sprint(role.ID), mm.URL, mm.OperateType)
		}
	}

	return nil
}

//userLoadPolicy 加载用户权限策略
func userLoadPolicy() error {
	_, users, err := models.ListUsers(nil, nil, "", -1, -1, false)
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := CasbinAddRoleForUser(user.ID); err != nil {
			return err
		}
	}

	return nil
}

//CasbinDeleteRole 删除角色
func CasbinDeleteRole(roleIds []uint) {
	if _casBinEnforcer == nil {
		return
	}
	for _, id := range roleIds {
		_casBinEnforcer.DeletePermissionsForUser(PrefixRoleID + fmt.Sprint(id))
		_casBinEnforcer.DeleteRole(PrefixRoleID + fmt.Sprint(id))
	}
}

//CasbinSetRolePermission 设置角色权限
func CasbinSetRolePermission(roleId uint) {
	if _casBinEnforcer == nil {
		return
	}
	_casBinEnforcer.DeletePermissionsForUser(PrefixRoleID + fmt.Sprint(roleId))
	setRolePermission(_casBinEnforcer, roleId)
}

//setRolePermission 设置角色权限
func setRolePermission(enforcer *casbin.Enforcer, roleId uint) {
	_, roleMenus, err := models.ListRoleMenus(
		map[string]interface{}{"role_id = ?": roleId},
		nil, "", -1, -1, false)
	if err != nil {
		return
	}

	for _, m := range roleMenus {
		menu, err := models.GetMenu(&models.TBMenu{Model: gorm.Model{ID: m.MenuID}})
		if err != nil {
			return
		}
		if menu.MenuType == 3 {
			enforcer.AddPermissionForUser(PrefixRoleID+fmt.Sprint(roleId), menu.URL, "GET|POST|DELETE|PUT")
		}
	}
}

//CasbinCheckPermission 检查用户是否有权限
func CasbinCheckPermission(userID, url, methodtype string) (bool, error) {
	if _casBinEnforcer == nil {
		return false, errors.New("casBinEnforcer is null")
	}
	return _casBinEnforcer.EnforceSafe(PrefixUserID+userID, url, methodtype)
}

//CasbinAddRoleForUser 用户角色权限处理 为用户分配权限和用户登录时调用
func CasbinAddRoleForUser(userId uint) (err error) {
	if _casBinEnforcer == nil {
		return
	}

	uid := PrefixUserID + fmt.Sprint(userId)
	_casBinEnforcer.DeleteRolesForUser(uid)

	user, err := models.GetUser(&models.TBUser{Model: gorm.Model{ID: userId}})
	if err != nil {
		return
	}

	var userRoles []*models.TBUser
	userRoles = append(userRoles, user)

	for _, user := range userRoles {
		_casBinEnforcer.AddRoleForUser(uid, PrefixRoleID+fmt.Sprint(user.RoleId))
	}
	return
}
