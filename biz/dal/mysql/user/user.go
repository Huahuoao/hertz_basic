package mysql

import (
	"github.com/huahuoao/hertz_base/biz/dal/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `json:"user_name"` // 用户名
	Password  string `json:"password"`  // 密码
	Gender    int64  `json:"gender"`    // 性别
	Age       int64  `json:"age"`       // 年龄
	Introduce string `json:"introduce"` // 个人介绍
}

func UserGender(gender int64) string {
	switch gender {
	case 1:
		return "男"
	case 2:
		return "女"
	case 3:
		return "保密"
	default:
		return "未知"
	}
}

func CreateUser(user *User) error {
	return mysql.DB.Create(user).Error
}

func GetUserByUsername(username string) (*User, error) {
	m := &User{}
	err := mysql.DB.Where("user_name =?", username).First(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ListAllUsers() ([]*User, error) {
	var users []*User
	err := mysql.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
