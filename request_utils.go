package gox

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func ParseRequest[T any](ctx iris.Context, t T) T {
	err := ctx.ReadJSON(&t)
	if err != nil {
		fmt.Println("请求参数解析错误", err)
		panic(NewCodeMessage(ParamError, "请求参数解析错误"))
	}
	return t
}
