package word

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	"nature/common/dolphin"
	"nature/common/leo"
	mWord "nature/model/word"
)

type Word struct {
	Id int
	UserId int
	Content string
	CreatedAt time.Time

	User *leo.User
	Replies []*dolphin.Reply
}

func init() {
}

func (this *Word) Update(content string) {
	params := xenon.Map{
		"content": content,
	}

	o := orm.NewOrm()
	qs := o.QueryTable(&mWord.Word{})
	_, err := qs.Filter(xenon.Map{
		"id": this.Id,
	}).Update(params)
	xenon.PanicNotNilError(err, "business:update failed", "update failed")
	this.Content = content
}

func InitWordFromModel(model *mWord.Word) *Word {
	instance := new(Word)
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewWord(user leo.User, content string) (word *Word) {
	model := mWord.Word{
		UserId: user.Id,
		Content: content,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitWordFromModel(&model)
}
