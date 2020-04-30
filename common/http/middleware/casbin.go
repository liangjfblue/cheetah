package middleware

import (
	"context"
	"strings"

	"github.com/liangjfblue/cheetah/common/proto"
	"github.com/micro/go-micro/client"

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

	a.userSrvClient = userv1.NewUserService(proto.UserSrvName, cli)

	return a
}

func (m *CasBin) PrivilegeMid() gin.HandlerFunc {
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
		ctx, span, err := tracer.TraceIntoContext(ctx, "PrivilegeMid")
		if err != nil {
			logger.Error(err.Error())
			result.Failure(c, errno.ErrPrivilegeIntoContext)
			c.Abort()
			return
		}
		defer span.Finish()

		username, ok := c.Get("username")
		if !ok {
			result.Failure(c, errno.ErrUserNotLogin)
			c.Abort()
			return
		}

		//sub obj act  etc: admin /v1/users/login GET
		reps, err := m.userSrvClient.PrivilegeMid(ctx, &userv1.PrivilegeMidRequest{
			Sub: username.(string),
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
			result.Failure(c, errno.ErrPrivilegeMid)
			c.Abort()
			return
		}

		c.Next()
	}
}
