package leo

import (
	"github.com/cisordeng/beego/xenon"
	"time"
)

type User struct {
	Id int
	Name string
	Avatar string
	Type string
	CreatedAt time.Time
}

func InitUserFromMap(userMap xenon.Map) *User {
	instance := new(User)
	instance.Id = int(userMap["id"].(float64))
	instance.Name = userMap["name"].(string)
	instance.Avatar = userMap["avatar"].(string)
	instance.Type = userMap["type"].(string)
	strCreatedAt := userMap["created_at"].(string)
	createdAt, _ := time.Parse("2006-01-02 15:04:05", strCreatedAt)
	instance.CreatedAt = createdAt

	return instance
}