package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	mysql "github.com/huahuoao/hertz_base/biz/dal/mysql/user"
	"github.com/huahuoao/hertz_base/biz/model/app/user"
	"github.com/huahuoao/hertz_base/biz/model/common"
	pack "github.com/huahuoao/hertz_base/biz/pack/user"
	"github.com/huahuoao/hertz_base/biz/util"
	"gorm.io/gorm"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

func (s *UserService) UserRegister(req *user.UserRegisterReq) (*user.UserRegisterResp, error) {
	var resp user.UserRegisterResp

	existUser, err := mysql.GetUserByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err // Return the error to be handled by the caller
	}
	if existUser != nil {
		return nil, fmt.Errorf("用户名已存在")
	}

	// Create the user
	err = mysql.CreateUser(&common.User{
		UserName: req.Username,
		Password: util.MD5Hash(req.Password),
	})
	if err != nil {
		return nil, err
	}

	resp.Msg = "注册成功"
	return &resp, nil
}

func (s *UserService) ListUsers(req *user.UserListReq) (*user.UserListResp, error) {
	resp := &user.UserListResp{}
	dbUsers, err := mysql.ListAllUsers()

	if err != nil {
		return resp, err
	}
	users := pack.PackUserList(dbUsers)
	resp.Users = users
	return resp, nil
}
