package leo

import (
	"github.com/cisordeng/beego/xenon"
)

func GetUsers(filters xenon.Map, orderExprs ...string ) []*User {
	params := xenon.Map{
		"page": 1,
		"count_per_page": 9999,
		"filters": filters,
	}
	if len(orderExprs) > 0 {
		params["orders"] = orderExprs
	}
	resp := xenon.Get("leo", "user.users", params)
	users := make([]*User, 0)
	userInterfaces := resp["data"].(xenon.Map)["users"].([]interface{})
	for _, user := range userInterfaces {
		userMap := user.(xenon.Map)
		users = append(users, InitUserFromMap(userMap))
	}
	return users
}
