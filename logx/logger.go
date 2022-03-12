package logx

import (
	"fmt"
	"github.com/sz27917344/gox/constx"
	"github.com/sz27917344/gox/context"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func getMdcFields(ctx *context.RequestContext) []zap.Field {
	zapFields := make([]zap.Field, 0, 3)
	if len(ctx.TraceId) > 0 {
		zapFields = append(zapFields, zap.String(constx.MdcTraceId, ctx.TraceId))
	}
	if len(ctx.RequestURI) > 0 {
		zapFields = append(zapFields, zap.String(constx.MdcRequestURI, ctx.RequestURI))
	}
	if ctx.Duration > 0 {
		zapFields = append(zapFields, zap.Int64(constx.MdcRequestURI, ctx.Duration.Milliseconds()))
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
func Debug(ctx *context.RequestContext, message string, a ...any) {
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

func Info(ctx *context.RequestContext, message string, a ...any) {
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

func Warn(ctx *context.RequestContext, message string, a ...any) {
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

func Error(ctx *context.RequestContext, message string, a ...any) {
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
