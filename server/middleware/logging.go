package middleware

import (
	"bytes"
	"io"
	"time"
	"what-to-wear/server/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 使用结构化日志记录请求信息
		fields := logger.Fields{
			"timestamp":    param.TimeStamp.Format("2006-01-02 15:04:05"),
			"status_code":  param.StatusCode,
			"latency":      param.Latency.String(),
			"client_ip":    param.ClientIP,
			"method":       param.Method,
			"path":         param.Path,
			"user_agent":   param.Request.UserAgent(),
			"request_id":   param.Request.Header.Get("X-Request-ID"),
		}

		// 根据状态码选择日志级别
		switch {
		case param.StatusCode >= 500:
			logger.Error("HTTP Request", fields)
		case param.StatusCode >= 400:
			logger.Warn("HTTP Request", fields)
		default:
			logger.Info("HTTP Request", fields)
		}

		return ""
	})
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
		
		// 将请求ID添加到日志上下文
		ctx := c.Request.Context()
		c.Request = c.Request.WithContext(ctx)
		
		c.Next()
	}
}

// DetailedLoggingMiddleware 详细日志中间件（包含请求体和响应体）
func DetailedLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetString("request_id")
		
		// 记录请求体（仅对POST、PUT、PATCH请求）
		var requestBody []byte
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			if c.Request.Body != nil {
				requestBody, _ = io.ReadAll(c.Request.Body)
				c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
			}
		}

		// 创建响应写入器来捕获响应体
		responseWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:          &bytes.Buffer{},
		}
		c.Writer = responseWriter

		// 记录请求开始
		logger.Info("Request started", logger.Fields{
			"request_id":   requestID,
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"query":        c.Request.URL.RawQuery,
			"client_ip":    c.ClientIP(),
			"user_agent":   c.Request.UserAgent(),
			"content_type": c.Request.Header.Get("Content-Type"),
			"request_body": string(requestBody),
		})

		c.Next()

		// 记录请求结束
		duration := time.Since(start)
		fields := logger.Fields{
			"request_id":    requestID,
			"status_code":   c.Writer.Status(),
			"duration":      duration.String(),
			"response_size": responseWriter.body.Len(),
		}

		// 只在开发环境记录响应体
		if gin.Mode() == gin.DebugMode {
			fields["response_body"] = responseWriter.body.String()
		}

		// 根据状态码和错误选择日志级别
		switch {
		case c.Writer.Status() >= 500:
			logger.Error("Request completed with server error", fields)
		case c.Writer.Status() >= 400:
			logger.Warn("Request completed with client error", fields)
		case duration > time.Second*5:
			logger.Warn("Slow request completed", fields)
		default:
			logger.Info("Request completed", fields)
		}
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
			for _, err := range c.Errors {
				logger.Error("Request error", logger.Fields{
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
