package router

import (
	"net/http"

	"github.com/liangjfblue/cheetah/app/interface/web/controllers"

	"github.com/liangjfblue/cheetah/common/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/liangjfblue/cheetah/app/interface/web/service"
)

type Router struct {
	G *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		G: gin.Default(),
	}
}

func (r *Router) Init() {
	r.G.Use(gin.Recovery())
	r.G.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	r.initRouter()
}

func (r *Router) initRouter() {
	gworkers := r.G.Group("/v1/workers")
	gworkers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gworker_groups := r.G.Group("/v1/worker_groups")
	gworker_groups.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gjobs := r.G.Group("/v1/jobs")
	gjobs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gschedulers := r.G.Group("/v1/schedulers")
	gschedulers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gusers := r.G.Group("/v1/users")
	gusers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{
		gusers.POST("/register", controllers.UserRegister)
		gusers.POST("/login", controllers.UserLogin)
		gusers.GET("/get", controllers.UserGet)
		gusers.GET("/list", controllers.UserList)
	}

	scheduler_logs := r.G.Group("/v1/scheduler_logs")
	scheduler_logs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}
}
