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

func (this *Article) Update(title string, content string) {
	params := xenon.Map{
		"title": title,
		"content": content,
	}

	o := orm.NewOrm()
	qs := o.QueryTable(&mArticle.Article{})
	_, err := qs.Filter(xenon.Map{
		"id": this.Id,
	}).Update(params)
	xenon.PanicNotNilError(err, "business:update failed", "update failed")
	this.Title = title
	this.Content = content
}

func InitArticleFromModel(model *mArticle.Article) *Article {
	instance := new(Article)
	instance.Id = model.Id
	instance.Title = model.Title
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewArticle(title string, content string) (article *Article) {
	model := mArticle.Article{
		Title: title,
		Content: content,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitArticleFromModel(&model)
}
