package pack

import (
	mysql "github.com/huahuoao/hertz_base/biz/dal/mysql/user"
	"github.com/huahuoao/hertz_base/biz/model/common"
)

func PackUserList(dbUsers []*mysql.User) (users []*common.User) {
	for _, dbUser := range dbUsers {
		user := PackUser(dbUser)
		users = append(users, user)
	}
	return users
}

func PackUser(dbUser *mysql.User) (user *common.User) {
	user = &common.User{
		Id:        int64(dbUser.ID),
		UserName:  dbUser.UserName,
		Password:  dbUser.Password,
		Gender:    mysql.UserGender(dbUser.Gender),
		Age:       dbUser.Age,
		Introduce: dbUser.Introduce,
	}
	return
}
