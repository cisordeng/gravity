package account

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "nature/model/account"

)

type User struct {
	Id int
	Name string
	Password string
	Avatar string
	CreatedAt time.Time
}

func init() {
}

func InitUserFromModel(model *mUser.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.Name = model.Name
	instance.Password = model.Password
	instance.Avatar = model.Avatar
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewUser(name string, password string, avatar string) (user *User) {
	model := mUser.User{
		Name: name,
		Password: xenon.EncodeMD5(password),
		Avatar: avatar,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitUserFromModel(&model)
}