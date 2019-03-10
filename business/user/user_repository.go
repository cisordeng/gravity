package user

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
	m_user "mango/models/user"
)

func GetUserByName(ctx *xenon.BCtx, Username string) (user *User)  {
	model := m_user.User{}
	Error := xenon.Error{}
	Error.Inner = orm.NewOrm().QueryTable("user_user").Filter("username", Username).One(&model)
	Error.Business = xenon.NewBusinessError("raise:user:not_exits", "用户不存在")
	ctx.Errors = append(ctx.Errors, Error)
	user = InitUserFromModel(&model)
	return user
}
