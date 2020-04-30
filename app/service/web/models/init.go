package models

import (
	"fmt"
	"os"

	"github.com/liangjfblue/cheetah/common/logger"

	gormadapter "github.com/casbin/gorm-adapter"

	"github.com/casbin/casbin"
	"github.com/liangjfblue/cheetah/app/service/web/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB              *gorm.DB
	_casBinEnforcer *casbin.Enforcer
)

func CasBinInstance() *casbin.Enforcer {
	return _casBinEnforcer
}

func Init() {
	var (
		err  error
		addr string
	)

	addr = os.Getenv("CONFIG_MYSQL_ADDR")
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

	initCasbin()

	return
}

//casbin  init casbin
func initCasbin() {
	_casBinEnforcer = casbin.NewEnforcer("rbac_model.conf", gormadapter.NewAdapterByDB(DB))
	_casBinEnforcer.EnableLog(true)
	if err := _casBinEnforcer.LoadPolicy(); err != nil {
		panic(err)
	}

	//test casbin
	//add a privilege need: username	roleName	resource	operator
	res := _casBinEnforcer.AddPolicy("dev", "/v1/users/get", "GET")
	if !res {
		logger.Warn("policy is exist")
	} else {
		logger.Info("policy is not exist, not add")
	}

	res = _casBinEnforcer.AddPolicy("dev", "/v1/users/list", "GET")
	if !res {
		logger.Warn("policy is exist")
	} else {
		logger.Info("policy is not exist, not add")
	}

	_casBinEnforcer.AddRoleForUser("liangjf", "dev")
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
