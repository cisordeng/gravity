package account

import (
	"github.com/cisordeng/beego/xenon"
	bUser "gravity/business/account"
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
	Username := this.GetString("username", "")

	bCtx := this.GetBusinessContext()

	article := bUser.GetUserByName(bCtx, Username)
	data := bUser.EncodeUser(article)
	this.ReturnJSON(data)
}

func (this *User) Put() {
	Username := this.GetString("username", "")
	Password := this.GetString("password", "")
	Avatar := this.GetString("avatar", "")

	bCtx := this.GetBusinessContext()

	user := bUser.NewUser(bCtx, Username, Password, Avatar)
	data := bUser.EncodeUser(user)
	this.ReturnJSON(data)
}