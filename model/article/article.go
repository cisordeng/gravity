package article

import (
	"time"
	
	"github.com/cisordeng/beego/orm"
)

type Article struct {
	Id int
	Title string
	Content string `orm:"type(string-text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *Article) TableName() string {
	return "article_article"
}

func init() {
	orm.RegisterModel(new(Article))
}
