package xenon

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"encoding/json"

	"github.com/cisordeng/beego"
)

var Resources []RestResourceInterface

type Map = map[string]interface{}
type FillOption = map[string]bool

type RestResource struct {
	beego.Controller
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

func (r *RestResource) CheckValidSign() {
	var signSecret = beego.AppConfig.String("api::signSecret")
	var signEffectiveSeconds, err = strconv.ParseInt(beego.AppConfig.String("api::signEffectiveSeconds"), 10, 64)
	PanicNotNilError(err)

	params := []string{"sign", "timestamp"}
	actualParams := r.Input()
	if len(r.Ctx.Input.RequestBody) > 0 {
		body := make(map[string]string, 0)
		err = json.Unmarshal(r.Ctx.Input.RequestBody, &body)
		PanicNotNilError(err)
		for k, v := range body {
			actualParams[k] = []string{v}
		}
	}
	for _, param := range params {
		if _, ok := actualParams[param]; !ok {
			RaiseException("rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s]", param))
		}
	}

	sign := actualParams.Get("sign")
	timestamp, err := strconv.ParseInt(actualParams.Get("timestamp"), 10, 64)
	PanicNotNilError(err, "rest:timestamp error", fmt.Sprintf("rest:timestamp error [%d]", timestamp))

	actualParams.Del("sign")
	unencryptedStr := signSecret + actualParams.Encode()
	t := time.Unix(timestamp, 0)
	if time.Now().Before(t) || time.Now().Sub(t) > time.Duration(signEffectiveSeconds * 1000000000) { // 签名有效时间15s
		RaiseException("rest:request expired", fmt.Sprintf("at [%s] request expired", sign))
	} else {
		if strings.ToLower(String2MD5(unencryptedStr)) != sign {
			RaiseException("rest:invalid sign", fmt.Sprintf("[%s] is invalid sign", sign))
		}
	}
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
					RaiseException("rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s]", param))
				}
			}
		}
	}
}

func (r *RestResource) Prepare() {
	r.CheckValidSign()
	r.CheckParams()
}

func RegisterResources() {
	for _, resource := range Resources {
		beego.Info("+resource: "+resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1), resource)
	}
}
