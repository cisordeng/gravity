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
	app := "article"
	resource := "articles"
	rest.Resources[app + "." + resource] = new(Articles)
}

func (o *Articles) Get() {
	articles, err, be := b_article.GetArticles()
	data := make([]business.Map, 0)
	for _, article := range articles {
		data = append(data, b_article.EncodeArticle(article))
	}
	o.ReturnJSON(business.Map{
		"articles": data,
	}, err, be)
}