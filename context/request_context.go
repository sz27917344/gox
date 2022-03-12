package context

import (
	"github.com/sz27917344/gox/tracex"
	"time"
)

type RequestContext struct {
	TraceId    string        // 追踪日志Id
	RequestURI string        // 访问地址
	Duration   time.Duration // 时长
}

// NewRequestContext 新的RequestContext实例
func NewRequestContext() *RequestContext {
	return &RequestContext{TraceId: tracex.New()}
}
