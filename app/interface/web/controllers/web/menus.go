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
	"github.com/liangjfblue/cheetah/app/interface/web/service/menus"
	"github.com/liangjfblue/cheetah/common/configs"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/liangjfblue/cheetah/common/tracer"
)

//MenuAdd 新增菜单
func MenuAdd(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuAddRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuAdd")
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

	resp, err := menus.Add(ctx, &req)
	if err != nil {
		logger.Error("web web menu add err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//MenuDelete 删除菜单
func MenuDelete(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuDeleteRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuDelete")
	if err != nil {
		logger.Error("web web menu err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	req.Id = append(req.Id, uint(i))

	resp, err := menus.Delete(ctx, &req)
	if err != nil {
		logger.Error("web web menu delete err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//MenuGet 获取单个菜单信息
func MenuGet(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuGetRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuGet")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	id := c.Param("id")
	roleId, _ := strconv.Atoi(id)
	req.Id = uint(roleId)

	resp, err := menus.Get(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//MenuList 获取菜单列表
func MenuList(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuListRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuList")
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

	resp, err := menus.List(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//MenuUpdate 更新菜单
func MenuUpdate(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuUpdateRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuUpdate")
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
	resp, err := menus.Update(ctx, &req)
	if err != nil {
		logger.Error("web web update err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

//MenuButtons 获取菜单menucode有权限的操作列表
func MenuButtons(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.MenuMenuButtonsRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebMenuList")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	UserId, ok := c.Get("id")
	if !ok {
		logger.Error("web web user err: not login")
		result.Failure(c, errno.ErrUserNotLogin)
		return
	}
	id := fmt.Sprint(UserId)
	userIdd, _ := strconv.Atoi(id)
	req.UserId = int32(userIdd)

	req.MenuCode = c.Param("menuCode")

	resp, err := menus.MenuButtons(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}
