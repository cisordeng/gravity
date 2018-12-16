package user

import (
	b_user "mango/business/user"
	"mango/rest"
)

type User struct {
	rest.RestResource
}

func init () {
	rest.Resources = append(rest.Resources, new(User))
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
	article, err, be := b_user.GetUserByName(Username)
	data := b_user.EncodeUser(article)
	this.ReturnJSON(data, err, be)
}

func (this *User) Put() {
	Username := this.GetString("username", "")
	Password := this.GetString("password", "")
	Avatar := this.GetString("avatar", "")
	article, err, be := b_user.Create(Username, Password, Avatar)
	data := b_user.EncodeUser(article)
	this.ReturnJSON(data, err, be)
}