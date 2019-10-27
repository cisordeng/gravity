package article

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mArticle "nature/model/article"
)

type Article struct {
	Id int
	Title string
	Content string
	CreatedAt time.Time
}

func init() {
}

func InitArticleFromModel(model *mArticle.Article) *Article {
	instance := new(Article)
	instance.Id = model.Id
	instance.Title = model.Title
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewArticle() (article *Article) {
	model := mArticle.Article{

	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitArticleFromModel(&model)
}
