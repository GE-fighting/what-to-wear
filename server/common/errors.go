package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError 自定义错误类型
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

// 预定义错误
var (
	ErrInvalidRequest   = &APIError{Code: http.StatusBadRequest, Message: "Invalid request"}
	ErrUnauthorized     = &APIError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden        = &APIError{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound         = &APIError{Code: http.StatusNotFound, Message: "Resource not found"}
	ErrConflict         = &APIError{Code: http.StatusConflict, Message: "Resource conflict"}
	ErrInternalServer   = &APIError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	ErrServiceUnavailable = &APIError{Code: http.StatusServiceUnavailable, Message: "Service unavailable"}
)

// NewAPIError 创建自定义错误
func NewAPIError(code int, message string, details ...string) *APIError {
	err := &APIError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// HandleError 统一错误处理
func HandleError(c *gin.Context, err error) {
	if apiErr, ok := err.(*APIError); ok {
		c.JSON(apiErr.Code, gin.H{
			"error":   apiErr.Message,
			"details": apiErr.Details,
		})
		return
	}

	// 处理其他类型的错误
	switch err.Error() {
	case "user not found":
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	case "username already exists":
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
	case "email already exists":
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
	case "invalid username or password":
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	case "invalid old password":
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid old password"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}, message ...string) {
	response := gin.H{
		"success": true,
		"data":    data,
	}
	
	if len(message) > 0 {
		response["message"] = message[0]
	}
	
	c.JSON(http.StatusOK, response)
}

// CreatedResponse 创建成功响应
func CreatedResponse(c *gin.Context, data interface{}, message ...string) {
	response := gin.H{
		"success": true,
		"data":    data,
	}
	
	if len(message) > 0 {
		response["message"] = message[0]
	}
	
	c.JSON(http.StatusCreated, response)
}

// PaginatedResponse 分页响应
func PaginatedResponse(c *gin.Context, data interface{}, total int64, page, pageSize int) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"pagination": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}
