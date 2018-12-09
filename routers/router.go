package routers

import (
	"mango/rest/article"
	"strings"

	"github.com/astaxie/beego"
)

type Routers map[string]beego.ControllerInterface

func init() {
	resources := Routers{
		"article.article": new(article.Article),
		"article.articles": new(article.Articles),
	}
	for resource, entrance := range resources {
		beego.Router(strings.Replace(resource, ".", "/", -1 ), entrance)
	}
}
