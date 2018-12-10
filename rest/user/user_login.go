package user

import (
	b_user "mango/business/user"
	"mango/rest"
)

type UserLogin struct {
	rest.RestResource
}

func init () {
	resource := "user.user_login"
	rest.Resources[resource] = new(UserLogin)
}

func (o *UserLogin) Put() {
	Username := o.GetString("username", "")
	Password := o.GetString("password", "")
	user, err, be := b_user.LoginCheck(Username, Password)
	data := b_user.EncodeUser(user)
	o.ReturnJSON(data, err, be)
}


