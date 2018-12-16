package article

import (
	b_article "mango/business/article"
	"mango/rest"
)

type Article struct {
	rest.RestResource
}

func init () {
	rest.Resources = append(rest.Resources, new(Article))
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
	article, err, be := b_article.GetArticleById(Id)
	data := b_article.EncodeArticle(article)
	this.ReturnJSON(data, err, be)
}

func (this *Article) Put() {
	Title := this.GetString("title", "")
	Content := this.GetString("content", "")
	article, err, be := b_article.Create(Title, Content)
	data := b_article.EncodeArticle(article)
	this.ReturnJSON(data, err, be)
}


