package selog

import "go.uber.org/zap"

// Logger 程序日志接口, 用于适配多种第三方日志插件
type Logger interface {
	StandardLogger
	FormatLogger
	WithMetaLogger
	// RecoveryLogger

	// 用于创建子Logger
	Named(name string) Logger
	With(fields ...zap.Field) Logger
}


// StandardLogger 标准的日志打印
type StandardLogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
}

// FormatLogger 携带format的日志打印
type FormatLogger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

// WithMetaLogger 携带额外的日志meta数据
type WithMetaLogger interface {
	Debugt(msg string)
	Infot(msg string)
	Warnt(msg string)
	Errort(msg string)
	Fatalt(msg string)
	Panict(msg string)
}

// RecoveryLogger 记录Panice的日志
// type RecoveryLogger interface {
// 	Recover(msg string)
// }
