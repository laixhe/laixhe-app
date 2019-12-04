package logs

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.SugaredLogger

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
		StacktraceKey:  "stack",
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
	callerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	logger := zap.New(core, caller, callerSkip, development)

	zapLogger = logger.Sugar()

}

// 日志时间格式
func zapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func Debug(args ...interface{}) {
	zapLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zapLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zapLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zapLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLogger.Errorf(template, args...)
}
