package logger

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// logrusLogger logrus实现的日志器
type logrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger 创建基于logrus的日志器
func NewLogrusLogger(config *Config) Logger {
	logger := logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(string(config.Level))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// 设置日志格式
	if config.Format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// 设置输出
	var output io.Writer
	switch config.Output {
	case "file":
		// 确保日志目录存在
		if err := os.MkdirAll(filepath.Dir(config.FilePath), 0755); err != nil {
			logger.Warnf("Failed to create log directory: %v", err)
			output = os.Stdout
		} else {
			output = &lumberjack.Logger{
				Filename:   config.FilePath,
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
				Compress:   config.Compress,
			}
		}
	case "both":
		// 同时输出到文件和控制台
		if err := os.MkdirAll(filepath.Dir(config.FilePath), 0755); err != nil {
			logger.Warnf("Failed to create log directory: %v", err)
			output = os.Stdout
		} else {
			fileOutput := &lumberjack.Logger{
				Filename:   config.FilePath,
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
				Compress:   config.Compress,
			}
			output = io.MultiWriter(os.Stdout, fileOutput)
		}
	default:
		output = os.Stdout
	}
	logger.SetOutput(output)

	return &logrusLogger{logger: logger}
}

// Debug 调试日志
func (l *logrusLogger) Debug(msg string, fields ...Fields) {
	entry := l.logger.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = l.logger.WithFields(logrus.Fields(fields[0]))
	}
	entry.Debug(msg)
}

// Info 信息日志
func (l *logrusLogger) Info(msg string, fields ...Fields) {
	entry := l.logger.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = l.logger.WithFields(logrus.Fields(fields[0]))
	}
	entry.Info(msg)
}

// Warn 警告日志
func (l *logrusLogger) Warn(msg string, fields ...Fields) {
	entry := l.logger.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = l.logger.WithFields(logrus.Fields(fields[0]))
	}
	entry.Warn(msg)
}

// Error 错误日志
func (l *logrusLogger) Error(msg string, fields ...Fields) {
	entry := l.logger.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = l.logger.WithFields(logrus.Fields(fields[0]))
	}
	entry.Error(msg)
}

// Fatal 致命错误日志
func (l *logrusLogger) Fatal(msg string, fields ...Fields) {
	entry := l.logger.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = l.logger.WithFields(logrus.Fields(fields[0]))
	}
	entry.Fatal(msg)
}

// DebugContext 带上下文的调试日志
func (l *logrusLogger) DebugContext(ctx context.Context, msg string, fields ...Fields) {
	entry := l.logger.WithContext(ctx)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Debug(msg)
}

// InfoContext 带上下文的信息日志
func (l *logrusLogger) InfoContext(ctx context.Context, msg string, fields ...Fields) {
	entry := l.logger.WithContext(ctx)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Info(msg)
}

// WarnContext 带上下文的警告日志
func (l *logrusLogger) WarnContext(ctx context.Context, msg string, fields ...Fields) {
	entry := l.logger.WithContext(ctx)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Warn(msg)
}

// ErrorContext 带上下文的错误日志
func (l *logrusLogger) ErrorContext(ctx context.Context, msg string, fields ...Fields) {
	entry := l.logger.WithContext(ctx)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Error(msg)
}

// ErrorWithErr 带错误的错误日志
func (l *logrusLogger) ErrorWithErr(err error, msg string, fields ...Fields) {
	entry := l.logger.WithError(err)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Error(msg)
}

// WarnWithErr 带错误的警告日志
func (l *logrusLogger) WarnWithErr(err error, msg string, fields ...Fields) {
	entry := l.logger.WithError(err)
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}
	entry.Warn(msg)
}

// WithFields 创建带字段的子日志器
func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogger{
		logger: l.logger.WithFields(logrus.Fields(fields)).Logger,
	}
}

// WithField 创建带单个字段的子日志器
func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{
		logger: l.logger.WithField(key, value).Logger,
	}
}

// WithError 创建带错误的子日志器
func (l *logrusLogger) WithError(err error) Logger {
	return &logrusLogger{
		logger: l.logger.WithError(err).Logger,
	}
}

// WithContext 创建带上下文的子日志器
func (l *logrusLogger) WithContext(ctx context.Context) Logger {
	return &logrusLogger{
		logger: l.logger.WithContext(ctx).Logger,
	}
}

// SetLevel 设置日志级别
func (l *logrusLogger) SetLevel(level LogLevel) {
	logrusLevel, err := logrus.ParseLevel(string(level))
	if err != nil {
		logrusLevel = logrus.InfoLevel
	}
	l.logger.SetLevel(logrusLevel)
}
