package routers

import (
	"github.com/astaxie/beego"
	"strings"

	"mango/rest"
	_ "mango/rest/article"
)

func init() {
	beego.Notice(rest.Resources)
	for resource, entrance := range rest.Resources {
		beego.Router(strings.Replace(resource, ".", "/", -1 ), entrance)
	}
}
