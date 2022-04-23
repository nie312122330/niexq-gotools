package logext

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoder zapcore.Encoder
var infoLevelFun zap.LevelEnablerFunc
var warnLevelFun zap.LevelEnablerFunc
var consoleLevelFun zap.LevelEnablerFunc

var OUT_LOG_LEVEL = zapcore.DebugLevel

//init 初始化
func init() {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	// 实现两个判断日志等级的
	infoLevelFun = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= OUT_LOG_LEVEL && level <= zapcore.WarnLevel
	})
	warnLevelFun = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	})
	consoleLevelFun = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= OUT_LOG_LEVEL
	})
}

//NewLogger 创建一个新的日志记录器
//  logDir 日志目录 如:./logs/
//  fileName 日志文件名称 如: main
//  maxTime 最大日志记录时间
//  rotationTime 日志轮换周期
func NewLogger(logDir string, fileName string, maxTime time.Duration, rotationTime time.Duration) *zap.Logger {
	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := GetRotationLogWriter(fmt.Sprintf("%s%s.log", logDir, fileName), maxTime, rotationTime)
	warnWriter := GetRotationLogWriter(fmt.Sprintf("%s%s_error.log", logDir, fileName), maxTime, rotationTime)
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), consoleLevelFun),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevelFun),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevelFun),
	)
	return zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
}

//DefaultLogger 默认的日志记录
//  位置为当前目录的logs文件夹,最大保留60天,轮回周期为每天
//  fileName 日志文件名称 如: main
func DefaultLogger(fileName string) *zap.Logger {
	return NewLogger("./logs/", fileName, time.Hour*24*60, time.Hour*24)
}

//LogLogger 默认的日志记录
//  位置为当前目录的logs文件夹,文件名为log.log,最大保留60天,轮回周期为每天
//  fileName 日志文件名称 如: main
func LogLogger() *zap.Logger {
	return DefaultLogger("log")
}
