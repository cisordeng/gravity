package article

import (
	"github.com/cisordeng/beego/xenon"
	b_article "mango/business/article"
)

type Article struct {
	xenon.RestResource
}

func init () {
	xenon.Resources = append(xenon.Resources, new(Article))
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
	}
}

func (this *Article) Get() {
	Id, _ := this.GetInt("id", 0)
	article, err := b_article.GetArticleById(Id)
	this.Error = err
	data := b_article.EncodeArticle(article)
	this.ReturnJSON(data)
}

func (this *Article) Put() {
	Title := this.GetString("title", "")
	Content := this.GetString("content", "")
	article, err := b_article.NewArticle(Title, Content)
	this.Error = err
	data := b_article.EncodeArticle(article)
	this.ReturnJSON(data)
}


