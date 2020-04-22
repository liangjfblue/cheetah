package controllers

import (
	"github.com/liangjfblue/cheetah/app/interface/web/service/users"
	"github.com/liangjfblue/cheetah/common/http/handle"
)

type Users struct {
}

func UserLogin(c *gin.context) {
	var (
		result handle.Result
	)
	resp, err := users.Login()
	if err != nil {

	}

	result.Success(c, resp)
}
