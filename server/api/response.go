package api

// Response 统一响应结构 - 简化版
type Response struct {
	Code    int         `json:"code"`           // 业务状态码
	Message string      `json:"message"`        // 响应消息
	Data    interface{} `json:"data,omitempty"` // 响应数据
}

// Success 成功响应
func Success(data interface{}, message ...string) Response {
	msg := "success"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	return Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

// Error 错误响应
func Error(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}

// BadRequest 400错误
func BadRequest(message string) Response {
	return Error(400, message)
}

// Unauthorized 401错误
func Unauthorized(message ...string) Response {
	msg := "unauthorized"
	if len(message) > 0 {
		msg = message[0]
	}
	return Error(401, msg)
}

// Forbidden 403错误
func Forbidden(message ...string) Response {
	msg := "forbidden"
	if len(message) > 0 {
		msg = message[0]
	}
	return Error(403, msg)
}

// NotFound 404错误
func NotFound(message ...string) Response {
	msg := "not found"
	if len(message) > 0 {
		msg = message[0]
	}
	return Error(404, msg)
}

// InternalError 500错误
func InternalError(message ...string) Response {
	msg := "internal server error"
	if len(message) > 0 {
		msg = message[0]
	}
	return Error(500, msg)
}

// PageData 分页数据结构
type PageData struct {
	Items interface{} `json:"items"`           // 数据列表
	Total int64       `json:"total"`           // 总数量
	Page  int         `json:"page"`            // 当前页
	Size  int         `json:"size"`            // 每页大小
	Pages int         `json:"pages,omitempty"` // 总页数
}

// SuccessWithPage 分页成功响应
func SuccessWithPage(items interface{}, total int64, page, size int, message ...string) Response {
	pages := int(total) / size
	if int(total)%size > 0 {
		pages++
	}

	pageData := PageData{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
		Pages: pages,
	}

	return Success(pageData, message...)
}
