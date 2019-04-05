package account

import (
	"github.com/cisordeng/beego/xenon"
	b_user "gravity/business/user"
)

type User struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(User))
}

func (this *User) Resource() string {
	return "account.account"
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

	article := b_user.GetUserByName(bCtx, Username)
	data := b_user.EncodeUser(article)
	this.ReturnJSON(data)
}

func (this *User) Put() {
	Username := this.GetString("username", "")
	Password := this.GetString("password", "")
	Avatar := this.GetString("avatar", "")

	bCtx := this.GetBusinessContext()

	user := b_user.NewUser(bCtx, Username, Password, Avatar)
	data := b_user.EncodeUser(user)
	this.ReturnJSON(data)
}