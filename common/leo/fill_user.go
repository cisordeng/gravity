package leo

import (
	"reflect"

	"github.com/cisordeng/beego/xenon"
)

func FillUser(resources interface{}) {
	userIds := make([]int, 0)
	for i := 0; i < reflect.ValueOf(resources).Len(); i ++ {
		resource := reflect.ValueOf(resources).Index(i)
		userIds = append(userIds, resource.Elem().FieldByName("UserId").Interface().(int))
	}

	users := GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*User)
	for _, user := range users {
		id2user[user.Id] = user
	}

	for i := 0; i < reflect.ValueOf(resources).Len(); i ++ {
		resource := reflect.ValueOf(resources).Index(i)
		if user, ok := id2user[resource.Elem().FieldByName("UserId").Interface().(int)]; ok {
			resource.Elem().FieldByName("User").Set(reflect.ValueOf(user))
		}
	}
	return
}