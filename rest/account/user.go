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
		"GET":  []string{
			"username",
		},
		"PUT": []string{
			"username",
			"password",
			"avatar",
		},
	}
}

func (this *User) Get() {
	username := this.GetString("username", "")

	user := bUser.GetUserByName(username)
	data := bUser.EncodeUser(user)
	this.ReturnJSON(data)
}

func (this *User) Put() {
	username := this.GetString("username", "")
	password := this.GetString("password", "")
	avatar := this.GetString("avatar", "")

	user := bUser.NewUser(username, password, avatar)
	data := bUser.EncodeUser(user)
	this.ReturnJSON(data)
}