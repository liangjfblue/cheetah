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
	gWorkers := r.G.Group("/v1/workers")
	gWorkers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gWorkerGroups := r.G.Group("/v1/worker_groups")
	gWorkerGroups.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gJobs := r.G.Group("/v1/jobs")
	gJobs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gSchedulers := r.G.Group("/v1/schedulers")
	gSchedulers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gUsers := r.G.Group("/v1/users")
	gUsers.Use(middleware.OpenTracingMid())
	{
		gUsers.POST("/register", controllers.UserRegister)
		gUsers.POST("/login", controllers.UserLogin)

		gUsers.Use(service.AuthMid.AuthMid())
		{
			gUsers.GET("/get", controllers.UserGet)
			gUsers.GET("/list", controllers.UserList)
		}
	}

	gSchedulerLogs := r.G.Group("/v1/scheduler_logs")
	gSchedulerLogs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}
}
