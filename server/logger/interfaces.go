package logger

import "context"

// LogLevel 日志级别
type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

// Fields 日志字段类型
type Fields map[string]interface{}

// Logger 日志接口
type Logger interface {
	// 基础日志方法
	Debug(msg string, fields ...Fields)
	Info(msg string, fields ...Fields)
	Warn(msg string, fields ...Fields)
	Error(msg string, fields ...Fields)
	Fatal(msg string, fields ...Fields)

	// 带上下文的日志方法
	DebugContext(ctx context.Context, msg string, fields ...Fields)
	InfoContext(ctx context.Context, msg string, fields ...Fields)
	WarnContext(ctx context.Context, msg string, fields ...Fields)
	ErrorContext(ctx context.Context, msg string, fields ...Fields)

	// 带错误的日志方法
	ErrorWithErr(err error, msg string, fields ...Fields)
	WarnWithErr(err error, msg string, fields ...Fields)

	// 创建子日志器
	WithFields(fields Fields) Logger
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	// 设置日志级别
	SetLevel(level LogLevel)
}

// Config 日志配置
type Config struct {
	Level      LogLevel `json:"level" yaml:"level"`
	Format     string   `json:"format" yaml:"format"`           // json, text
	Output     string   `json:"output" yaml:"output"`           // stdout, file
	FilePath   string   `json:"file_path" yaml:"file_path"`     // 日志文件路径
	MaxSize    int      `json:"max_size" yaml:"max_size"`       // 文件最大大小(MB)
	MaxBackups int      `json:"max_backups" yaml:"max_backups"` // 保留的备份文件数
	MaxAge     int      `json:"max_age" yaml:"max_age"`         // 保留天数
	Compress   bool     `json:"compress" yaml:"compress"`       // 是否压缩
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Level:      InfoLevel,
		Format:     "json",
		Output:     "stdout",
		FilePath:   "logs/app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
}
