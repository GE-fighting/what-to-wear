# 统一响应系统使用指南

## 概述

本项目实现了统一的API响应格式，确保前后端交互的一致性和可预测性。

## 响应结构

### 基础响应结构 (`dto.BaseResponse`)

```go
type BaseResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
}
```

### 错误信息结构 (`dto.ErrorInfo`)

```go
type ErrorInfo struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}
```

### 分页响应结构 (`dto.PaginatedResponse`)

```go
type PaginatedResponse struct {
    BaseResponse
    Pagination *PaginationInfo `json:"pagination,omitempty"`
}
```

## 使用方法

### 1. 成功响应

```go
// 基本成功响应
common.Success(c, data, "操作成功")

// 创建成功响应 (201)
common.Created(c, data, "创建成功")

// 分页响应
common.Paginated(c, data, total, page, pageSize, "获取列表成功")
```

### 2. 错误响应

```go
// 400 - 请求参数错误
common.BadRequest(c, "请求参数无效", "详细错误信息")

// 401 - 未授权
common.Unauthorized(c, "未授权访问", "请先登录")

// 403 - 权限不足
common.Forbidden(c, "权限不足", "您没有访问此资源的权限")

// 404 - 资源不存在
common.NotFound(c, "资源不存在", "用户不存在")

// 409 - 资源冲突
common.Conflict(c, "资源冲突", "用户名已存在")

// 500 - 服务器错误
common.InternalServerError(c, "服务器内部错误", "数据库连接失败")

// 通用错误处理
common.Error(c, err)
```

## 错误代码规范

### 通用错误代码
- `INVALID_REQUEST` - 请求参数无效
- `UNAUTHORIZED` - 未授权访问
- `FORBIDDEN` - 权限不足
- `NOT_FOUND` - 资源不存在
- `CONFLICT` - 资源冲突
- `INTERNAL_SERVER_ERROR` - 服务器内部错误
- `SERVICE_UNAVAILABLE` - 服务不可用

### 业务错误代码
- `USER_NOT_FOUND` - 用户不存在
- `USER_EXISTS` - 用户已存在
- `INVALID_CREDENTIALS` - 凭据无效
- `INVALID_PASSWORD` - 密码错误
- `TOKEN_EXPIRED` - Token已过期
- `TOKEN_INVALID` - Token无效

## 响应示例

### 成功响应

```json
{
  "success": true,
  "message": "获取用户信息成功",
  "data": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

### 错误响应

```json
{
  "success": false,
  "error": {
    "code": "INVALID_REQUEST",
    "message": "请求参数无效",
    "details": "用户名不能为空"
  }
}
```

### 分页响应

```json
{
  "success": true,
  "message": "获取用户列表成功",
  "data": [
    {"id": 1, "username": "user1"},
    {"id": 2, "username": "user2"}
  ],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

## 中间件配置

确保在路由中正确配置中间件：

```go
// 在 main.go 或路由配置中
r.Use(middleware.RequestIDMiddleware())  // 生成请求ID（用于日志追踪）
r.Use(middleware.LoggingMiddleware())    // 日志记录
```

## 最佳实践

1. **统一使用响应函数**：不要直接使用 `c.JSON()`，始终使用 `common` 包中的响应函数
2. **提供有意义的消息**：错误消息应该对用户友好，详细信息用于调试
3. **使用合适的HTTP状态码**：通过响应函数自动设置正确的状态码
4. **简洁的响应结构**：响应结构保持简洁，专注于业务数据
5. **国际化支持**：消息文本应该支持多语言（后续可扩展）
6. **请求追踪**：虽然响应中不包含请求ID，但可通过日志和响应头进行追踪

## 前端处理建议

前端可以根据 `success` 字段判断请求是否成功：

```javascript
// 成功处理
if (response.success) {
  console.log(response.message);
  // 处理 response.data
} else {
  // 错误处理
  console.error(response.error.message);
  // 根据 response.error.code 进行特定处理
}
```
