package middleware

import (
	"context"
	"time"

	"github.com/liangjfblue/cheetah/common/proto"

	"github.com/liangjfblue/cheetah/common/errno"

	"github.com/liangjfblue/cheetah/common/configs"

	"github.com/gin-gonic/gin"
	userv1 "github.com/liangjfblue/cheetah/app/service/user/proto/v1"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/liangjfblue/cheetah/common/tracer"
	"github.com/micro/go-micro/client"
)

type Auth struct {
	userSrvClient userv1.UserService
}

func New() *Auth {
	a := new(Auth)

	a.userSrvClient = userv1.NewUserService(proto.UserSrvName, client.NewClient(
		client.Retries(0),
		client.DialTimeout(time.Minute*2),
	))

	return a
}

func (m *Auth) AuthMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err    error
			result handle.Result
		)

		//tracer
		cc, ok := c.Get(configs.TraceContext)
		if !ok {
			logger.Error("no TraceContext")
			result.Failure(c, errno.ErrTraceNoContext)
			c.Abort()
			return
		}

		ctx := cc.(context.Context)
		ctx, span, err := tracer.TraceIntoContext(ctx, "VerifyToken")
		if err != nil {
			logger.Error(err.Error())
			result.Failure(c, errno.ErrTraceIntoContext)
			c.Abort()
			return
		}
		defer span.Finish()

		//jwt
		token := c.Request.Header.Get("Authorization")

		req := userv1.AuthRequest{
			Token: token,
		}

		resp, err := m.userSrvClient.Auth(c, &req)
		if err != nil {
			logger.Error(err.Error())
			result.Failure(c, errno.ErrUserAuthMid)
			c.Abort()
			return
		}

		c.Set("uid", resp.UID)

		c.Next()
	}
}
