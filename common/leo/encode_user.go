package leo

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeUser(user *User) xenon.Map {
	if user == nil {
		return nil
	}
	mapUser := xenon.Map{
		"id": user.Id,
		"name": user.Name,
		"avatar": user.Avatar,
		"type": user.Type,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapUser
}