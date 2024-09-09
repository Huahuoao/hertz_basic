package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	mysql "github.com/huahuoao/hertz_base/biz/dal/mysql/user"
	"github.com/huahuoao/hertz_base/biz/model/app/user"
	consts "github.com/huahuoao/hertz_base/biz/model/const"
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
	err = mysql.CreateUser(&mysql.User{
		UserName:  req.Username,
		Password:  util.MD5Hash(req.Password),
		Gender:    consts.DefaultGender,
		Age:       consts.DefaultAge,
		Introduce: consts.DefaultIntroduce,
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

func (s *UserService) Login(req *user.UserLoginReq) (*user.UserLoginResp, error) {
	dbUser, err := mysql.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, fmt.Errorf("用户名不存在")
	}
	if dbUser.Password != util.MD5Hash(req.Password) {
		return nil, fmt.Errorf("密码错误")
	}
	resp := &user.UserLoginResp{}
	resp.Token = "token"
	return resp, nil
}
