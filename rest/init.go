package rest

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"mango/business"
)

var Resources []RestResourceInterface

type RestResource struct {
	beego.Controller
}

type RestResourceInterface interface {
	beego.ControllerInterface
	Resource() string
	Params() map[string][]string
}

type BusinessError struct {
	ErrCode string
	ErrMsg string
}

type Response struct {
	Code 		int32 		`json:"code"`
	Data 		interface{} `json:"data"`
	ErrCode 	string 		`json:"errCode"`
	ErrMsg 		string 		`json:"errMsg"`
	InnerErrMsg string 		`json:"innerErrMsg"`
}

func (r *RestResource) Resource() string {
	return ""
}

func (r *RestResource) Params() map[string][]string {
	return nil
}

func (r *RestResource) CheckParams () {
	method := r.Ctx.Input.Method()
	app := r.AppController.(RestResourceInterface)
	method2params := app.Params()
	if method2params != nil {
		if params, ok := method2params[method]; ok {
			actualParams := r.Input()
			for _, param := range params {
				if _, ok := actualParams[param]; !ok {
					err := errors.New("no param provided")
					be  := BusinessError{
						"rest:missing_argument",
						fmt.Sprintf("missing or invalid argument: %s", param),
					}
					r.ReturnJSON(nil, err, be)
					return
				}
			}
		}
	}
}

func (r *RestResource) Prepare() {
	r.CheckParams()
}

func CheckErr(response *Response, err error, be BusinessError) {
	if err != nil {
		*response = Response{ // 指针引用
			500,
			"",
			be.ErrCode,
			be.ErrMsg,
			err.Error(),
		}
	}
}

func MakeResponse(data business.Map, err error, be BusinessError) *Response {
	response := &Response{
		200,
		data,
		"",
		"",
		"",
	}
	CheckErr(response, err, be)
	return response
}

func (r *RestResource) ReturnJSON(data business.Map, err error, be BusinessError) {
	response := MakeResponse(data, err, be)
	r.Data["json"] = response
	r.ServeJSON()
}