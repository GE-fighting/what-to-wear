# 配置文档 (Configuration Documentation)

## 概述

本项目使用环境变量进行配置管理，支持多环境部署。配置文件位于项目根目录：

- `.env.example` - 配置模板和说明
- `.env.development` - 开发环境配置
- `.env.production` - 生产环境配置

## 配置分类

### 1. 数据库配置 (Database Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `DB_HOST` | string | localhost | 数据库主机地址 |
| `DB_PORT` | int | 5432 | 数据库端口 |
| `DB_USER` | string | admin | 数据库用户名 |
| `DB_PASSWORD` | string | admin | 数据库密码 |
| `DB_NAME` | string | what_to_wear | 数据库名称 |
| `DB_SSLMODE` | string | disable | SSL模式 (disable/require/verify-full) |

### 2. JWT 认证配置 (JWT Authentication)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `JWT_SECRET` | string | - | JWT签名密钥 (生产环境必须修改) |

### 3. 日志配置 (Logging Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `LOG_LEVEL` | string | info | 日志级别 (debug/info/warn/error/fatal) |
| `LOG_FORMAT` | string | json | 日志格式 (json/text) |
| `LOG_OUTPUT` | string | both | 输出方式 (stdout/file/both) |
| `LOG_TO_FILE` | bool | true | 是否输出到文件 (兼容旧配置) |
| `LOG_FILE_PATH` | string | logs/app.log | 日志文件路径 |
| `LOG_MAX_SIZE` | int | 100 | 日志文件最大大小 (MB) |
| `LOG_MAX_BACKUPS` | int | 3 | 保留的备份文件数量 |
| `LOG_MAX_AGE` | int | 7 | 日志文件保留天数 |
| `LOG_COMPRESS` | bool | true | 是否压缩旧日志文件 |

### 4. 服务器配置 (Server Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `GIN_MODE` | string | release | 运行模式 (debug/release/test) |
| `PORT` | int | 8080 | 服务器端口 |
| `HOST` | string | 0.0.0.0 | 服务器主机地址 |

### 5. CORS 跨域配置 (CORS Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `CORS_ORIGINS` | string | localhost:1420 | 允许的源地址 (逗号分隔) |
| `CORS_METHODS` | string | GET,POST,PUT,DELETE,OPTIONS | 允许的HTTP方法 |
| `CORS_HEADERS` | string | Origin,Content-Type,Authorization | 允许的请求头 |
| `CORS_EXPOSE_HEADERS` | string | Content-Length,X-Request-ID | 暴露的响应头 |
| `CORS_CREDENTIALS` | bool | true | 是否允许携带凭证 |

### 6. 开发环境配置 (Development Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `ENABLE_DETAILED_LOGGING` | bool | false | 是否启用详细日志 (包含请求体/响应体) |
| `ENABLE_SQL_LOGGING` | bool | false | 是否启用SQL查询日志 |
| `API_TIMEOUT` | int | 30 | API请求超时时间 (秒) |

### 7. 生产环境配置 (Production Configuration)

| 变量名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `ENABLE_METRICS` | bool | false | 是否启用性能监控 |
| `ENABLE_HEALTH_CHECK` | bool | true | 是否启用健康检查端点 |
| `MAX_CONNECTIONS` | int | 1000 | 最大并发连接数 |
| `MAX_REQUEST_SIZE` | int | 10 | 请求体最大大小 (MB) |

## 环境配置建议

### 开发环境 (Development)
```bash
LOG_LEVEL=debug
LOG_FORMAT=text
LOG_OUTPUT=stdout
GIN_MODE=debug
ENABLE_DETAILED_LOGGING=true
ENABLE_SQL_LOGGING=true
```

### 测试环境 (Testing)
```bash
LOG_LEVEL=info
LOG_FORMAT=json
LOG_OUTPUT=both
GIN_MODE=test
ENABLE_DETAILED_LOGGING=false
ENABLE_SQL_LOGGING=false
```

### 生产环境 (Production)
```bash
LOG_LEVEL=warn
LOG_FORMAT=json
LOG_OUTPUT=file
GIN_MODE=release
ENABLE_DETAILED_LOGGING=false
ENABLE_SQL_LOGGING=false
ENABLE_METRICS=true
```

## 使用方法

1. **复制配置模板**
   ```bash
   cp .env.example .env
   ```

2. **根据环境修改配置**
   - 开发环境：使用 `.env.development` 作为参考
   - 生产环境：使用 `.env.production` 作为参考

3. **加载环境变量**
   ```bash
   # 使用 godotenv 自动加载 .env 文件
   # 或者手动设置环境变量
   export LOG_LEVEL=debug
   ```

## 安全注意事项

1. **敏感信息保护**
   - 生产环境的 `JWT_SECRET` 必须使用强密码
   - 数据库密码不应使用默认值
   - `.env` 文件不应提交到版本控制

2. **权限控制**
   - 日志文件目录需要适当的写权限
   - 生产环境建议使用专用用户运行服务

3. **网络安全**
   - 生产环境的 CORS 配置应严格限制允许的域名
   - 数据库连接建议启用 SSL

## 故障排除

### 常见问题

1. **日志文件无法创建**
   - 检查 `LOG_FILE_PATH` 目录是否存在
   - 检查目录写权限

2. **数据库连接失败**
   - 验证数据库配置参数
   - 检查网络连接和防火墙设置

3. **CORS 错误**
   - 检查 `CORS_ORIGINS` 是否包含前端域名
   - 确认 `CORS_CREDENTIALS` 设置正确

### 调试技巧

1. **启用调试日志**
   ```bash
   LOG_LEVEL=debug
   GIN_MODE=debug
   ```

2. **查看详细请求日志**
   ```bash
   ENABLE_DETAILED_LOGGING=true
   ```

3. **监控SQL查询**
   ```bash
   ENABLE_SQL_LOGGING=true
   ```
