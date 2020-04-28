package handle

import (
	"net/http"

	"github.com/liangjfblue/cheetah/common/errno"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Result) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 1,
		Data: data,
		Msg:  "ok",
	})
}

func (r *Result) Failure(c *gin.Context, err error) {
	if e, ok := err.(*errno.Errno); ok {
		c.JSON(http.StatusOK, Result{
			Code: 0,
			Data: map[string]interface{}{
				"code": e.Code,
				"msg":  e.Msg,
			},
			Msg: "error",
		})
	} else {
		c.JSON(http.StatusOK, Result{
			Code: 0,
			Data: map[string]interface{}{
				"code": -1,
				"msg":  "system error",
			},
			Msg: "error",
		})
	}
}
