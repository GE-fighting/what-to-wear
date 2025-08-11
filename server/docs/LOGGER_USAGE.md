# 日志系统使用指南

## 概述

本项目的日志系统支持自动包含 `request_id` 的便捷函数，实现链路追踪。通过中间件预创建日志器，避免重复实例化，提高性能。

## 推荐使用方式（高性能）

### 方式一：获取预创建的日志器（推荐）

在请求开始时获取一次预创建的日志器，整个请求过程中复用：

```go
func (c *Controller) SomeHandler(ctx *gin.Context) {
    // 在请求开始时获取预创建的日志器（只创建一次）
    log := logger.FromGinContext(ctx)

    log.Info("Processing request", logger.Fields{
        "user_id": userID,
        "action":  "get_profile",
    })

    log.Error("Request failed", logger.Fields{
        "error": err.Error(),
    })
}
```

### 方式二：直接使用便捷函数

如果不关心性能，可以直接使用便捷函数：

```go
// 每次调用都会重新获取日志器
logger.InfoFromContext(ctx, "Processing request", logger.Fields{...})
logger.ErrorFromContext(ctx, "Request failed", logger.Fields{...})
```

### 可用的便捷函数

```go
// 调试日志
logger.DebugFromContext(c, "Debug message", logger.Fields{...})

// 信息日志
logger.InfoFromContext(c, "Info message", logger.Fields{...})

// 警告日志
logger.WarnFromContext(c, "Warning message", logger.Fields{...})

// 错误日志
logger.ErrorFromContext(c, "Error message", logger.Fields{...})
```

### 创建子日志器（高级用法）

如果需要在函数间传递日志器：

```go
// 从 Gin Context 创建子日志器
subLogger := logger.FromGinContext(c)
subLogger.Info("Message with request_id", logger.Fields{...})
```

## 使用示例

### 控制器中的完整示例（推荐方式）

```go
package controllers

import (
    "what-to-wear/server/common"
    "what-to-wear/server/logger"
    "github.com/gin-gonic/gin"
)

func (uc *UserController) GetProfile(c *gin.Context) {
    // 在请求开始时获取预创建的日志器（性能最佳）
    log := logger.FromGinContext(c)

    userID := c.GetUint("user_id")

    log.Info("Getting user profile", logger.Fields{
        "user_id": userID,
    })

    user, err := uc.userService.GetProfile(userID)
    if err != nil {
        log.Error("Failed to get user profile", logger.Fields{
            "user_id": userID,
            "error":   err.Error(),
        })
        common.Error(c, err)
        return
    }

    log.Info("User profile retrieved successfully", logger.Fields{
        "user_id": userID,
    })

    common.Success(c, user, "获取用户资料成功")
}
```

### 性能对比

**❌ 低性能方式（每次都创建新日志器）：**
```go
func Handler(c *gin.Context) {
    logger.InfoFromContext(c, "Step 1", logger.Fields{...})  // 创建日志器1
    logger.InfoFromContext(c, "Step 2", logger.Fields{...})  // 创建日志器2
    logger.InfoFromContext(c, "Step 3", logger.Fields{...})  // 创建日志器3
    // 一个请求创建了3个日志器实例
}
```

**✅ 高性能方式（复用一个日志器）：**
```go
func Handler(c *gin.Context) {
    log := logger.FromGinContext(c)  // 只创建一次

    log.Info("Step 1", logger.Fields{...})  // 复用
    log.Info("Step 2", logger.Fields{...})  // 复用
    log.Info("Step 3", logger.Fields{...})  // 复用
    // 一个请求只创建了1个日志器实例
}
```

### 服务层中的使用

如果需要在服务层记录日志，可以传递子日志器：

```go
// 在控制器中
subLogger := logger.FromGinContext(c)
user, err := uc.userService.GetProfile(userID, subLogger)

// 在服务层中
func (s *userService) GetProfile(userID uint, log logger.Logger) (*models.User, error) {
    log.Info("Fetching user from database", logger.Fields{
        "user_id": userID,
    })

    // ... 业务逻辑

    return user, nil
}
```

## 日志输出示例

使用这些便捷函数后，日志会自动包含 `request_id`：

```json
{
  "level": "info",
  "msg": "Processing user registration",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "john_doe",
  "email": "john@example.com",
  "time": "2024-01-01T12:00:00Z"
}
```

## 对比：使用前后

### 使用前（手动传递 request_id）

```go
func (ac *AuthController) Register(c *gin.Context) {
    requestID := c.GetString("request_id")  // 手动获取
    
    logger.Info("Processing user registration", logger.Fields{
        "request_id": requestID,  // 手动添加
        "username":   req.Username,
        "email":      req.Email,
    })
}
```

### 使用后（自动包含 request_id）

```go
func (ac *AuthController) Register(c *gin.Context) {
    // 自动包含 request_id，代码更简洁
    logger.InfoFromContext(c, "Processing user registration", logger.Fields{
        "username": req.Username,
        "email":    req.Email,
    })
}
```

## 最佳实践

1. **在控制器中优先使用 `*FromContext` 函数**：这些函数最简洁，自动处理 `request_id`

2. **保持日志字段一致性**：使用统一的字段名，如 `user_id`、`error`、`action` 等

3. **记录关键操作**：在重要的业务操作前后记录日志，便于问题追踪

4. **错误日志包含详细信息**：错误日志应该包含足够的上下文信息

5. **避免敏感信息**：不要在日志中记录密码、token 等敏感信息

## 注意事项

- 便捷函数依赖于 `RequestIDMiddleware` 中间件，确保在路由中正确配置
- `request_id` 字段会自动添加到日志中，无需在 `Fields` 中手动指定
- 如果 `request_id` 不存在，函数会正常工作，只是不会包含该字段
- 推荐在控制器中使用 `*FromContext` 函数，代码最简洁
