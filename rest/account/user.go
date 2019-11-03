package account

import (
	"github.com/cisordeng/beego/xenon"
	bUser "nature/business/account"
)

type User struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(User))
}

func (this *User) Resource() string {
	return "account.user"
}

func (this *User) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"name",
			"password",
			"avatar",
		},
	}
}

func (this *User) Put() {
	name := this.GetString("name", "")
	password := this.GetString("password", "")
	avatar := this.GetString("avatar", "")

	user := bUser.NewUser(name, password, avatar)
	data := bUser.EncodeUser(user)
	this.ReturnJSON(data)
}