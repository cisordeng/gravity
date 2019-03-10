package user

import (
	"fmt"
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
	m_user "mango/models/user"
	"mango/pandora/encrypt"
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

func NewUser(ctx *xenon.BCtx, Username string, Password string, Avatar string) (user *User) {
	model := m_user.User{
		Username: Username,
		Password: encrypt.String2MD5(Password),
		Avatar: Avatar,
	}
	Error := xenon.Error{}
	_, Error.Inner = orm.NewOrm().Insert(&model)
	Error.Business = xenon.NewBusinessError("user:create_fail", fmt.Sprintf("创建%suser失败", Username))
	ctx.Errors = append(ctx.Errors, Error)
	return InitUserFromModel(&model)
}