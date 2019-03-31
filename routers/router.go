package routers

import (
	"strings"

	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/plugins/cors"
	"github.com/cisordeng/beego/xenon"

	_ "gravity/rest"
)

func init() {
	for _, resource := range xenon.Resources {
		beego.Notice("+resource: " + resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1 ), resource)
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}
