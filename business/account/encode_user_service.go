package account

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeUser(user *User) xenon.Map {
	mapUser := xenon.Map{
		"id": user.Id,
		"username": user.Username,
		"avatar": user.Avatar,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapUser
}