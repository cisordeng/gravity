package rest

import (
	"github.com/astaxie/beego"
	"mango/business"
)

type Response struct {
	Code 		int32 		`json:"code"`
	Data 		interface{} `json:"data"`
	ErrCode 	string 		`json:"errCode"`
	ErrMsg 		string 		`json:"errMsg"`
	InnerErrMsg string 		`json:"innerErrMsg"`
}

func MakeResponse(code int32, data business.Map, errCode string, errMsg string, innerErrMsg string) *Response {
	return &Response{
		code,
		data,
		errCode,
		errMsg,
		innerErrMsg,
	}
}

type RestResource struct {
	beego.Controller
}

func (r *RestResource) ReturnJSON(response *Response) {
	r.Data["json"] = response
	r.ServeJSON()
}