package xenon

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cisordeng/beego"
)

var Resources []RestResourceInterface

type Map = map[string]interface{}
type FillOption = map[string]bool

type RestResource struct {
	beego.Controller
	bCtx Ctx
}

type Ctx struct {
	Req    *http.Request
	Errors []Error
}

type RestResourceInterface interface {
	beego.ControllerInterface
	Resource() string
	Params() map[string][]string
}

func RegisterResource(resourceInterface RestResourceInterface) {
	Resources = append(Resources, resourceInterface)
}

func (r *RestResource) Resource() string {
	return ""
}

func (r *RestResource) Params() map[string][]string {
	return nil
}

func (r *RestResource) CheckParams() {
	method := r.Ctx.Input.Method()
	app := r.AppController.(RestResourceInterface)
	method2params := app.Params()
	if method2params != nil {
		if params, ok := method2params[method]; ok {
			actualParams := r.Input()
			for _, param := range params {
				if _, ok := actualParams[param]; !ok {
					r.bCtx.Errors = append(r.bCtx.Errors, Error{NewBusinessError("rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s]", param)), errors.New("missing param")})
					return
				}
			}
		}
	}
}

func (r *RestResource) Prepare() {
	r.bCtx.Req = r.Ctx.Input.Context.Request
	r.CheckParams()
}

func (r *RestResource) GetBusinessContext() *Ctx {
	return &r.bCtx
}

func RegisterResources() {
	for _, resource := range Resources {
		beego.Info("+resource: "+resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1), resource)
	}
}
