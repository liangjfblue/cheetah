package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/liangjfblue/cheetah/common/proto"
	"github.com/micro/go-micro/v2/client"

	"github.com/gin-gonic/gin"
	userv1 "github.com/liangjfblue/cheetah/app/service/web/proto/v1"
	"github.com/liangjfblue/cheetah/common/configs"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/liangjfblue/cheetah/common/tracer"
)

type CasBin struct {
	userSrvClient userv1.UserService
}

func NewCasBin(cli client.Client) *CasBin {
	a := new(CasBin)

	a.userSrvClient = userv1.NewUserService(proto.WebSrvName, cli)

	return a
}

func (m *CasBin) CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err    error
			result handle.Result
		)

		cc, ok := c.Get(configs.TraceContext)
		if !ok {
			logger.Error("no TraceContext")
			result.Failure(c, errno.ErrTraceNoContext)
			c.Abort()
			return
		}

		ctx := cc.(context.Context)
		ctx, span, err := tracer.TraceIntoContext(ctx, "CasbinMiddleware")
		if err != nil {
			logger.Error(err.Error())
			result.Failure(c, errno.ErrPrivilegeIntoContext)
			c.Abort()
			return
		}
		defer span.Finish()

		id, ok := c.Get("id")
		if !ok {
			result.Failure(c, errno.ErrUserNotLogin)
			c.Abort()
			return
		}

		roleId, ok := c.Get("roleId")
		if !ok {
			result.Failure(c, errno.ErrUserNotLogin)
			c.Abort()
			return
		}

		//若是超级管理员则跳过权限检查
		if fmt.Sprint(roleId) == "1" {
			c.Next()
			return
		}

		log.Println("=======casbin=======")
		log.Println("id:", id, "roleId:", roleId, c.Request.URL.Path, c.Request.Method)
		log.Println("=======casbin=======")
		//sub obj act  etc: admin /v1/users/login GET
		reps, err := m.userSrvClient.PrivilegeMiddleware(ctx, &userv1.UserPrivilegeMiddlewareRequest{
			Sub: fmt.Sprint(id),
			Obj: c.Request.URL.Path,
			Act: c.Request.Method,
		})
		if err != nil {
			logger.Error(err.Error())
			if strings.Contains(err.Error(), "too many request") {
				err = errno.ErrTooManyRequest
			} else {
				err = errno.ErrUserAuthMid
			}
			result.Failure(c, err)
			c.Abort()
			return
		}

		if reps.Code != errno.Success.Code {
			logger.Error(fmt.Sprintf("casbin code %d", reps.Code))
			result.Failure(c, errno.ErrPrivilegeMid)
			c.Abort()
			return
		}

		c.Next()
	}
}
