package article

import (
	"github.com/cisordeng/beego/xenon"
	b_article "mango/business/article"
)

type Articles struct {
	xenon.RestResource
}

func init () {
	xenon.Resources = append(xenon.Resources, new(Articles))
}

func (this *Articles) Resource() string {
	return "article.articles"
}

func (this *Articles) Get() {
	articles, err := b_article.GetArticles()
	this.Error = err
	data := make([]xenon.Map, 0)
	for _, article := range articles {
		data = append(data, b_article.EncodeArticle(article))
	}
	this.ReturnJSON(xenon.Map{
		"articles": data,
	})
}