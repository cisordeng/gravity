package comment

import (
	"time"
	
	"github.com/cisordeng/beego/orm"
)

type Comment struct {
	Id int
	UserId int

	ResourceType string  // 资源类型
	ResourceId int // 资源ID
	CommentId int `orm:"default(0)"` // 回复评论ID

	Content string `orm:"type(text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (o *Comment) TableName() string {
	return "comment_comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}
