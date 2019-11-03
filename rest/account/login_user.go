package account

import (
	"github.com/cisordeng/beego/xenon"
	bUser "nature/business/account"
)

type LoginUser struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(LoginUser))
}

func (this *LoginUser) Resource() string {
	return "account.login_user"
}

func (this *LoginUser) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"name",
			"password",
		},
	}
}

func (this *LoginUser) Put() {
	name := this.GetString("name", "")
	password := this.GetString("password", "")
	sid := bUser.AuthUser(name, password)
	if sid != "" {
		user := bUser.GetUserByName(name)
		data := bUser.EncodeUser(user)
		data["sid"] = sid
		this.ReturnJSON(data)
	} else {
		xenon.RaiseException("rest:name or password is wrong", "用户名或密码错误")
	}
}