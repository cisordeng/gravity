package dolphin

import (
	"github.com/cisordeng/beego/xenon"
	"time"

	"nature/common/leo"
)

type Reply struct {
	Id int
	UserId int
	ResourceType string
	ResourceId int
	ReplyId int
	Content string
	CreatedAt time.Time

	User *leo.User
	Reply *Reply
}

func InitReplyFromMap(replyMap xenon.Map) *Reply {
	instance := new(Reply)
	instance.Id = int(replyMap["id"].(float64))
	user := leo.InitUserFromMap(replyMap["user"].(xenon.Map))
	instance.UserId = user.Id
	instance.User = user
	instance.ResourceType = replyMap["resource_type"].(string)
	instance.ResourceId = int(replyMap["resource_id"].(float64))
	if replyMap["reply"] != nil {
		reply := InitReplyFromMap(replyMap["reply"].(xenon.Map))
		instance.ReplyId = reply.Id
		instance.Reply = reply
	}
	instance.Content = replyMap["content"].(string)
	strCreatedAt := replyMap["created_at"].(string)
	createdAt, err := time.Parse("2006-01-02 15:04:05", strCreatedAt)
	xenon.PanicNotNilError(err)
	instance.CreatedAt = createdAt
	return instance
}
