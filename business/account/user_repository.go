package account

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "nature/model/account"
)

func GetUserByName(name string) (user *User)  {
	model := mUser.User{}
	err := orm.NewOrm().QueryTable(&mUser.User{}).Filter(xenon.Map{
		"name": name,
	}).One(&model)
	xenon.PanicNotNilError(err, "raise:account:not_exits", "用户不存在")
	user = InitUserFromModel(&model)
	return user
}
