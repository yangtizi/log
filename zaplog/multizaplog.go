package zaplog

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/yangtizi/go/ioutils"
	"github.com/yangtizi/log/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var mp = map[string]*TZaplog{}
var mutexLock *sync.RWMutex

func Map(strMap string) *TZaplog {
	mutexLock.RLock()
	v, ok := mp[strMap]
	mutexLock.RUnlock()

	if ok {
		return v
	}

	v = NewZaplog("./log/"+strMap+".log", 500, true, true, 60, "2006-01-02 15:04:05.000", DebugLevel)

	mutexLock.Lock()
	mp[strMap] = v
	mutexLock.Unlock()

	return v
}

//
func NewZaplog(strFilename string, nMaxSizeMB int, bLocalTime bool, bCompress bool, nMaxAge int, strTimeFormat string, nLevel int8) *TZaplog {
	p := &TZaplog{}
	p.init(strFilename, nMaxSizeMB, bLocalTime, bCompress, nMaxAge, strTimeFormat, nLevel)
	return p
}

//
type TZaplog struct {
	log *zap.SugaredLogger
}

func (m *TZaplog) init(strFilename string, nMaxSizeMB int, bLocalTime bool, bCompress bool, nMaxAge int, strTimeFormat string, nLevel int8) {
	// 提取路径
	strPath := ioutils.Path.GetDirectoryName(strFilename)
	// 创建目录
	os.MkdirAll(strPath, os.ModePerm)

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  strFilename,
		MaxSize:   nMaxSizeMB,
		LocalTime: bLocalTime,
		Compress:  bCompress, // 是否压缩
		MaxAge:    nMaxAge,   // 文件最多保存多少天
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(strTimeFormat)) // 时间格式
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zapcore.Level(nLevel)))
	m.log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

// Printf 为了兼容
func (m *TZaplog) Printf(template string, args ...interface{}) {
	fmt.Println(color.Green)
	fmt.Printf(template, args...)
	fmt.Println(color.Reset)
	m.log.Debugf(template, args...)
}

// Print 为了兼容
func (m *TZaplog) Print(args ...interface{}) {
	fmt.Println(color.Green)
	fmt.Println(args...)
	fmt.Println(color.Reset)
	Println(args...)
}

// Println 为了兼容
func (m *TZaplog) Println(args ...interface{}) {
	fmt.Println(color.Green)
	fmt.Println(args...)
	fmt.Println(color.Reset)
	m.log.Info(args...)
}

// Debug ()
func (m *TZaplog) Debug(args ...interface{}) {
	fmt.Println(color.Blue)
	fmt.Println(args...)
	fmt.Println(color.Reset)
	m.log.Debug(args...)
}

// Debugf ()
func (m *TZaplog) Debugf(template string, args ...interface{}) {
	fmt.Println(color.Blue)
	fmt.Printf(template, args...)
	fmt.Println(color.Reset)
	m.log.Debugf("[+] "+template, args...)
}

// Info ()
func (m *TZaplog) Info(args ...interface{}) {
	fmt.Println(color.Cyan)
	fmt.Println(args...)
	fmt.Println(color.Reset)
	m.log.Info(args...)
}

// Infof ()
func (m *TZaplog) Infof(template string, args ...interface{}) {
	fmt.Println(color.Cyan)
	fmt.Printf("[√] "+template, args...)
	fmt.Println(color.Reset)
	m.log.Infof("[√] "+template, args...)
}

// Warn ()
func (m *TZaplog) Warn(args ...interface{}) {
	fmt.Println(color.Yellow)
	fmt.Println(args...)
	fmt.Println(color.Reset)
	m.log.Warn(args...)
}

// Warnf ()
func (m *TZaplog) Warnf(template string, args ...interface{}) {
	fmt.Println(color.Yellow)
	fmt.Printf("[!] "+template, args...)
	fmt.Println(color.Reset)
	m.log.Warnf("[!] "+template, args...)
}

// Error ()
func (m *TZaplog) Error(args ...interface{}) {
	fmt.Println(args...)
	m.log.Error(args...)
}

// Errorf ()
func (m *TZaplog) Errorf(template string, args ...interface{}) {
	fmt.Println(color.Red)
	fmt.Printf("[x] "+template, args...)
	fmt.Println(color.Reset)
	m.log.Errorf("[x] "+template, args...)
}

// DPanic ()
func (m *TZaplog) DPanic(args ...interface{}) {
	fmt.Println(args...)
	m.log.DPanic(args...)
}

// DPanicf ()
func (m *TZaplog) DPanicf(template string, args ...interface{}) {
	fmt.Println(`[D] ` + color.Yellow)
	fmt.Printf(template, args...)
	fmt.Println(`[D] ` + color.Reset)
	m.log.DPanicf(`[D] `+template, args...)
}

// Panic ()
func (m *TZaplog) Panic(args ...interface{}) {
	fmt.Println(args...)
	m.log.Panic(args...)
}

// Panicf ()
func (m *TZaplog) Panicf(template string, args ...interface{}) {
	fmt.Println(color.RedBg)
	fmt.Printf(template, args...)
	fmt.Println(color.Reset)
	m.log.Panicf(`[P] `+template, args...)
}

// Fatal ()
func (m *TZaplog) Fatal(args ...interface{}) {
	fmt.Println(args...)
	m.log.Fatal(args...)
}

// Fatalf ()
func (m *TZaplog) Fatalf(template string, args ...interface{}) {
	fmt.Println(color.YellowBg)
	fmt.Printf(template, args...)
	fmt.Println(color.Reset)
	m.log.Fatalf(`[F] `+template, args...)
}

// Flush ()
func (m *TZaplog) Flush() {
	m.log.Sync()
}
