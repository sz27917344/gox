package xmvc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/sz27917344/gox/xconst"
	"github.com/sz27917344/gox/xctx"
	"github.com/sz27917344/gox/xerr"
	"github.com/sz27917344/gox/xhack"
	"github.com/sz27917344/gox/xlog"
	"github.com/sz27917344/gox/xres"
)

func ParseRequest[T any](ctx iris.Context, t T) T {
	rawData, err := ctx.GetBody()
	if err != nil {
		panic(xerr.NewErrorMessage(xres.ParamError, err, "请求参数解析错误"))
	}
	requestCtx := getRequestCtx(ctx)
	xlog.Info(requestCtx, "请求参数为：%s", xhack.SliceToString(rawData))
	err = jsoniter.Unmarshal(rawData, &t)
	if err != nil {
		fmt.Println("请求参数解析错误", err)
		panic(xerr.NewErrorMessage(xres.ParamError, err, "请求参数解析错误"))
	}
	return t
}

func getRequestCtx(ctx iris.Context) *xctx.RequestContext {
	result := ctx.Values().Get(xconst.IrisKeyRequestContext)
	if result != nil {
		return result.(*xctx.RequestContext)
	}
	return nil
}
