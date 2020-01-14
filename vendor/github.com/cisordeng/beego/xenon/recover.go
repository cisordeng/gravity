package xenon

import (
	"context"
	"fmt"
	"strings"
	"runtime"

	"github.com/cisordeng/beego"
	beegoContext "github.com/cisordeng/beego/context"
	"github.com/cisordeng/beego/logs"
)

func rollBackTx(ctx *beegoContext.Context) {
	dbUsed, _ := beego.AppConfig.Bool("db::DB_USED")
	if !dbUsed {
		return
	}
	if ctx.Input.GetData("bContext") != nil {
		if bCtx, ok := ctx.Input.GetData("bContext").(context.Context); ok {
			o := GetOrmFromContext(bCtx)
			if o != nil {
				if ctx.Input.GetData("disableTx") != nil {
					if disableTx, ok := ctx.Input.GetData("disableTx").(bool); ok {
						if !disableTx {
							err := o.Rollback()
							beego.Warn("[ORM] rollback transaction")
							if err != nil {
								beego.Error(err)
							}
						}
					}
				}
			}
		}
	}
}

func RecoverPanic(ctx *beegoContext.Context) {
	if err := recover(); err != nil {
		// 回滚
		rollBackTx(ctx)

		logs.Critical("the request url is ", ctx.Input.URL())
		logs.Error("panic: ", fmt.Sprintf("%s", err))
		innerErrMsg := ""
		msg := make([]string, 0)
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = append(msg, fmt.Sprintf("%s:%d", file, line))
		}
		msg = msg[3:len(msg) - 4]
		for _, m := range msg {
			logs.Critical(m)
		}

		var resp Map
		if e, ok := err.(Error); ok {
			innerErrMsg = e.InnerErr.Error() + ";" + strings.Join(msg, ";")
			resp = Map{
				"code":        531,
				"data":        "",
				"errCode":     e.ErrCode,
				"errMsg":      e.ErrMsg,
				"innerErrMsg": innerErrMsg,
			}
		} else {
			innerErrMsg = fmt.Sprintf("%s", err) + ";" + strings.Join(msg, ";")
			resp = Map{
				"code":        531,
				"data":        "",
				"errCode":     "unknown error",
				"errMsg":	   "未知错误",
				"innerErrMsg": innerErrMsg,
			}
		}

		err = ctx.Output.JSON(resp, true, true)
	}
}