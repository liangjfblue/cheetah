package router

import (
	"net/http"

	"github.com/liangjfblue/cheetah/app/interface/web/controllers/web"

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
	initApi(r.G)
	initWeb(r.G)
}

//initApi init api
func initApi(g *gin.Engine) {
	a := g.Group("/api")

	a.Use(middleware.OpenTracingMid())

	gWorkers := a.Group("/v1/workers")
	gWorkers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gWorkerGroups := a.Group("/v1/worker_groups")
	gWorkerGroups.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gJobs := a.Group("/v1/jobs")
	gJobs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gSchedulers := a.Group("/v1/schedulers")
	gSchedulers.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}

	gSchedulerLogs := a.Group("/v1/scheduler_logs")
	gSchedulerLogs.Use(middleware.OpenTracingMid(), service.AuthMid.AuthMid())
	{

	}
}

//initWeb init web
func initWeb(g *gin.Engine) {
	w := g.Group("/web")

	w.Use(middleware.OpenTracingMid())

	gUsers := w.Group("/v1/users")
	gUsers.POST("/login", web.UserLogin)
	gUsers.POST("/logout", web.UserLogout)
	gUsers.Use(service.AuthMid.AuthMid(), service.CasBinMid.CasbinMiddleware())
	{
		gUsers.POST("", web.UserAdd)
		gUsers.DELETE("/:id", web.UserDelete)
		gUsers.GET("/:id", web.UserGet)
		gUsers.GET("", web.UserList)
		gUsers.PUT("/:id", web.UserUpdate)
		gUsers.POST("/set_role", web.UserSetRole)
	}

	gRoles := w.Group("/v1/roles")
	gRoles.Use(service.AuthMid.AuthMid(), service.CasBinMid.CasbinMiddleware())
	{
		gRoles.POST("", web.RoleAdd)
		gRoles.DELETE("/:id", web.RoleDelete)
		gRoles.GET("/:id", web.RoleGet)
		gRoles.GET("", web.RoleList)
		gRoles.PUT("/:id", web.RoleUpdate)
		gRoles.POST("/set_menus", web.RoleSetMenus)
		gRoles.GET("/:id/all_menus", web.RoleAllMenus)
	}

	gMenus := w.Group("/v1/menus")
	gMenus.Use(service.AuthMid.AuthMid(), service.CasBinMid.CasbinMiddleware())
	{
		gMenus.POST("", web.MenuAdd)
		gMenus.DELETE("/:id", web.MenuDelete)
		gMenus.GET("/menu/:id", web.MenuGet)
		gMenus.GET("", web.MenuList)
		gMenus.PUT("/:id", web.MenuUpdate)
		gMenus.GET("/buttons/:menuCode", web.MenuButtons) ///buttons/:roleId/:menucode
		gMenus.GET("/tree", web.MenuButtons)              //获取菜单树
		gMenus.GET("/buttons/:menuCode", web.MenuButtons)
	}
}
