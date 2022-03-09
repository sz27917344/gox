package xmvc

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sz27917344/gox/xerr"
	"github.com/sz27917344/gox/xres"
)

func ParseRequest[T any](ctx iris.Context, t T) T {
	err := ctx.ReadJSON(&t)
	if err != nil {
		fmt.Println("请求参数解析错误", err)
		panic(xerr.NewCodeMessage(xres.ParamError, "请求参数解析错误"))
	}
	return t
}
