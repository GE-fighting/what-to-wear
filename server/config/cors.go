package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSConfig CORS配置结构
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}

// LoadCORSConfig 加载CORS配置
func LoadCORSConfig() *CORSConfig {
	return &CORSConfig{
		AllowOrigins:     getAllowOrigins(),
		AllowMethods:     getAllowMethods(),
		AllowHeaders:     getAllowHeaders(),
		ExposeHeaders:    getExposeHeaders(),
		AllowCredentials: getAllowCredentials(),
	}
}

// GetCORSMiddleware 获取CORS中间件
func GetCORSMiddleware() gin.HandlerFunc {
	config := LoadCORSConfig()
	return cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		ExposeHeaders:    config.ExposeHeaders,
		AllowCredentials: config.AllowCredentials,
	})
}

// getAllowOrigins 获取允许的源地址
func getAllowOrigins() []string {
	origins := os.Getenv("CORS_ORIGINS")
	if origins != "" {
		// 解析逗号分隔的源地址列表
		originList := strings.Split(origins, ",")
		// 去除每个元素的前后空格
		for i, origin := range originList {
			originList[i] = strings.TrimSpace(origin)
		}
		return originList
	}

	// 根据运行模式设置默认值
	if gin.Mode() == gin.DebugMode {
		return []string{
			"http://localhost:1420",
			"https://tauri.localhost",
			"http://localhost:3000",
			"http://localhost:8080",
		}
	}

	// 生产环境默认值 - 更严格的配置
	return []string{
		"https://tauri.localhost",
	}
}

// getAllowMethods 获取允许的HTTP方法
func getAllowMethods() []string {
	methods := os.Getenv("CORS_METHODS")
	if methods != "" {
		methodList := strings.Split(methods, ",")
		for i, method := range methodList {
			methodList[i] = strings.TrimSpace(method)
		}
		return methodList
	}

	// 默认允许的方法
	return []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
}

// getAllowHeaders 获取允许的请求头
func getAllowHeaders() []string {
	headers := os.Getenv("CORS_HEADERS")
	if headers != "" {
		headerList := strings.Split(headers, ",")
		for i, header := range headerList {
			headerList[i] = strings.TrimSpace(header)
		}
		return headerList
	}

	// 默认允许的请求头
	return []string{
		"Origin",
		"Content-Type",
		"Authorization",
		"X-Request-ID",
		"Accept",
		"X-Requested-With",
	}
}

// getExposeHeaders 获取暴露的响应头
func getExposeHeaders() []string {
	headers := os.Getenv("CORS_EXPOSE_HEADERS")
	if headers != "" {
		headerList := strings.Split(headers, ",")
		for i, header := range headerList {
			headerList[i] = strings.TrimSpace(header)
		}
		return headerList
	}

	// 默认暴露的响应头
	return []string{
		"Content-Length",
		"X-Request-ID",
	}
}

// getAllowCredentials 获取是否允许携带凭证
func getAllowCredentials() bool {
	credentials := os.Getenv("CORS_CREDENTIALS")
	if credentials != "" {
		if allow, err := strconv.ParseBool(credentials); err == nil {
			return allow
		}
	}

	// 默认允许携带凭证（支持JWT认证）
	return true
}
