package model

import (
	"fmt"
	"os"

	"github.com/liangjfblue/cheetah/app/service/web/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init() {
	var (
		err  error
		addr string
	)

	addr = os.Getenv("CONFIGOR_MYSQL_ADDR")
	if addr == "" {
		addr = config.ConfigInstance().MysqlConf.Addr
	}
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ConfigInstance().MysqlConf.User, config.ConfigInstance().MysqlConf.Password, addr, config.ConfigInstance().MysqlConf.Db)

	DB, err = gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(config.ConfigInstance().MysqlConf.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.ConfigInstance().MysqlConf.MaxOpenConns)

	DB.AutoMigrate(&TBUser{})

	return
}

func CheckPageSize(page, pageSize int32) (int32, int32) {
	if page < 1 {
		page = 1
	}
	if pageSize < 15 {
		pageSize = 15
	}
	return page, pageSize
}
