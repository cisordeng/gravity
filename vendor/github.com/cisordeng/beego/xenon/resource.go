package xenon

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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


func (r *RestResource) GetUserFromToken(user interface{}) {
	actualParams := r.Input()
	token := actualParams.Get("token")
	if token != "" {
		commonKey := beego.AppConfig.String("api::aesCommonKey")
		decodedToken, err := DecodeAesWithCommonKey(token, commonKey)
		PanicNotNilError(err, "rest:invalid token", fmt.Sprintf("[%s] is invalid token", token))
		err = json.Unmarshal([]byte(decodedToken), user)
		PanicNotNilError(err, "rest:invalid token", fmt.Sprintf("[%s] is invalid token", token))
	}
}

func (r *RestResource) checkValidSign() {
	var enableSign, _ = beego.AppConfig.Bool("api::enableSign")
	if !enableSign {
		return
	}
	var signSecret = beego.AppConfig.String("api::signSecret")
	var signEffectiveSeconds, err = strconv.ParseInt(beego.AppConfig.String("api::signEffectiveSeconds"), 10, 64)
	PanicNotNilError(err)

	params := []string{"sign", "timestamp"}
	actualParams := r.Input()
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
		if strings.ToLower(EncodeMD5(unencryptedStr)) != sign {
			RaiseException("rest:invalid sign", fmt.Sprintf("[%s] is invalid sign", sign))
		}
	}
	actualParams.Del("timestamp")
}

func (r *RestResource) checkParams() {
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

func (r *RestResource) checkValidToken() {
	actualParams := r.Input()
	token := actualParams.Get("token")
	user := make(map[string]interface{}, 0)
	if token != "" {
		commonKey := beego.AppConfig.String("api::aesCommonKey")
		decodedToken, err := DecodeAesWithCommonKey(token, commonKey)
		PanicNotNilError(err, "rest:invalid token", fmt.Sprintf("[%s] is invalid token", token))
		err = json.Unmarshal([]byte(decodedToken), &user)
		PanicNotNilError(err, "rest:invalid token", fmt.Sprintf("[%s] is invalid token", token))
		if id, ok := user["id"].(float64); !ok || id <= 0 {
			RaiseException("rest:invalid token", fmt.Sprintf("[%s] is invalid token", token))
		}
	}
}

func (r *RestResource) Prepare() {
	r.checkValidSign()
	r.checkParams()
	r.checkValidToken()
}

func RegisterResources() {
	for _, resource := range Resources {
		beego.Info("+resource: "+resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1), resource)
	}
}
