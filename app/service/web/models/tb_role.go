package models

import (
	"time"
)

//TBRole  角色表
type TBRole struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	RoleName    string     `gorm:"column:role_name;type:varchar(100);unique_index" description:"角色名称"`
	RoleDesc    string     `gorm:"column:role_desc;type:varchar(100)" description:"角色描述"`
	IsAvailable int8       `gorm:"column:is_available;null" description:"是否可用 1-可用 0-不可用" `
	IsAdmin     int8       `gorm:"column:is_admin;null" description:"是否是超级管理员 1-是 0-不是" `
	IsBase      int8       `gorm:"column:is_admin;null" description:"基础角色不能删除 1-是 0-不是" `
}

func (t *TBRole) TableName() string {
	return "tb_role"
}

func (t *TBRole) Create() error {
	return DB.Create(t).Error
}

func GetRole(u *TBRole) (*TBRole, error) {
	var role TBRole
	err := DB.Where(u).First(&role).Error
	return &role, err
}

func ListRoles(query map[string]interface{}, orders []string, group string,
	offset int32, limit int32) (uint64, []*TBRole, error) {
	var (
		err   error
		roles = make([]*TBRole, 0)
		count uint64
	)

	db := DB.Model(&TBRole{})

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
	err = db.Offset(offset).Limit(limit).Find(&roles).Error

	return count, roles, err
}

func DeleteRole(id uint) error {
	role := TBRole{
		ID: id,
	}
	return DB.Delete(&role).Error
}

func (t *TBRole) Update() error {
	return DB.Save(t).Error
}
