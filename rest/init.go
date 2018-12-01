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

type BusinessError struct {
	ErrCode string
	ErrMsg string
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

type RestResource struct {
	beego.Controller
}

func (r *RestResource) ReturnJSON(response *Response) {
	r.Data["json"] = response
	r.ServeJSON()
}