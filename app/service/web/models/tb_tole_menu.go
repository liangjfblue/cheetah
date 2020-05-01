/*
@Time : 2020/5/1 22:57
@Author : liangjiefan
*/
package models

import (
	"github.com/jinzhu/gorm"
)

//TBRoleMenu 角色菜单表
type TBRoleMenu struct {
	gorm.Model
	RoleID uint `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;"`
	MenuID uint `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;"`
}

func (t *TBRoleMenu) TableName() string {
	return "tb_role_menu"
}

func (t *TBRoleMenu) Create() error {
	return DB.Create(t).Error
}

func GetRoleMenu(u *TBRoleMenu) (*TBRoleMenu, error) {
	var roleMenu TBRoleMenu
	err := DB.Where(u).First(&roleMenu).Error
	return &roleMenu, err
}

func ListRoleMenus(query map[string]interface{}, orders []string, group string,
	offset int32, limit int32, isLimit bool) (uint64, []*TBRoleMenu, error) {
	var (
		err       error
		roleMenus = make([]*TBRoleMenu, 0)
		count     uint64
	)

	db := DB.Model(&TBRoleMenu{})

	for k, v := range query {
		db = db.Where(k, v)
	}

	for _, v := range orders {
		db = db.Order(v)
	}

	if group != "" {
		db = db.Group(group)
	}

	err = db.Count(&count).Error

	if isLimit {
		db = db.Offset(offset).Limit(limit)
	}

	err = db.Find(&roleMenus).Error

	return count, roleMenus, err
}

func DeleteRoleMenu(id uint) error {
	roleMenu := TBRoleMenu{
		Model: gorm.Model{ID: id},
	}
	return DB.Delete(&roleMenu).Error
}

func (t *TBRoleMenu) Update() error {
	return DB.Save(t).Error
}

func (t *TBRoleMenu) SetRoleMenus(roleId uint, menuIds []uint) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("id = ?", roleId).Delete(&TBRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, id := range menuIds {
		roleMenu := &TBRoleMenu{
			RoleID: roleId,
			MenuID: id,
		}
		if err := tx.Create(&roleMenu).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
