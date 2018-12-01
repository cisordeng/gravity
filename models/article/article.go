package article

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id		  int
	Title 	  string
	Content   string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (o *Article) TableName() string {
	return "article_article"
}

func init() {
	orm.RegisterModel(new(Article))
}