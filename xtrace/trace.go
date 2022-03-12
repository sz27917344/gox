package xtrace

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/sz27917344/gox/xerr"
	"github.com/sz27917344/gox/xres"
	"io"
)

// New 生成新的TraceId
func New() string {
	buf := make([]byte, 10)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(xerr.NewErrorMessage(xres.TraceIdGenerateFailed, err, "生成TraceId失败"))
	}
	return hex.EncodeToString(buf)
}
