package article

import (
	"github.com/cisordeng/beego/xenon"
	"nature/common/leo"

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
			"token",
			"title",
			"content",
		},
		"POST": []string{
			"token",
			"id",
			"title",
			"content",
		},
	}
}

func (this *Article) Get() {
	id, _ := this.GetInt("id", 0)

	article := bArticle.GetArticleById(id)
	bArticle.Fill([]*bArticle.Article{ article })
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}

func (this *Article) Put() {
	title := this.GetString("title", "")
	content := this.GetString("content", "")

	user := leo.User{}
	this.GetUserFromToken(&user)

	article := bArticle.NewArticle(user, title, content)
	bArticle.Fill([]*bArticle.Article{ article })
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}

func (this *Article) Post() {
	id, _ := this.GetInt("id", 0)
	title := this.GetString("title")
	content := this.GetString("content")

	article := bArticle.GetArticleById(id)
	article.Update(title, content)
	bArticle.Fill([]*bArticle.Article{ article })
	data := bArticle.EncodeArticle(article)
	this.ReturnJSON(data)
}
