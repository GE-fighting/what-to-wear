package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	JWT      JWTConfig      `json:"jwt"`
	OSS      OSSConfig      `json:"oss"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type JWTConfig struct {
	Secret     string `json:"secret"`
	ExpireTime int    `json:"expire_time"`
}

type OSSConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	BucketName      string `json:"bucket_name"`
	Region          string `json:"region"`
	Expires         int64  `json:"expires"`
}

func LoadConfig() (*Config, error) {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		// .env 文件不存在时不报错，使用系统环境变量
	}

	config := &Config{
		Server: ServerConfig{
			Port: getEnvWithDefault("SERVER_PORT", "8080"),
			Mode: getEnvWithDefault("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnvWithDefault("DB_HOST", "localhost"),
			User:     getEnvWithDefault("DB_USER", "admin"),
			Password: getEnvWithDefault("DB_PASSWORD", "admin"),
			DBName:   getEnvWithDefault("DB_NAME", "what_to_wear"),
			Port:     getEnvWithDefault("DB_PORT", "5432"),
			SSLMode:  getEnvWithDefault("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:     getEnvWithDefault("JWT_SECRET", "your-secret-key"),
			ExpireTime: getEnvIntWithDefault("JWT_EXPIRE_TIME", 86400),
		},
		OSS: OSSConfig{
			Endpoint:        getEnvWithDefault("OSS_ENDPOINT", "oss-cn-hangzhou.aliyuncs.com"),
			AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
			AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
			BucketName:      getEnvWithDefault("OSS_BUCKET_NAME", "your-bucket-name"),
			Region:          getEnvWithDefault("OSS_REGION", "cn-hangzhou"),
			Expires:         getEnvInt64WithDefault("OSS_EXPIRES", 3600),
		},
	}

	return config, nil
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.Host, d.User, d.Password, d.DBName, d.Port, d.SSLMode)
}

func getEnvIntWithDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue := parseInt(value); intValue != 0 {
			return intValue
		}
	}
	return defaultValue
}

func getEnvInt64WithDefault(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue := parseInt64(value); intValue != 0 {
			return intValue
		}
	}
	return defaultValue
}

func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}

func parseInt64(s string) int64 {
	var result int64
	fmt.Sscanf(s, "%d", &result)
	return result
}
