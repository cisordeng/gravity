package user

import (
	"github.com/cisordeng/beego/xenon"
	b_user "mango/business/user"
)

type User struct {
	xenon.RestResource
}

func init () {
	xenon.Resources = append(xenon.Resources, new(User))
}

func (this *User) Resource() string {
	return "user.user"
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
	article := b_user.GetUserByName(&this.BCtx, Username)
	data := b_user.EncodeUser(article)
	this.ReturnJSON(data)
}

func (this *User) Put() {
	Username := this.GetString("username", "")
	Password := this.GetString("password", "")
	Avatar := this.GetString("avatar", "")
	user := b_user.NewUser(&this.BCtx, Username, Password, Avatar)
	data := b_user.EncodeUser(user)
	this.ReturnJSON(data)
}