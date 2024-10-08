// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/huahuoao/hertz_base/biz/model/app/user"
	"github.com/huahuoao/hertz_base/biz/model/common"
	service "github.com/huahuoao/hertz_base/biz/service/user"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userService := service.NewUserService(ctx, c)
	resp, err := userService.UserRegister(&req)
	if err != nil {
		if err.Error() == "用户名已存在" {
			c.JSON(consts.StatusOK, common.NewResult().Error(301, err.Error()))
		} else {
			c.String(consts.StatusInternalServerError, "数据库错误: "+err.Error())
		}
		return
	}

	c.JSON(consts.StatusOK, common.NewResult().Success(resp))
}

// ListUsers .
// @router /user/list [GET]
func ListUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userService := service.NewUserService(ctx, c)
	resp, err := userService.ListUsers(&req)
	if err != nil {
		c.String(consts.StatusInternalServerError, "数据库错误: "+err.Error())
		return
	}
	c.JSON(consts.StatusOK, common.NewResult().Success(resp))
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userService := service.NewUserService(ctx, c)
	resp, err := userService.Login(&req)
	if err != nil {
		c.JSON(consts.StatusOK, common.NewResult().Error(301, "登录失败: "+err.Error()))
		return
	}
	c.JSON(consts.StatusOK, common.NewResult().Success(resp))
}
