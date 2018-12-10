package user

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	m_user "mango/models/user"
	"mango/pandora/encrypt"
	"mango/rest"
	"time"
)

type User struct {
	Id int
	Username string
	Avatar string
	CreatedAt time.Time
}

func init() {
}

func InitUserFromModel(model *m_user.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.Username = model.Username
	instance.Avatar = model.Avatar
	instance.CreatedAt = model.CreatedAt

	return instance
}

func Create(Username string, Password string, Avatar string) (user *User, err error, be rest.BusinessError) {
	model := m_user.User{
		Username: Username,
		Password: encrypt.String2MD5(Password),
		Avatar: Avatar,
	}
	_, err = orm.NewOrm().Insert(&model)
	return InitUserFromModel(&model), err, rest.BusinessError{
		ErrCode: "user:create_fail",
		ErrMsg:  fmt.Sprintf("创建%suser失败", Username),
	}
}