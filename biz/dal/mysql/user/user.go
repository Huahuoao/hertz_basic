package mysql

import (
	"github.com/huahuoao/hertz_base/biz/dal/mysql"
	"github.com/huahuoao/hertz_base/biz/model/do"
)

func CreateUser(user *do.User) error {
	return mysql.DB.Create(user).Error
}
