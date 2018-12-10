package user

import (
	"mango/business"
)

func EncodeUser(user *User) business.Map {
	mapUser := business.Map{
		"id": user.Id,
		"username": user.Username,
		"avatar": user.Avatar,
		"created_at": user.CreatedAt,
	}
	if user.Id !=0 {
		return mapUser
	} else {
		return nil
	}
}