package article

import (
	"github.com/cisordeng/beego/xenon"

	bArticle "nature/business/article"
)

type Article struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Article))
}

func (this *Article) Resource() string {
	return "article.article"
}

func (this *Article) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
	}
}

func (this *Article) Get() {
	id, _ := this.GetInt("id", 0)

	article := bArticle.GetArticleById(id)
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}
