package model

import (
	"fmt"
	"os"

	"github.com/liangjfblue/cheetah/app/service/worker/config"

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
		addr = config.GetInstance().MysqlConf.Addr
	}
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetInstance().MysqlConf.User, config.GetInstance().MysqlConf.Password, addr, config.GetInstance().MysqlConf.Db)

	DB, err = gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(config.GetInstance().MysqlConf.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.GetInstance().MysqlConf.MaxOpenConns)

	//DB.AutoMigrate(&TBUser{})

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
