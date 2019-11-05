package account

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mUser "nature/model/account"
)

func GetUsers(filters xenon.Map, orderExprs ...string ) []*User {
	o := orm.NewOrm()
	qs := o.QueryTable(&mUser.User{})

	var models []*mUser.User
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	users := make([]*User, 0)
	for _, model := range models {
		users = append(users, InitUserFromModel(model))
	}
	return users
}

func GetUserByName(name string) (user *User)  {
	model := mUser.User{}
	err := orm.NewOrm().QueryTable(&mUser.User{}).Filter(xenon.Map{
		"name": name,
	}).One(&model)
	xenon.PanicNotNilError(err, "raise:account:not_exits", "用户不存在")
	user = InitUserFromModel(&model)
	return user
}
