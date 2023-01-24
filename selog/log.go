package selog

import (
	"os"
	"path"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type SeLogger struct {
	selogger *zap.Logger
}

// create standard logger
func NewStdLogger(level ...string) *SeLogger {
	encoderConfig := GeneralConfig()

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var core zapcore.Core
	if level == nil {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	} else {
		switch level[0] {
		case "debug":
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
		case "info":
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
		default:
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
		}
	}

	logger = zap.New(core)

	return &SeLogger{selogger: logger}

}

func (sl *SeLogger) Debug(msg string) {
	sl.selogger.Debug(msg)
}

func (sl *SeLogger) Info(msg string) {
	sl.selogger.Info(msg)
}

func (sl *SeLogger) Warn(msg string) {
	sl.selogger.Warn(msg)
}

func (sl *SeLogger) Error(msg string) {
	sl.selogger.Error(msg)
}

func (sl *SeLogger) Fatal(msg string) {
	sl.selogger.Fatal(msg)
}

func (sl *SeLogger) Panic(msg string) {
	sl.selogger.Panic(msg)
}

func (sl *SeLogger) Debugf(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Infof(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Warnf(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Errorf(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Fatalf(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Panicf(format string, args ...interface{}) {
	sl.selogger.Sugar().Debugf(format, args...)
}

func (sl *SeLogger) Debugt(msg string) {
	callerFields := getCallerInfoForLog()
	logger.Debug(msg, callerFields...)
}

func (sl *SeLogger) Infot(msg string) {
	callerFields := getCallerInfoForLog()
	sl.selogger.Info(msg, callerFields...)
}

func (sl *SeLogger) Warnt(msg string) {
	callerFields := getCallerInfoForLog()
	sl.selogger.Warn(msg, callerFields...)
}

func (sl *SeLogger) Errort(msg string) {
	callerFields := getCallerInfoForLog()
	sl.selogger.Error(msg, callerFields...)
}

func (sl *SeLogger) Fatalt(msg string) {
	callerFields := getCallerInfoForLog()
	sl.selogger.Fatal(msg, callerFields...)
}

func (sl *SeLogger) Panict(msg string) {
	callerFields := getCallerInfoForLog()
	sl.selogger.Panic(msg, callerFields...)
}

// 用于创建子Logger
func (sl *SeLogger) Named(name string) Logger {
	sl.selogger = sl.selogger.Named(name)
	return sl
}

func (sl *SeLogger) With(fields ...zap.Field) Logger {
	return &SeLogger{selogger: sl.selogger.With(fields...)}
}

func Infoc(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}
func Debugc(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}
func Errorc(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}
func Warnc(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
	fileName := path.Base(file)
	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", fileName), zap.Int("line", line))

	return
}

func GeneralConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

func init() {
	// encoderConfig := zap.NewProductionEncoderConfig()
	// // 设置日志记录中时间的格式
	// encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// // 在日志文件中使用大写字母记录日志级别
	// encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// // 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	// encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// file, err := os.OpenFile("../logs/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// fileWriteSyncer := zapcore.AddSync(file)
	// core := zapcore.NewTee(
	// 	// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
	// 	zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	// 	zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	// )
	// logger = zap.New(core)

	logger = NewStdLogger().selogger

}
