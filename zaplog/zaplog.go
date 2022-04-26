package zaplog

import (
	"os"
	"path/filepath"
	"time"
)

// import log "github.com/yangtizi/go/log/zaplog"

// theZap 新的日志库,据说性能更好
var Ins *TZaplog

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel int8 = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// 不要扩展名的文件
func noExt(path string) string {
	path = filepath.Base(path)
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[0:i]
		}
	}
	return path
}

func init() {
	if Ins != nil {
		return
	}
	strNoExt := noExt(os.Args[0])
	Ins = Map(strNoExt)
}

// Since 给 defer 用的函数 defer zaplog.Since(time.Now())
func Since(t time.Time) time.Duration {
	return time.Since(t)
}
