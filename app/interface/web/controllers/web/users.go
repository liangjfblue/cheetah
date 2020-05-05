package web

import (
	"context"
	"strconv"

	"github.com/liangjfblue/cheetah/common/configs"
	"github.com/liangjfblue/cheetah/common/verify"

	"github.com/liangjfblue/cheetah/app/interface/web/service/users"

	"github.com/liangjfblue/cheetah/common/logger"

	"github.com/gin-gonic/gin"
	"github.com/liangjfblue/cheetah/app/interface/web/models"
	"github.com/liangjfblue/cheetah/common/errno"
	"github.com/liangjfblue/cheetah/common/http/handle"
	"github.com/liangjfblue/cheetah/common/tracer"
)

func UserAdd(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserAddRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserAdd")
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

	resp, err := users.Add(ctx, &req)
	if err != nil {
		logger.Error("web web add err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserDelete(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserDeleteRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserUpdate")
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

	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	req.Id = append(req.Id, uint(i))
	resp, err := users.Delete(ctx, &req)
	if err != nil {
		logger.Error("web web role delete err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserGet(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserGetRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserGet")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	uid, ok := c.Get("uid")
	if !ok {
		logger.Error("web web err: token no uid")
		result.Failure(c, errno.ErrNoTokenUid)
		return
	}
	req.Uid = uid.(string)

	resp, err := users.Get(ctx, &req)
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserList(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserListRequest
	)

	//tracer
	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserList")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
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
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserUpdate(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserUpdateRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}
	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserUpdate")
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

	id := c.Query("id")
	req.Id, _ = strconv.Atoi(id)
	resp, err := users.Update(ctx, &req)
	if err != nil {
		logger.Error("web web update err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserLogin(c *gin.Context) {
	var (
		err    error
		result handle.Result
		req    models.UserLoginRequest
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}

	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserLogin")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	if err = c.BindJSON(&req); err != nil {
		logger.Error("web web Login err: %s", err.Error())
		result.Failure(c, errno.ErrBind)
		return
	}

	if err = verify.Validate(req); err != nil {
		logger.Error("web web Login err: %s", err.Error())
		result.Failure(c, errno.ErrParams)
		return
	}

	resp, err := users.Login(ctx, &req)
	if err != nil {
		logger.Error("web web Login err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}

func UserLogout(c *gin.Context) {

}

func UserSetRole(c *gin.Context) {
	var (
		err    error
		result handle.Result
	)

	cc, exist := c.Get(configs.TraceContext)
	if !exist {
		logger.Error("no TraceContext")
		result.Failure(c, errno.ErrTraceNoContext)
		return
	}

	ctx := cc.(context.Context)
	ctx, span, err := tracer.TraceIntoContext(ctx, "WebUserSetRole")
	if err != nil {
		logger.Error("web web err: %s", err.Error())
		result.Failure(c, errno.ErrTraceIntoContext)
		return
	}
	defer span.Finish()

	id := c.Query("roleId")
	roleId, _ := strconv.Atoi(id)

	resp, err := users.SetRole(ctx, &models.UserSetRoleRequest{
		UserId: 0,
		RoleId: uint(roleId),
	})
	if err != nil {
		logger.Error("web web SetRole err: %s", err.Error())
		result.Failure(c, err)
		return
	}

	result.Success(c, resp)
}
