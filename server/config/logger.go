package config

import (
	"os"
	"strconv"
	"what-to-wear/server/logger"

	"github.com/gin-gonic/gin"
)

// InitLogger 初始化日志系统
func InitLogger() {
	config := &logger.Config{
		Level:      getLogLevel(),
		Format:     getLogFormat(),
		Output:     getLogOutput(),
		FilePath:   getLogFilePath(),
		MaxSize:    getLogMaxSize(),
		MaxBackups: getLogMaxBackups(),
		MaxAge:     getLogMaxAge(),
		Compress:   getLogCompress(),
	}

	logger.Init(config)
}

// getLogLevel 获取日志级别
func getLogLevel() logger.LogLevel {
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		return logger.DebugLevel
	case "info":
		return logger.InfoLevel
	case "warn":
		return logger.WarnLevel
	case "error":
		return logger.ErrorLevel
	case "fatal":
		return logger.FatalLevel
	default:
		// 根据运行模式设置默认级别
		if gin.Mode() == gin.DebugMode {
			return logger.DebugLevel
		}
		return logger.InfoLevel
	}
}

// getLogOutput 获取日志输出方式
func getLogOutput() string {
	output := os.Getenv("LOG_OUTPUT")
	if output != "" {
		return output
	}

	// 兼容旧的配置方式
	if os.Getenv("LOG_TO_FILE") == "true" {
		return "both"
	}

	// 根据运行模式设置默认输出
	if gin.Mode() == gin.DebugMode {
		return "stdout"
	}
	return "both"
}

// getLogFormat 获取日志格式
func getLogFormat() string {
	format := os.Getenv("LOG_FORMAT")
	if format == "" {
		if gin.Mode() == gin.DebugMode {
			return "text"
		}
		return "json"
	}
	return format
}

// getLogFilePath 获取日志文件路径
func getLogFilePath() string {
	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {
		return "logs/app.log"
	}
	return path
}

// getLogMaxSize 获取日志文件最大大小
func getLogMaxSize() int {
	sizeStr := os.Getenv("LOG_MAX_SIZE")
	if sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil && size > 0 {
			return size
		}
	}
	// 默认100MB
	return 100
}

// getLogMaxBackups 获取日志备份文件数量
func getLogMaxBackups() int {
	backupsStr := os.Getenv("LOG_MAX_BACKUPS")
	if backupsStr != "" {
		if backups, err := strconv.Atoi(backupsStr); err == nil && backups >= 0 {
			return backups
		}
	}
	// 默认保留3个备份文件
	return 3
}

// getLogMaxAge 获取日志保留天数
func getLogMaxAge() int {
	ageStr := os.Getenv("LOG_MAX_AGE")
	if ageStr != "" {
		if age, err := strconv.Atoi(ageStr); err == nil && age > 0 {
			return age
		}
	}
	// 默认保留7天
	return 7
}

// getLogCompress 获取是否压缩日志
func getLogCompress() bool {
	compressStr := os.Getenv("LOG_COMPRESS")
	if compressStr != "" {
		if compress, err := strconv.ParseBool(compressStr); err == nil {
			return compress
		}
	}
	// 默认压缩
	return true
}
