package middleware

import (
	"bytes"
	"context"
	"time"
	"what-to-wear/server/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 计算处理耗时
		latency := time.Since(start)

		// 从 Gin context 获取 request_id
		requestID := c.GetString("request_id")

		// 使用结构化日志记录请求信息
		fields := logger.Fields{
			"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
			"status_code": c.Writer.Status(),
			"latency":     latency.String(),
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"user_agent":  c.Request.UserAgent(),
			"request_id":  requestID,
		}

		// 根据状态码选择日志级别
		switch {
		case c.Writer.Status() >= 500:
			logger.GetLogger().Error("HTTP Request", fields)
		case c.Writer.Status() >= 400:
			logger.GetLogger().Warn("HTTP Request", fields)
		default:
			logger.GetLogger().Info("HTTP Request", fields)
		}
	}
}

// RequestIDMiddleware 请求ID中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)

		// 将请求ID添加到标准 context 中
		ctx := context.WithValue(c.Request.Context(), "requestID", requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// responseBodyWriter 用于捕获响应体的写入器
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseBodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// ErrorLoggingMiddleware 错误日志中间件
func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 记录错误
		if len(c.Errors) > 0 {
			requestID := c.GetString("request_id")
			log := logger.GetLogger()
			for _, err := range c.Errors {
				log.Error("Request error", logger.Fields{
					"request_id": requestID,
					"error":      err.Error(),
					"type":       err.Type,
					"meta":       err.Meta,
					"method":     c.Request.Method,
					"path":       c.Request.URL.Path,
				})
			}
		}
	}
}
