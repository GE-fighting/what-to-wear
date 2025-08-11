package errors

import (
	"fmt"
	"net/http"
)

// APIError 自定义API错误类型
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// NewAPIError 创建新的API错误
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

func NewInternalError(message string, details ...string) *APIError {
	err := &APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// 常用错误封装 (使用标准 HTTP 状态码作为 code)

func ErrInvalidRequest(message string, details ...string) *APIError {
	return NewAPIError(http.StatusBadRequest, message, details...)
}

func ErrUnauthorized(message string, details ...string) *APIError {
	return NewAPIError(http.StatusUnauthorized, message, details...)
}

func ErrForbidden(message string, details ...string) *APIError {
	return NewAPIError(http.StatusForbidden, message, details...)
}

func ErrNotFound(message string, details ...string) *APIError {
	return NewAPIError(http.StatusNotFound, message, details...)
}

func ErrConflict(message string, details ...string) *APIError {
	return NewAPIError(http.StatusConflict, message, details...)
}

func ErrInternalServer(message string, details ...string) *APIError {
	return NewAPIError(http.StatusInternalServerError, message, details...)
}
