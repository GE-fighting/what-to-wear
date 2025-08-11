package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

// slogLogger 使用 slog 实现的日志器
type slogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger 创建基于 slog 的日志器
func NewSlogLogger(config *Config) Logger {
	var handler slog.Handler

	// 输出目标
	var output io.Writer
	switch config.Output {
	case "file":
		if err := os.MkdirAll(filepath.Dir(config.FilePath), 0755); err != nil {
			slog.Warn("Failed to create log directory", "error", err)
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
		if err := os.MkdirAll(filepath.Dir(config.FilePath), 0755); err != nil {
			slog.Warn("Failed to create log directory", "error", err)
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

	// 日志级别
	level := slog.LevelInfo
	switch config.Level {
	case DebugLevel:
		level = slog.LevelDebug
	case InfoLevel:
		level = slog.LevelInfo
	case WarnLevel:
		level = slog.LevelWarn
	case ErrorLevel:
		level = slog.LevelError
	case FatalLevel:
		level = slog.LevelError
	}

	// 格式
	if config.Format == "json" {
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	}

	return &slogLogger{
		logger: slog.New(handler),
	}
}

// fieldsToAttrs 将 Fields 转换为 slog.Attr
func fieldsToAttrs(fields Fields) []any {
	attrs := make([]any, 0, len(fields))
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}

// Debug 调试日志
func (l *slogLogger) Debug(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Debug(msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.Debug(msg)
	}
}

// Info 信息日志
func (l *slogLogger) Info(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Info(msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.Info(msg)
	}
}

// Warn 警告日志
func (l *slogLogger) Warn(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Warn(msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.Warn(msg)
	}
}

// Error 错误日志
func (l *slogLogger) Error(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Error(msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.Error(msg)
	}
}

// Fatal 致命错误日志（slog 没有 fatal，模拟）
func (l *slogLogger) Fatal(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Error(msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.Error(msg)
	}
	os.Exit(1)
}

// DebugContext 带上下文的调试日志
func (l *slogLogger) DebugContext(ctx context.Context, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.DebugContext(ctx, msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.DebugContext(ctx, msg)
	}
}

// InfoContext 带上下文的信息日志
func (l *slogLogger) InfoContext(ctx context.Context, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.InfoContext(ctx, msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.InfoContext(ctx, msg)
	}
}

// WarnContext 带上下文的警告日志
func (l *slogLogger) WarnContext(ctx context.Context, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.WarnContext(ctx, msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.WarnContext(ctx, msg)
	}
}

// ErrorContext 带上下文的错误日志
func (l *slogLogger) ErrorContext(ctx context.Context, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.ErrorContext(ctx, msg, fieldsToAttrs(fields[0])...)
	} else {
		l.logger.ErrorContext(ctx, msg)
	}
}

// ErrorWithErr 带错误的错误日志
func (l *slogLogger) ErrorWithErr(err error, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Error(msg, append(fieldsToAttrs(fields[0]), slog.Any("error", err))...)
	} else {
		l.logger.Error(msg, slog.Any("error", err))
	}
}

// WarnWithErr 带错误的警告日志
func (l *slogLogger) WarnWithErr(err error, msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.logger.Warn(msg, append(fieldsToAttrs(fields[0]), slog.Any("error", err))...)
	} else {
		l.logger.Warn(msg, slog.Any("error", err))
	}
}

// WithFields 创建带字段的子日志器
func (l *slogLogger) WithFields(fields Fields) Logger {
	attrs := fieldsToAttrs(fields)
	return &slogLogger{
		logger: l.logger.With(attrs...),
	}
}

// WithField 创建带单个字段的子日志器
func (l *slogLogger) WithField(key string, value interface{}) Logger {
	return &slogLogger{
		logger: l.logger.With(slog.Any(key, value)),
	}
}

// WithError 创建带错误的子日志器
func (l *slogLogger) WithError(err error) Logger {
	return &slogLogger{
		logger: l.logger.With(slog.Any("error", err)),
	}
}

// WithContext 创建带上下文的子日志器（注入 requestID）
func (l *slogLogger) WithContext(ctx context.Context) Logger {
	if requestID, ok := ctx.Value("requestID").(string); ok {
		return &slogLogger{
			logger: l.logger.With(slog.String("requestID", requestID)),
		}
	}
	return l
}

// SetLevel 设置日志级别（slog 不支持动态设置，忽略）
func (l *slogLogger) SetLevel(level LogLevel) {
	// slog 不支持运行时动态修改 level
}
