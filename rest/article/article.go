package article

import (
	b_article "mango/business/article"
	"mango/rest"
)

type Article struct {
	rest.RestResource
}

func (o *Article) Get() {
	articleId, _ := o.GetInt("article_id", 0)
	article := b_article.GetArticleById(articleId)
	data := b_article.EncodeArticle(article)
	response := rest.MakeResponse(
		200,
		data,
		"",
		"",
		"",
	)
	o.ReturnJSON(response)
}

func (o *Article) GetAll() {
}


func (o *Article) Put() {
	Title := o.GetString("title", "")
	Content := o.GetString("content", "")
	article := b_article.Create(Title, Content)
	data := b_article.EncodeArticle(article)
	response := rest.MakeResponse(
		200,
		data,
		"",
		"",
		"",
	)
	o.ReturnJSON(response)
}


