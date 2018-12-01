package routers

import (
	"mango/rest/article"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/article/article", &article.Article{})
}
