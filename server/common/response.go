package common

import (
	"net/http"
	"what-to-wear/server/dto"
	"what-to-wear/server/errors"
	"what-to-wear/server/logger"

	"github.com/gin-gonic/gin"
)

// Success 返回成功响应
func Success(c *gin.Context, data interface{}, message ...string) {
	response := dto.CreateSuccess(data, message...)
	c.JSON(http.StatusOK, response)
}

// Created 创建成功响应
func Created(c *gin.Context, data interface{}, message ...string) {
	response := dto.CreateSuccess(data, message...)
	c.JSON(http.StatusCreated, response)
}

// Paginated 返回分页响应
func Paginated(c *gin.Context, data interface{}, total int64, page, pageSize int, message ...string) {
	response := dto.CreatePaginatedResult(data, total, page, pageSize, message...)
	c.JSON(http.StatusOK, response)
}

// Error 统一错误处理入口
func Error(c *gin.Context, err error) {
	log := logger.FromContext(c.Request.Context())

	// 如果是 APIError，记录警告日志并直接使用其信息构建响应
	if apiErr, ok := err.(*errors.APIError); ok {
		log.Error("API error occurred",
			logger.Fields{
				"code":    apiErr.Code,
				"message": apiErr.Message,
				"details": apiErr.Details,
				"path":    c.Request.URL.Path,
				"method":  c.Request.Method,
			})
		response := dto.CreateError(apiErr.Code, apiErr.Message)
		c.JSON(apiErr.Code, response)
		return
	}

	// 对于其他类型的错误，记录详细错误日志并返回通用的500错误给客户端
	log.Error("Internal server error",
		logger.Fields{
			"error":  err.Error(),
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})
	response := dto.CreateError(http.StatusInternalServerError, "服务器内部错误")
	c.JSON(http.StatusInternalServerError, response)
}
