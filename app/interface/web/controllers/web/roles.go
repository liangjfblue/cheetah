/*
@Time : 2020/5/2 13:29
@Author : liangjiefan
*/
package web

import (
	"context"
	"fmt"
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
	var (
		err    error
		result handle.Result
		req    models.RoleAddRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleAdd")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		result.Failure(c, errno.ErrBind)
		return
	}

	resp, err := roles.Add(ctx, &req)
	if err != nil {
		logger.Error("web web role add err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
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

	id := c.Param("id")
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
	var (
		err    error
		result handle.Result
		req    models.RoleGetRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleGet")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	id := c.Param("id")
	roleId, _ := strconv.Atoi(id)
	req.Id = uint(roleId)
	fmt.Println(req)
	resp, err := roles.Get(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//RoleList 获取角色列表
func RoleList(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RoleListRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleList")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	search := c.Query("search")

	req.Page, req.PageSize = CheckPage(page, pageSize)
	req.Name = search

	resp, err := roles.List(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//RoleUpdate 更新角色
func RoleUpdate(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RoleUpdateRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleUpdate")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		result.Failure(c, errno.ErrBind)
		return
	}

	id := c.Param("id")
	req.Id, _ = strconv.Atoi(id)
	resp, err := roles.Update(ctx, &req)
	if err != nil {
		logger.Error("web web update err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//RoleSetMenus 设置角色菜单权限
func RoleSetMenus(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RoleSetMenusRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleSetMenus")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		result.Failure(c, errno.ErrBind)
		return
	}

	resp, err := roles.SetMenus(ctx, &req)
	if err != nil {
		logger.Error("web web role SetMenus err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//RoleAllMenus 获取角色菜单权限列表
func RoleAllMenus(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RoleAllMenusRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebRoleAllMenus")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	id := c.Param("id")
	roleId, _ := strconv.Atoi(id)
	req.RoleId = int32(roleId)

	resp, err := roles.AllMenus(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}
