package word

import (
	"time"
	
	"github.com/cisordeng/beego/orm"
)

type Word struct {
	Id int
	UserId int

	Content string `orm:"type(text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (o *Word) TableName() string {
	return "word_word"
}

func init() {
	orm.RegisterModel(new(Word))
}
