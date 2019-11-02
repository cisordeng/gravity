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
		"PUT": []string{
			"title",
			"content",
		},
		"POST": []string{
			"id",
			"title",
			"content",
		},
	}
}

func (this *Article) Get() {
	id, _ := this.GetInt("id", 0)

	article := bArticle.GetArticleById(id)
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}

func (this *Article) Put() {
	title := this.GetString("title", "")
	content := this.GetString("content", "")

	article := bArticle.NewArticle(title, content)
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}

func (this *Article) Post() {
	id, _ := this.GetInt("id", 0)
	title := this.GetString("title")
	content := this.GetString("content")

	article := bArticle.GetArticleById(id)
	article.Update(title, content)
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}
