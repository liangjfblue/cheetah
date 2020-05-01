/*
@Time : 2020/5/1 22:56
@Author : liangjiefan
*/
package models

import (
	"github.com/jinzhu/gorm"
)

type TBMenu struct {
	gorm.Model
	URL         string `gorm:"column:url;size:100;" description:"菜单URL"`
	Name        string `gorm:"column:name;size:32;not null;" description:"菜单名称"`
	ParentID    uint   `gorm:"column:parent_id;not null;" description:"父级ID"`
	Sequence    int    `gorm:"column:sequence;not null;" description:"排序值"`
	MenuType    uint8  `gorm:"column:menu_type;type:tinyint(1);not null;" description:"菜单类型 1模块 2菜单 3操作"`
	Code        string `gorm:"column:code;size:32;not null;unique_index:uk_menu_code;" description:"菜单代码"`
	Icon        string `gorm:"column:icon;size:32;" description:"icon"`
	OperateType string `gorm:"column:operate_type;size:32;not null;" description:"操作类型 add/delete/get/update/list"`
	IsAvailable int8   `gorm:"column:is_available;null" description:"是否可用 1-可用 0-不可用" `
	Remark      string `gorm:"column:remark;size:64;" description:"备注"`
}

func (t *TBMenu) TableName() string {
	return "tb_menu"
}

func (t *TBMenu) Create() error {
	return DB.Create(t).Error
}

func GetMenu(u *TBMenu) (*TBMenu, error) {
	var menu TBMenu
	err := DB.Where(u).First(&menu).Error
	return &menu, err
}

func ListMenus(query map[string]interface{}, orders []string, group string,
	offset int32, limit int32) (uint64, []*TBMenu, error) {
	var (
		err   error
		menus = make([]*TBMenu, 0)
		count uint64
	)

	db := DB.Model(&TBMenu{})

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
	err = db.Offset(offset).Limit(limit).Find(&menus).Error

	return count, menus, err
}

func DeleteMenu(id uint) error {
	menu := TBMenu{
		Model: gorm.Model{ID: id},
	}
	return DB.Delete(&menu).Error
}

func (t *TBMenu) Update() error {
	return DB.Save(t).Error
}

//GetMenuButton 获取菜单有权限的操作列表
func (t *TBMenu) GetMenuButton(userId uint64, menuCode string) (bs *[]string, err error) {
	sql := `select operate_type from tb_menu
	      where id in (
					select menu_id from tb_role_menu where 
					menu_id in (select id from tb_menu where parent_id in (select id from tb_menu where code=?))
					and role_id in (select role_id from tb_user where user_id=?)
				)`
	err = DB.Raw(sql, menuCode, userId).Pluck("operate_type", bs).Error
	return
}

//GetMenuByUserId 获取用户权限下所有菜单
func (t *TBMenu) GetMenuByUserId(userId uint64) (menus *[]TBMenu, err error) {
	sql := `select * from tb_menu
	      where id in (
					select menu_id from tb_role_menu where 
				  	role_id in (select role_id from tb_user where user_id=?)
				)`
	err = DB.Raw(sql, userId).Find(&menus).Error
	return
}

// 删除菜单及关联数据
func (t *TBMenu) Delete(ids []uint) error {
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, id := range ids {
		if err := deleteMenuRecursive(tx, id); err != nil {
			tx.Rollback()
			return err
		}
	}

	//delete tole menu
	if err := tx.Where("menu_id in (?)", ids).Delete(&TBRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//delete role
	if err := tx.Where("id in (?)", ids).Delete(&TBMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func deleteMenuRecursive(db *gorm.DB, parentID uint) error {
	var menus []TBMenu
	if err := db.Where(&TBMenu{ParentID: parentID}).Find(&menus).Error; err != nil {
		return err
	}

	//delete parent menu and child
	for _, menu := range menus {
		if err := db.Where("menu_id = ?", menu.ID).Delete(&TBRoleMenu{}).Error; err != nil {
			return err
		}
		if err := deleteMenuRecursive(db, menu.ID); err != nil {
			return err
		}
	}

	if err := db.Where(&TBMenu{ParentID: parentID}).Delete(&TBMenu{}).Error; err != nil {
		return err
	}

	return nil
}
