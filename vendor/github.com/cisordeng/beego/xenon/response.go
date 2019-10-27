package xenon

type Response struct {
	Code        int32       `json:"code"`
	Data        interface{} `json:"data"`
	ErrCode     string      `json:"errCode"`
	ErrMsg      string      `json:"errMsg"`
	InnerErrMsg string    	`json:"innerErrMsg"`
}

func (r *RestResource) MakeResponse(data Map) *Response {
	response := &Response{
		200,
		data,
		"",
		"",
		"",
	}
	return response
}

func (r *RestResource) ReturnJSON(data Map) {
	response := r.MakeResponse(data)
	r.Data["json"] = response
	r.ServeJSON()
}
