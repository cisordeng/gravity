package comment

import (
	"nature/business"
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	"nature/business/account"
	mComment "nature/model/comment"
)

type Comment struct {
	Id int
	UserId int
	ResourceType string
	ResourceId int
	CommentId int
	Content string
	CreatedAt time.Time

	User *account.User
	Comment *Comment
}

func init() {
}

func InitCommentFromModel(model *mComment.Comment) *Comment {
	instance := new(Comment)
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.ResourceType = model.ResourceType
	instance.ResourceId = model.ResourceId
	instance.CommentId = model.CommentId
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewComment(user account.User, restResource string, resource interface{}, comment *Comment, content string) *Comment {
	commentId := 0
	if comment != nil {
		commentId = comment.Id
	}
	model := mComment.Comment{
		UserId: user.Id,
		ResourceType: restResource,
		ResourceId: business.GetValueByName(resource,"Id").(int),
		CommentId: commentId,
		Content: content,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitCommentFromModel(&model)
}
