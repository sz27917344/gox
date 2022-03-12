package tracex

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/sz27917344/gox/errx"
	"github.com/sz27917344/gox/resx"
	"io"
)

// New 生成新的TraceId
func New() string {
	buf := make([]byte, 10)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(errx.NewErrorMessage(resx.TraceIdGenerateFailed, err, "生成TraceId失败"))
	}
	return hex.EncodeToString(buf)
}
