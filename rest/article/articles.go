package article

import (
	"mango/business"
	b_article "mango/business/article"
	"mango/rest"
)

type Articles struct {
	rest.RestResource
}

func init () {
	rest.Resources = append(rest.Resources, new(Articles))
}

func (this *Articles) Resource() string {
	return "article.articles"
}

func (this *Articles) Get() {
	articles, err, be := b_article.GetArticles()
	data := make([]business.Map, 0)
	for _, article := range articles {
		data = append(data, b_article.EncodeArticle(article))
	}
	this.ReturnJSON(business.Map{
		"articles": data,
	}, err, be)
}