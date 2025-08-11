package logger

import (
	"context"
	"sync"
)

var (
	defaultLogger Logger
	once          sync.Once
)

// Init 初始化全局日志器
func Init(config *Config) {
	once.Do(func() {
		if config == nil {
			config = DefaultConfig()
		}
		defaultLogger = NewSlogLogger(config)
	})
}

// GetLogger 获取全局日志器
func GetLogger() Logger {
	if defaultLogger == nil {
		Init(DefaultConfig())
	}
	return defaultLogger
}

// FromContext 从 context 中获取带 requestID 的 logger
func FromContext(ctx context.Context) Logger {
	return GetLogger().WithContext(ctx)
}

// 全局日志方法，方便直接使用

// Debug 调试日志
func Debug(msg string, fields ...Fields) {
	GetLogger().Debug(msg, fields...)
}

// Info 信息日志
func Info(msg string, fields ...Fields) {
	GetLogger().Info(msg, fields...)
}

// Warn 警告日志
func Warn(msg string, fields ...Fields) {
	GetLogger().Warn(msg, fields...)
}

// Error 错误日志
func Error(msg string, fields ...Fields) {
	GetLogger().Error(msg, fields...)
}

// Fatal 致命错误日志
func Fatal(msg string, fields ...Fields) {
	GetLogger().Fatal(msg, fields...)
}

// DebugContext 带上下文的调试日志
func DebugContext(ctx context.Context, msg string, fields ...Fields) {
	GetLogger().DebugContext(ctx, msg, fields...)
}

// InfoContext 带上下文的信息日志
func InfoContext(ctx context.Context, msg string, fields ...Fields) {
	GetLogger().InfoContext(ctx, msg, fields...)
}

// WarnContext 带上下文的警告日志
func WarnContext(ctx context.Context, msg string, fields ...Fields) {
	GetLogger().WarnContext(ctx, msg, fields...)
}

// ErrorContext 带上下文的错误日志
func ErrorContext(ctx context.Context, msg string, fields ...Fields) {
	GetLogger().ErrorContext(ctx, msg, fields...)
}

// ErrorWithErr 带错误的错误日志
func ErrorWithErr(err error, msg string, fields ...Fields) {
	GetLogger().ErrorWithErr(err, msg, fields...)
}

// WarnWithErr 带错误的警告日志
func WarnWithErr(err error, msg string, fields ...Fields) {
	GetLogger().WarnWithErr(err, msg, fields...)
}

// WithFields 创建带字段的子日志器
func WithFields(fields Fields) Logger {
	return GetLogger().WithFields(fields)
}

// WithField 创建带单个字段的子日志器
func WithField(key string, value interface{}) Logger {
	return GetLogger().WithField(key, value)
}

// WithError 创建带错误的子日志器
func WithError(err error) Logger {
	return GetLogger().WithError(err)
}

// SetLevel 设置日志级别
func SetLevel(level LogLevel) {
	GetLogger().SetLevel(level)
}
