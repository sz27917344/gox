package gox

type HttpCode struct {
	Code    string
	Message string
}

var COPY_FAILED = HttpCode{Code: "C0001", Message: "属性拷贝失败"}
var DbError = HttpCode{Code: "C0002", Message: "用户获取失败"}
var ParamError = HttpCode{Code: "C0003", Message: "参数错误"}
