package user

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
	m_user "gravity/models/account"
	"time"
)

type User struct {
	Id int
	Username string
	Avatar string
	CreatedAt time.Time
}

func init() {
}

func InitUserFromModel(model *m_user.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.Username = model.Username
	instance.Avatar = model.Avatar
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewUser(ctx *xenon.Ctx, Username string, Password string, Avatar string) (user *User) {
	model := m_user.User{
		Username: Username,
		Password: xenon.String2MD5(Password),
		Avatar: Avatar,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.RaiseError(ctx, err)
	return InitUserFromModel(&model)
}