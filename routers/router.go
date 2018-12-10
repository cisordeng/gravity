package routers

import (
	"github.com/astaxie/beego"
	"strings"

	"mango/rest"
	_ "mango/rest/article"
	_ "mango/rest/user"
)

func init() {
	for resource, entrance := range rest.Resources {
		beego.Notice("+resource: " + resource)
		beego.Router(strings.Replace(resource, ".", "/", -1 ), entrance)
	}
}
