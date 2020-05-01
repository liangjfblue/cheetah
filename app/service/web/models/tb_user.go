package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/liangjfblue/cheetah/common/auth"

	"gopkg.in/go-playground/validator.v9"
)

//TBUser 用户表
type TBUser struct {
	gorm.Model
	Uid         string    `gorm:"column:uid;type:varchar(100);unique_index" description:"uuid"`
	Username    string    `gorm:"column:username;type:varchar(100);unique_index" description:"账号"`
	Password    string    `gorm:"column:password;type:varchar(80);null" description:"密码"`
	Age         int32     `gorm:"column:age;not null" description:"年龄"`
	Address     string    `gorm:"column:address;type:varchar(250);null" description:"地址"`
	IsAvailable int8      `gorm:"column:is_available;null" description:"是否可用 1-可用 0-不可用" `
	LastLogin   time.Time `gorm:"column:last_login;type(datetime);null" description:"最后登录时间"`
	LoginIp     string    `gorm:"column:login_ip;type:varchar(20);null" description:"登录IP"`
	RoleId      uint      `gorm:"column:role_id;not null" description:"年龄"`
}

func (t *TBUser) TableName() string {
	return "tb_user"
}

func (t *TBUser) Create() error {
	return DB.Create(t).Error
}

func GetUser(u *TBUser) (*TBUser, error) {
	var user TBUser
	err := DB.Where(u).First(&user).Error
	return &user, err
}

func ListUsers(query map[string]interface{}, orders []string, group string,
	offset int32, limit int32, isLimit bool) (uint64, []*TBUser, error) {
	var (
		err   error
		users = make([]*TBUser, 0)
		count uint64
	)

	db := DB.Model(&TBUser{})

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
	err = db.Find(&users).Error

	return count, users, err
}

func DeleteUser(id uint) error {
	user := TBUser{
		Model: gorm.Model{ID: id},
	}
	return DB.Delete(&user).Error
}

func (t *TBUser) Update() error {
	return DB.Save(t).Error
}

func (t *TBUser) Encrypt() (err error) {
	t.Password, err = auth.Encrypt(t.Password)
	return
}

func (t *TBUser) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
