/*
@Time : 2020/5/2 13:29
@Author : liangjiefan
*/
package web

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/app/interface/web/service/roles"
	"github.com/liangjfblue/cheetah/common/configs"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/liangjfblue/cheetah/common/tracer"
)

//RoleAdd 新增角色
func RoleAdd(c *gin.Context) {

}

//RoleDelete 删除角色
func RoleDelete(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RoleDeleteRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleDelete")
	if err != nil {
		logger.Error("web web role err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		result.Failure(c, errno.ErrBind)
		return
	}

	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	req.Id = append(req.Id, uint(i))

	resp, err := roles.Delete(ctx, &req)
	if err != nil {
		logger.Error("web web role delete err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//RoleGet 获取角色信息
func RoleGet(c *gin.Context) {

}

//RoleList 获取角色列表
func RoleList(c *gin.Context) {

}

//RoleUpdate 更新角色
func RoleUpdate(c *gin.Context) {

}

//RoleSetMenus 设置角色菜单权限
func RoleSetMenus(c *gin.Context) {

}

//RoleAllMenus 获取角色菜单权限列表
func RoleAllMenus(c *gin.Context) {

}
