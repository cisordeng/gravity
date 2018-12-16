package routers

import (
	"github.com/astaxie/beego"
	"strings"

	"mango/rest"
	_ "mango/rest/article"
	_ "mango/rest/user"
)

func init() {
	for _, resource := range rest.Resources {
		beego.Notice("+resource: " + resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1 ), resource)
	}
}
