package xenon

type Response struct {
	Code        int32       `json:"code"`
	Data        interface{} `json:"data"`
	ErrCode     string      `json:"errCode"`
	ErrMsg      string      `json:"errMsg"`
	InnerErrMsg string      `json:"innerErrMsg"`
}

func (r *RestResource) MakeResponse(data Map) *Response {
	response := &Response{
		200,
		data,
		"",
		"",
		"",
	}
	for _, Error := range r.bCtx.Errors {
		if Error.Inner != nil {
			response = &Response{ // 指针引用
				500,
				"",
				Error.Business.ErrCode,
				Error.Business.ErrMsg,
				Error.Inner.Error(),
			}
			return response
		}
	}
	return response
}

func (r *RestResource) ReturnJSON(data Map) {
	response := r.MakeResponse(data)
	r.Data["json"] = response
	r.ServeJSON()
}
