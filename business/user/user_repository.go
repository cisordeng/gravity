package user

import (
	"errors"
	"github.com/cisordeng/beego/orm"
	m_user "mango/models/user"
	"mango/pandora/encrypt"
	"mango/rest"
)

func GetUserByName(Username string) (user *User, err error, be rest.BusinessError)  {
	model := m_user.User{}
	err = orm.NewOrm().QueryTable("user_user").Filter("Username", Username).One(&model)
	user = InitUserFromModel(&model)
	return user, err, rest.BusinessError{
		ErrCode: "raise:user:not_exits",
		ErrMsg: "用户不存在",
	}
}

func LoginCheck(Username string, Password string) (user *User, err error, be rest.BusinessError) {
	model := m_user.User{}
	err = orm.NewOrm().QueryTable("user_user").Filter("Username", Username).One(&model)
	User := InitUserFromModel(&model)
	if err != nil || model.Password != encrypt.String2MD5(Password) {
		err = errors.New("username_or_password_incorrect")
		be = rest.BusinessError{
			ErrCode: "raise:user:username_or_password_incorrect",
			ErrMsg: "用户名 || 密码不正确",
		}
	}
	return User, err, be
}
