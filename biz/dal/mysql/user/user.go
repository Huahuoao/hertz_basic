package mysql

import (
	"github.com/huahuoao/hertz_base/biz/dal/mysql"
	"github.com/huahuoao/hertz_base/biz/model/common"
)

func CreateUser(user *common.User) error {
	return mysql.DB.Create(user).Error
}

func GetUserByUsername(username string) (*common.User, error) {
	m := &common.User{}
	err := mysql.DB.Where("user_name =?", username).First(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ListAllUsers() ([]*common.User, error) {
	var users []*common.User
	err := mysql.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
