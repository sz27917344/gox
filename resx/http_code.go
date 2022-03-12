package resx

type HttpCode struct {
	Code    string
	Message string
}

var CopyFailed = HttpCode{Code: "00001", Message: "属性拷贝失败"}
var DbError = HttpCode{Code: "00002", Message: "数据库访问异常"}
var ParamError = HttpCode{Code: "00003", Message: "参数错误"}
var TraceIdGenerateFailed = HttpCode{Code: "00004", Message: "TraceId生成失败"}
