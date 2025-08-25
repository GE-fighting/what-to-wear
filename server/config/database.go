package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// DatabaseConfig 数据库配置结构
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// LoadDatabaseConfig 加载数据库配置
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		// .env 文件不存在时不报错，使用系统环境变量
	}

	config := &DatabaseConfig{
		Host:     getEnvWithDefault("DB_HOST", "localhost"),
		User:     getEnvWithDefault("DB_USER", "admin"),
		Password: getEnvWithDefault("DB_PASSWORD", "admin"),
		DBName:   getEnvWithDefault("DB_NAME", "what_to_wear"),
		Port:     getEnvWithDefault("DB_PORT", "5432"),
		SSLMode:  getEnvWithDefault("DB_SSLMODE", "disable"),
	}

	return config, nil
}

// ConnectDatabase 连接数据库
func ConnectDatabase(config *DatabaseConfig) (*gorm.DB, error) {
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	// 配置GORM日志
	gormConfig := &gorm.Config{
		Logger: getGormLogger(),
	}

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

// getGormLogger 获取GORM日志配置
func getGormLogger() gormLogger.Interface {
	if os.Getenv("ENABLE_SQL_LOGGING") == "true" {
		return gormLogger.Default.LogMode(getGormLogLevel())
	}
	return gormLogger.Default.LogMode(gormLogger.Silent)
}

// getEnvWithDefault 获取环境变量，如果不存在则使用默认值
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getGormLogLevel 根据环境变量获取 GORM 日志级别
func getGormLogLevel() gormLogger.LogLevel {
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		return gormLogger.Info
	case "info":
		return gormLogger.Warn
	default:
		// 根据运行模式设置默认级别
		if os.Getenv("GIN_MODE") == "debug" {
			return gormLogger.Info
		}
		return gormLogger.Warn
	}
}
