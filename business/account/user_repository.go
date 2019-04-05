package account

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "nature/model/account"
)

func GetUserByName(ctx *xenon.Ctx, Username string) (user *User)  {
	model := mUser.User{}
	err := orm.NewOrm().QueryTable(&mUser.User{}).Filter(xenon.Map{
		"username": Username,
	}).One(&model)
	xenon.RaiseError(ctx, err, xenon.NewBusinessError("raise:account:not_exits", "用户不存在"))
	user = InitUserFromModel(&model)
	return user
}
