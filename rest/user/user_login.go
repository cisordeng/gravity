package user

import (
	b_user "mango/business/user"
	"mango/rest"
)

type UserLogin struct {
	rest.RestResource
}

func init () {
	rest.Resources = append(rest.Resources, new(UserLogin))
}

func (this *UserLogin) Resource() string {
	return "user.user_login"
}

func (this *UserLogin) Params() map[string][]string {
	return map[string][]string{
		"PUT":  []string{
			"username",
			"password",
		},
	}
}

func (this *UserLogin) Put() {
	Username := this.GetString("username", "")
	Password := this.GetString("password", "")
	user, err, be := b_user.LoginCheck(Username, Password)
	data := b_user.EncodeUser(user)
	this.ReturnJSON(data, err, be)
}


