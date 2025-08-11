package dto

// Result 响应结果结构体
type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// PaginationInfo 分页信息结构体
type PaginationInfo struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

// PaginatedResult 分页响应结构体
type PaginatedResult struct {
	Result
	Pagination *PaginationInfo `json:"pagination,omitempty"`
}

// CreateSuccess 创建成功响应
func CreateSuccess(data interface{}, message ...string) *Result {
	result := &Result{
		Success: true,
		Code:    200, // http.StatusOK
		Data:    data,
	}

	if len(message) > 0 && message[0] != "" {
		result.Message = message[0]
	}

	return result
}

// CreateError 创建错误响应
func CreateError(code int, message string) *Result {
	return &Result{
		Success: false,
		Code:    code,
		Message: message,
	}
}

// CreatePaginatedResult 创建分页响应
func CreatePaginatedResult(data interface{}, total int64, page, pageSize int, message ...string) *PaginatedResult {
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)
	if totalPages < 0 {
		totalPages = 0
	}

	result := &PaginatedResult{
		Result: Result{
			Success: true,
			Code:    200, // http.StatusOK
			Data:    data,
		},
		Pagination: &PaginationInfo{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}

	if len(message) > 0 && message[0] != "" {
		result.Message = message[0]
	}

	return result
}
