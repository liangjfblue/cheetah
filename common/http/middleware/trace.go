package middleware

import (
	"context"

	"github.com/liangjfblue/cheetah/common/comConfigs"

	"github.com/gin-gonic/gin"

	"github.com/liangjfblue/cheetah/common/tracer"
)

func OpenTracingMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span, err := tracer.TraceFromHeader(context.Background(), "api:"+c.Request.URL.Path, c.Request.Header)
		if err == nil {
			defer span.Finish()
			c.Set(comConfigs.TraceContext, ctx)
		} else {
			c.Set(comConfigs.TraceContext, context.Background())
		}

		c.Next()
	}
}
