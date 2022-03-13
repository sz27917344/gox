package mvcx

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/sz27917344/gox/constx"
	"github.com/sz27917344/gox/context"
	"github.com/sz27917344/gox/errx"
	"github.com/sz27917344/gox/hackx"
	"github.com/sz27917344/gox/logx"
	"github.com/sz27917344/gox/resx"
)

func GetRequestCtx(ctx iris.Context) *context.RequestContext {
	result := ctx.Values().Get(constx.IrisKeyRequestContext)
	if result != nil {
		return result.(*context.RequestContext)
	}
	return nil
}

// ParseRequest 解析请求参数，同时完成校验
func ParseRequest[T any](ctx iris.Context, t T) T {
	request := ParseRequestOnly(ctx, t)
	// 对请求参数进行校验
	MvcValidate(request)
	return request
}

func ParseRequestOnly[T any](ctx iris.Context, t T) T {
	rawData, err := ctx.GetBody()
	if err != nil {
		panic(errx.NewErrorMessage(resx.ParamError, err, "请求参数解析错误"))
	}
	requestCtx := getRequestCtx(ctx)
	logx.Info(requestCtx, "请求参数为：%s", hackx.SliceToString(rawData))
	err = jsoniter.Unmarshal(rawData, &t)
	if err != nil {
		fmt.Println("请求参数解析错误", err)
		panic(errx.NewErrorMessage(resx.ParamError, err, "请求参数解析错误"))
	}
	return t
}

func getRequestCtx(ctx iris.Context) *context.RequestContext {
	result := ctx.Values().Get(constx.IrisKeyRequestContext)
	if result != nil {
		return result.(*context.RequestContext)
	}
	return nil
}
