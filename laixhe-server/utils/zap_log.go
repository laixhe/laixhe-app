package utils

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

// 初始 zap 日志
// path      日志文件路径
// logLevel  日志级别
// maxSize   每个日志文件保存大小
// maxBackup 保留N个备份
// maxAge    保留N天
func ZapLogInit(path, logLevel string, maxSize, maxBackup, maxAge int) {

	// 日志分割
	hook := lumberjack.Logger{
		Filename:   path,      // 日志文件路径，默认 os.TempDir()
		MaxSize:    maxSize,   // 每个日志文件保存大小，默认 100M
		MaxBackups: maxBackup, // 保留N个备份，默认不限
		MaxAge:     maxAge,    // 保留N天，默认不限
		Compress:   true,      // 是否压缩，默认不压缩
	}

	write := zapcore.AddSync(&hook)

	// 设置日志级别
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line_num",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		write,
		level,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	zapLogger = zap.New(core, caller, development)

}

// 日志时间格式
func zapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func ZapSugar() *zap.SugaredLogger {
	return zapLogger.Sugar()
}

func Zap() *zap.Logger {
	return zapLogger
}
