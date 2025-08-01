# 今天穿什么 (What to Wear) 项目文档

## 📖 项目概述

"今天穿什么"是一个基于天气数据的智能穿衣建议桌面应用，帮助用户根据当天天气情况选择合适的服装搭配。

### 技术栈
- **前端**: React + TypeScript + Tauri (桌面应用)
- **后端**: Go + Gin + GORM + PostgreSQL
- **认证**: JWT Token
- **数据库**: PostgreSQL
- **开发工具**: Vite, pnpm

## 🏗️ 项目架构

### 混合架构设计
```
┌─────────────────────────────────────┐
│           Tauri 前端                │
│         (React + TS)                │
└─────────────┬───────────────────────┘
              │
        ┌─────┴─────┐
        │           │
        ▼           ▼
┌─────────────┐ ┌─────────────┐    GORM     ┌─────────────────┐
│ Rust 后端   │ │ Go 后端服务器│ ──────────→ │   PostgreSQL    │
│ (本地功能)  │ │ (网络 API)  │ ←────────── │     数据库      │
└─────────────┘ └─────────────┘             └─────────────────┘
```

### 架构分工
- **Go 后端**: 处理网络 API、数据库操作、用户认证、第三方服务集成
- **Rust 后端**: 处理本地文件操作、系统集成、性能计算、离线功能
- **前端**: 用户界面、状态管理、用户交互

## 📁 项目结构

```
what-to-wear/
├── server/                          # Go 后端服务
│   ├── config/                      # 配置模块
│   │   └── database.go              # 数据库连接配置
│   ├── controllers/                 # 控制器层
│   │   └── auth.go                  # 认证控制器
│   ├── middleware/                  # 中间件
│   │   └── auth.go                  # JWT 认证中间件
│   ├── models/                      # 数据模型
│   │   └── user.go                  # 用户模型
│   ├── routes/                      # 路由配置
│   │   ├── routes.go                # 主路由配置
│   │   ├── auth_routes.go           # 认证路由
│   │   ├── public_routes.go         # 公开路由
│   │   ├── user_routes.go           # 用户路由
│   │   └── weather_routes.go        # 天气路由
│   ├── utils/                       # 工具函数
│   │   ├── jwt.go                   # JWT 工具
│   │   └── password.go              # 密码加密工具
│   ├── .env                         # 环境变量配置
│   ├── go.mod                       # Go 模块依赖
│   ├── go.sum                       # 依赖版本锁定
│   └── main.go                      # 应用入口
├── client/what-to-wear-client/      # Tauri 前端应用
│   ├── src/                         # React 源码
│   │   ├── AuthPage.tsx             # 认证页面组件
│   │   ├── App.tsx                  # 主应用组件
│   │   ├── App.css                  # 样式文件
│   │   └── main.tsx                 # React 入口
│   ├── src-tauri/                   # Tauri 配置
│   │   ├── src/                     # Rust 源码
│   │   │   ├── lib.rs               # Rust 主库文件
│   │   │   └── main.rs              # Rust 入口文件
│   │   ├── Cargo.toml               # Rust 依赖配置
│   │   └── tauri.conf.json          # Tauri 应用配置
│   ├── package.json                 # Node.js 依赖
│   ├── pnpm-lock.yaml              # 依赖版本锁定
│   └── vite.config.ts               # Vite 构建配置
├── test.html                        # API 测试页面
├── test.json                        # 测试数据
└── claude.md                        # 项目文档 (本文件)
```

## 🚀 快速开始

### 环境要求
- Go 1.23+
- Node.js 18+
- pnpm
- Rust (用于 Tauri)
- PostgreSQL

### 1. 克隆项目
```bash
git clone <repository-url>
cd what-to-wear
```

### 2. 启动后端服务

```bash
# 进入后端目录
cd server

# 安装依赖
go mod tidy

# 配置环境变量 (编辑 .env 文件)
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=your_username
# DB_PASSWORD=your_password
# DB_NAME=what_to_wear
# DB_SSLMODE=disable
# JWT_SECRET=your_jwt_secret

# 启动服务
go run main.go
```

### 3. 启动前端应用

```bash
# 进入前端目录
cd client/what-to-wear-client

# 安装依赖
pnpm install

# 启动开发服务器
pnpm tauri dev
```

## 📡 API 接口文档

### 基础信息
- **Base URL**: `http://localhost:8080/api`
- **认证方式**: Bearer Token (JWT)

### 公开接口

#### 健康检查
```http
GET /ping
```
**响应示例**:
```json
{
  "message": "pong from go server!"
}
```

#### 用户注册
```http
POST /auth/register
Content-Type: application/json

{
  "username": "string (3-20字符)",
  "password": "string (最少6字符)"
}
```

#### 用户登录
```http
POST /auth/login
Content-Type: application/json

{
  "username": "string",
  "password": "string"
}
```
**响应示例**:
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 需要认证的接口

#### 获取用户资料
```http
GET /user/profile
Authorization: Bearer <token>
```

#### 获取当前天气
```http
GET /weather/current
Authorization: Bearer <token>
```

#### 获取天气预报
```http
GET /weather/forecast
Authorization: Bearer <token>
```

## 🦀 Tauri Rust 后端功能

### 适用场景

#### 1. 系统级操作
- **文件系统访问**: 保存用户穿搭照片、读取本地配置
- **系统通知**: 天气变化提醒、穿衣建议推送
- **系统托盘**: 显示当前天气、快速访问功能

#### 2. 硬件访问
- **摄像头调用**: 拍摄今日穿搭记录
- **位置服务**: 获取本地天气数据
- **传感器数据**: 环境温度、湿度检测

#### 3. 性能敏感计算
- **图像处理**: 分析衣服颜色搭配、风格识别
- **算法计算**: 高性能穿搭组合生成
- **数据分析**: 用户穿衣习惯模式分析

#### 4. 离线功能
- **本地缓存**: 天气数据、穿搭建议缓存
- **离线推荐**: 基于历史数据的离线建议
- **数据同步**: 离线数据与云端同步

### 示例 Tauri 命令

```rust
// 保存穿搭记录
#[tauri::command]
async fn save_daily_outfit(
    photo: Vec<u8>,
    weather: WeatherData,
    rating: u8
) -> Result<(), String> {
    // 保存到本地数据库
    let local_db = get_local_database()?;
    local_db.save_outfit_record(photo, weather, rating)?;
    Ok(())
}

// 系统通知
#[tauri::command]
async fn send_weather_notification(message: String) -> Result<(), String> {
    // 发送系统级通知
    show_notification(&message).await?;
    Ok(())
}

// 快速建议
#[tauri::command]
async fn get_quick_recommendation() -> Result<QuickOutfit, String> {
    // 基于缓存数据快速生成建议
    let cached_weather = get_cached_weather()?;
    let outfit = generate_quick_suggestion(cached_weather)?;
    Ok(outfit)
}
```

## 🗄️ 数据库设计

### 用户表 (users)
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔧 开发指南

### 添加新的 Go API 接口

1. **创建控制器方法** (在 `controllers/` 目录)
2. **添加路由配置** (在 `routes/` 目录)
3. **在主路由文件中注册** (`routes/routes.go`)

### 添加新的 Tauri 命令

1. **在 `src-tauri/src/lib.rs` 中定义命令**:
```rust
#[tauri::command]
async fn your_command(param: String) -> Result<String, String> {
    // 实现逻辑
    Ok("success".to_string())
}
```

2. **注册命令到 Tauri**:
```rust
pub fn run() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![your_command])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
```

3. **在前端调用**:
```typescript
import { invoke } from '@tauri-apps/api/core';

const result = await invoke('your_command', { param: 'value' });
```

### 添加新的前端页面

1. **创建 React 组件** (在 `src/` 目录)
2. **添加路由配置** (如果需要)
3. **在主应用中引入**

### 功能开发建议

#### 选择 Go 后端的场景:
- 网络 API 调用 (天气服务、第三方 API)
- 数据库操作 (用户数据、历史记录)
- 用户认证和授权
- 复杂的业务逻辑处理

#### 选择 Rust 后端的场景:
- 文件系统操作 (保存图片、配置文件)
- 系统集成 (通知、托盘、自启动)
- 性能敏感计算 (图像处理、算法)
- 离线功能 (本地缓存、离线推荐)

### 环境变量配置

后端环境变量 (`.env`):
```env
# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=what_to_wear
DB_SSLMODE=disable

# JWT 配置
JWT_SECRET=your_secret_key
```

## 🧪 测试

### API 测试
使用提供的 `test.html` 文件进行 API 测试，或使用 Postman/curl 等工具。

### 前端测试
在 Tauri 应用中直接测试用户界面功能。

## 📦 部署

### 后端部署
```bash
# 构建可执行文件
go build -o what-to-wear-server main.go

# 运行
./what-to-wear-server
```

### 前端部署
```bash
# 构建 Tauri 应用
pnpm tauri build
```

## 🔮 未来计划

### 核心功能
- [ ] 集成真实天气 API (Go 后端)
- [ ] 实现穿衣建议算法 (Rust 后端 - 性能优化)
- [ ] 添加用户偏好设置 (混合实现)
- [ ] 实现历史记录功能 (Go 后端 + 本地缓存)

### Tauri 原生功能
- [ ] 穿搭照片拍摄和管理 (Rust 后端)
- [ ] 系统托盘天气显示 (Rust 后端)
- [ ] 开机自启动和定时提醒 (Rust 后端)
- [ ] 本地数据加密存储 (Rust 后端)
- [ ] 离线模式支持 (Rust 后端)

### 用户体验
- [ ] 服装搭配推荐 (Go 后端 + AI)
- [ ] 智能颜色搭配分析 (Rust 后端 - 图像处理)
- [ ] 个性化穿衣习惯分析 (Rust 后端 - 数据分析)
- [ ] 支持多语言 (前端)
- [ ] 添加主题切换 (前端)

### 高级功能
- [ ] 语音穿衣建议 (Rust 后端 - 系统 TTS)
- [ ] 天气变化智能提醒 (Rust 后端 - 系统通知)
- [ ] 穿搭社交分享功能 (Go 后端)
- [ ] 服装购买建议集成 (Go 后端)

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 📄 许可证

[MIT License](LICENSE)

---

**开发时间**: 2025年8月
**技术支持**: Claude AI Assistant
