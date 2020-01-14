package xenon

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/orm"
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
	DisableTx() bool
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

func (r *RestResource) DisableTx() bool {
	if r.Ctx.Request.Method == "GET" {
		return true
	} else {
		return false
	}
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

func (r *RestResource) GetMap(key string) Map {
	strM := r.GetString(key, "{}")
	m := Map{}
	err := json.Unmarshal([]byte(strM), &m)
	PanicNotNilError(err, "rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s](%s)", key, "map"))
	return m
}

func (r *RestResource) GetSlice(key string) []interface{} {
	strS := r.GetString(key, "")
	s := make([]interface{}, 0)
	if len(strS) == 0 {
		return s
	}
	if strS[0] != '[' {
		strs := strings.Split(r.GetString(key, ""), ",")
		for _, str := range strs {
			s = append(s, str)
		}
		return s
	}
	err := json.Unmarshal([]byte(strS), &s)
	PanicNotNilError(err, "rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s](%s)", key, "slice"))
	return s
}

func (r *RestResource) GetPage() *Paginator {
	page, _ := r.GetInt("page", 1)
	countPerPage, _ := r.GetInt("count_per_page", 10)
	return NewPaginator(page, countPerPage)
}

func (r *RestResource) GetFilters() Map {
	filters := r.GetMap("filters")
	return filters
}

func (r *RestResource) GetOrders() []string {
	orders := r.GetStrings("orders", []string{})
	return orders
}

func (r *RestResource) encodeURIComponent() string {
	replaceMap := map[string]string{
		"+": "%20",
		"%27": "'",
		"%28": "(",
		"%29": ")",
		"%21": "!",
		"%2A": "*",
	}
	temp1 := r.Input().Encode()
	temp2 := ""
	for key, value := range replaceMap {
		temp2 = strings.Replace(temp1, key, value, -1)
		temp1 = temp2
	}
	return temp1
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
	unencryptedStr := signSecret + r.encodeURIComponent()
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
			actualParams := make(map[string]interface{}, 0)
			for k, v := range r.Input() {
				actualParams[k] = v
			}
			if r.Ctx.Request.MultipartForm != nil {
				for k, v := range r.Ctx.Request.MultipartForm.File {
					actualParams[k] = v
				}
			}

			for _, param := range params {
				paramStrs := strings.Split(param, ":")
				paramCode := paramStrs[0]
				canMissParam := false
				if paramCode[0] == '?' {
					paramCode = paramCode[1:]
					canMissParam = true
				}
				if _, ok := actualParams[paramCode]; !ok {
					if !canMissParam {
						RaiseException("rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s]", paramCode))
					}
				} else {
					if len(paramStrs) > 1 {
						paramType := paramStrs[1]
						var err error = nil
						switch paramType {
						case "string":
						case "int":
							_, err = r.GetInt(paramCode)
						case "float":
							_, err = r.GetFloat(paramCode)
						case "bool":
							_, err = r.GetBool(paramCode)
						case "map":
							r.GetMap(paramCode)
						case "slice":
							r.GetSlice(paramCode)
						case "file":
							_, _, err = r.GetFile(paramCode)
						case "files":
							_, err = r.GetFiles(paramCode)
						default:
							beego.Warn(fmt.Sprintf("unset type %s", paramType))
						}
						PanicNotNilError(err, "rest:missing_argument", fmt.Sprintf("missing or invalid argument: [%s](%s)", paramCode, paramType))
					}
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

func (r *RestResource) mergeParams() {
	token := r.Ctx.GetCookie("token")
	if token != "" {
		r.Input().Set("token", token)
	}

	// merge body params
	bodyParams := make(map[string]interface{}, 0)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &bodyParams)
	if err == nil {
		for k, v := range bodyParams {
			strV := ""
			switch t := v.(type) {
			case string:
				strV = fmt.Sprintf("%s", v.(string))
			case int:
				strV = fmt.Sprintf("%d", v.(int))
			case float64:
				strV = fmt.Sprintf("%g", v.(float64))
			case Map:
				bytes, _ := json.Marshal(v.(Map))
				strV = string(bytes)
			case []interface{}:
				bytes, _ := json.Marshal(v.([]interface{}))
				strV = string(bytes)
			default:
				beego.Warn(fmt.Sprintf("unknown type %t", t))
			}
			r.Input().Set(k, strV)
		}
	}
}

func (r *RestResource) setBusinessContext() {
	dbUsed, _ := beego.AppConfig.Bool("db::DB_USED")
	if !dbUsed {
		return
	}
	bContext := context.Background()
	bContext = context.WithValue(bContext, "orm", orm.NewOrm())
	r.Ctx.Input.SetData("bContext", bContext)
}

func (r *RestResource) GetBusinessContext() context.Context {
	dbUsed, _ := beego.AppConfig.Bool("db::DB_USED")
	if !dbUsed {
		return nil
	}
	if r.Ctx.Input.GetData("bContext") == nil {
		r.setBusinessContext()
	}
	return r.Ctx.Input.GetData("bContext").(context.Context)
}

func (r *RestResource) beginTx() {
	dbUsed, _ := beego.AppConfig.Bool("db::DB_USED")
	if !dbUsed {
		return
	}
	app := r.AppController.(RestResourceInterface)

	ctx := r.GetBusinessContext()
	o := GetOrmFromContext(ctx)
	if o != nil {
		r.Ctx.Input.SetData("disableTx", app.DisableTx())
		if !app.DisableTx() {
			err := o.Begin()
			beego.Debug("[ORM] start transaction")
			if err != nil {
				beego.Error(err)
			}
		}
	}
}

func (r *RestResource) commitTx() {
	dbUsed, _ := beego.AppConfig.Bool("db::DB_USED")
	if !dbUsed {
		return
	}
	app := r.AppController.(RestResourceInterface)

	ctx := r.GetBusinessContext()
	o := GetOrmFromContext(ctx)
	if o != nil {
		if !app.DisableTx() {
			err := o.Commit()
			beego.Debug("[ORM] commit transaction")
			if err != nil {
				beego.Error(err)
			}
		}
	}
}

func (r *RestResource) Prepare() {
	r.mergeParams() // merge params
	r.checkValidSign()
	r.checkParams()
	r.checkValidToken()
	r.beginTx()
}

func (r *RestResource) Finish() {
	r.commitTx()
}

func RegisterResources() {
	for _, resource := range Resources {
		beego.Info("+resource: "+resource.Resource(), resource.Params())
		beego.Router(strings.Replace(resource.Resource(), ".", "/", -1), resource)
	}
}
