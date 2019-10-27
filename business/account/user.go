package account

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "nature/model/account"

)

type User struct {
	Id int
	Username string
	Avatar string
	CreatedAt time.Time
}

func init() {
}

func InitUserFromModel(model *mUser.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.Username = model.Username
	instance.Avatar = model.Avatar
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewUser(username string, password string, avatar string) (user *User) {
	model := mUser.User{
		Username: username,
		Password: xenon.String2MD5(password),
		Avatar: avatar,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitUserFromModel(&model)
}