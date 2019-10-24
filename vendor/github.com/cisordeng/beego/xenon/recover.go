package xenon

import (
	"fmt"
	"runtime"

	"github.com/cisordeng/beego/context"
	"github.com/cisordeng/beego/logs"
)

func RecoverPanic(ctx *context.Context) {
	if err := recover(); err != nil {
		var resp Map
		if be, ok := err.(BusinessError); ok {
			resp = Map{
				"code":        531,
				"data":        "",
				"errCode":     be.ErrCode,
				"errMsg":      be.ErrMsg,
				"innerErrMsg": "",
			}
		} else {
			resp = Map{
				"code":        531,
				"data":        "",
				"errCode":     "",
				"errMsg":      err,
				"innerErrMsg": "",
			}
		}
		logs.Critical("the request url is ", ctx.Input.URL())
		logs.Critical("Handler crashed with error", err)
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Critical(fmt.Sprintf("%s:%d", file, line))
		}

		err = ctx.Output.JSON(resp, true, true)
	}
}