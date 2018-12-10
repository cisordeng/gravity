package user

import (
	b_user "mango/business/user"
	"mango/rest"
)

type User struct {
	rest.RestResource
}

func init () {
	resource := "user.user"
	rest.Resources[resource] = new(User)
}

func (o *User) Get() {
	Username := o.GetString("username", "")
	article, err, be := b_user.GetUserByName(Username)
	data := b_user.EncodeUser(article)
	o.ReturnJSON(data, err, be)
}

func (o *User) Put() {
	Username := o.GetString("username", "")
	Password := o.GetString("password", "")
	Avatar := o.GetString("avatar", "")
	article, err, be := b_user.Create(Username, Password, Avatar)
	data := b_user.EncodeUser(article)
	o.ReturnJSON(data, err, be)
}