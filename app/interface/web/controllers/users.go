package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/liangjfblue/cheetah/common/comConfigs"

	"github.com/pkg/errors"

	"github.com/liangjfblue/cheetah/app/interface/web/service/users"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/gin-gonic/gin"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/tracer"
)

func UserLogin(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.LoginRequest
	)

	//tracer
	cc, exist := c.Get(comConfigs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserLogin")
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		logger.Error("web user Login err: %s", err.Error())
		result.Failure(c, errno.ErrBind)
		return
	}

	if req.Username == "" || req.Password == "" {
		logger.Error("web user Login err: %s", fmt.Sprintf("params empty: %s %s", req.Username, req.Password))
		result.Failure(c, errno.ErrParams)
		return
	}

	resp, err := users.Login(ctx, &req)
	if err != nil {
		logger.Error("web user Login err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserRegister(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.RegisterRequest
	)

	//tracer
	cc, exist := c.Get(comConfigs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserRegister")
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		result.Failure(c, errno.ErrBind)
		return
	}

	resp, err := users.Register(ctx, &req)
	if err != nil {
		logger.Error("web user Register err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserGet(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.GetRequest
	)

	//tracer
	cc, exist := c.Get(comConfigs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserGet")
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	uid, ok := c.Get("uid")
	if !ok {
		logger.Error("web user err: token no uid")
		result.Failure(c, errors.New("web user err: token no uid"))
		return
	}

	req.Uid = uid.(string)

	resp, err := users.Get(ctx, &req)
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserList(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.ListRequest
	)

	//tracer
	cc, exist := c.Get(comConfigs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserList")
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")

	req.Page = int32(page)
	req.PageSize = int32(pageSize)
	req.Username = username

	resp, err := users.List(ctx, &req)
	if err != nil {
		logger.Error("web user err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}
