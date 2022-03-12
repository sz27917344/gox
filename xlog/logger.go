package xlog

import (
	"fmt"
	"github.com/sz27917344/gox/xconst"
	"github.com/sz27917344/gox/xctx"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func getMdcFields(ctx *xctx.RequestContext) []zap.Field {
	zapFields := make([]zap.Field, 0, 3)
	if len(ctx.TraceId) > 0 {
		zapFields = append(zapFields, zap.String(xconst.MdcTraceId, ctx.TraceId))
	}
	if len(ctx.RequestURI) > 0 {
		zapFields = append(zapFields, zap.String(xconst.MdcRequestURI, ctx.RequestURI))
	}
	if ctx.Duration > 0 {
		zapFields = append(zapFields, zap.Int64(xconst.MdcRequestURI, ctx.Duration.Milliseconds()))
	}
	return zapFields
}

func getFormattedMessage(message string, a ...any) string {
	if len(a) == 0 {
		return message
	} else {
		return fmt.Sprintf(message, a...)
	}
}

// Debug debug日志输出
func Debug(ctx *xctx.RequestContext, message string, a ...any) {
	text := getFormattedMessage(message, a...)
	if ctx != nil {
		zapLogger.Debug(text, getMdcFields(ctx)...)
	} else {
		zapLogger.Debug(text)
	}
}

func DebugWithoutCtx(message string, a ...any) {
	Debug(nil, message, a...)
}

func Info(ctx *xctx.RequestContext, message string, a ...any) {
	text := getFormattedMessage(message, a...)
	if ctx != nil {
		zapLogger.Info(text, getMdcFields(ctx)...)
	} else {
		zapLogger.Info(text)
	}
}

func InfoWithoutCtx(message string, a ...any) {
	Info(nil, message, a...)
}

func Warn(ctx *xctx.RequestContext, message string, a ...any) {
	text := getFormattedMessage(message, a...)
	if ctx != nil {
		zapLogger.Warn(text, getMdcFields(ctx)...)
	} else {
		zapLogger.Warn(text)
	}
}

func WarnWithoutCtx(message string, a ...any) {
	Warn(nil, message, a...)
}

func Error(ctx *xctx.RequestContext, message string, a ...any) {
	text := getFormattedMessage(message, a...)
	if ctx != nil {
		zapLogger.Error(text, getMdcFields(ctx)...)
	} else {
		zapLogger.Error(text)
	}
}

func SetZapLog(logger *zap.Logger) {
	zapLogger = logger
}

func ErrorWithoutCtx(message string, a ...any) {
	Error(nil, message, a...)
}
