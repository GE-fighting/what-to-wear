# 数据库管理系统

## 概述

数据库管理系统采用分层设计，将数据库连接、迁移和种子数据分离，提供清晰的职责划分和灵活的使用方式。

## 目录结构

```
database/
├── README.md          # 说明文档
├── database.go        # 统一的数据库初始化入口
└── migrations.go      # 数据库迁移管理

config/
└── database.go        # 数据库连接配置

seeds/
├── seeder.go          # 种子数据管理器
├── categories.go      # 分类种子数据
└── tags.go           # 标签种子数据

cmd/
├── migrate/           # 数据库迁移命令行工具
│   └── main.go
└── seed/              # 种子数据命令行工具
    └── main.go
```

## 架构设计

### 1. 职责分离

- **config/database.go** - 数据库连接配置和连接管理
- **database/migrations.go** - 数据库表结构迁移
- **database/database.go** - 统一的数据库初始化入口
- **seeds/** - 种子数据管理

### 2. 分层结构

```
应用层 (main.go)
    ↓
数据库初始化层 (database/database.go)
    ↓
配置层 (config/database.go) + 迁移层 (database/migrations.go) + 种子数据层 (seeds/)
```

## 使用方式

### 1. 应用启动时自动初始化

```go
// main.go
func main() {
    // 初始化数据库（连接 + 迁移 + 种子数据）
    if err := database.Initialize(); err != nil {
        log.Fatal("Database initialization failed:", err)
    }
    
    // 获取数据库实例
    db := database.GetDB()
    
    // 启动应用...
}
```

### 2. 不包含种子数据的初始化

```go
// 只进行连接和迁移，不运行种子数据
if err := database.InitializeWithoutSeeds(); err != nil {
    log.Fatal("Database initialization failed:", err)
}
```

### 3. 命令行工具

**数据库迁移工具:**
```bash
# 运行数据库迁移
go run cmd/migrate/main.go -action migrate

# 运行种子数据
go run cmd/migrate/main.go -action seed

# 运行指定种子数据
go run cmd/migrate/main.go -action seed -seeder categories

# 重置数据库（删除所有表并重新初始化）
go run cmd/migrate/main.go -action reset

# 检查数据库状态
go run cmd/migrate/main.go -action status

# 删除所有表（危险操作）
go run cmd/migrate/main.go -action drop
```

**种子数据工具:**
```bash
# 运行所有种子数据
go run cmd/seed/main.go

# 运行指定种子数据
go run cmd/seed/main.go -seeder categories
```

### 4. 程序中调用

```go
import (
    "what-to-wear/server/database"
    "what-to-wear/server/seeds"
)

// 只运行迁移
err := database.AutoMigrate(db)

// 只运行种子数据
err := database.SeedInitialData(db)

// 运行指定种子数据
err := seeds.RunSeeder(db, "categories")

// 数据库健康检查
err := database.HealthCheck()

// 重置数据库
err := database.Reset()
```

## 配置管理

### 环境变量

```bash
# 数据库连接配置
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=what_to_wear
DB_PORT=5432
DB_SSLMODE=disable

# 日志配置
ENABLE_SQL_LOGGING=true
LOG_LEVEL=debug
GIN_MODE=debug
```

### 配置结构

```go
type DatabaseConfig struct {
    Host     string
    User     string
    Password string
    DBName   string
    Port     string
    SSLMode  string
}
```

## 数据库迁移

### 自动迁移

系统会自动迁移以下模型：

- User (用户)
- ClothingCategory (衣物分类)
- ClothingTag (衣物标签)
- ClothingItem (衣物)
- Outfit (穿搭)
- OutfitItem (穿搭项目)
- WearRecord (穿着记录)
- MaintenanceRecord (保养记录)
- PurchaseRecord (购买记录)
- Attachment (附件)

### 迁移管理

```go
// 迁移所有模型
database.AutoMigrate(db)

// 迁移指定模型
database.MigrateSpecificModels(db, &models.User{}, &models.ClothingItem{})

// 检查迁移状态
database.CheckMigrationStatus(db)

// 删除所有表（开发环境）
database.DropAllTables(db)
```

## 种子数据管理

### 数据类型

- **分类数据**: 6个一级分类，30+个二级分类
- **标签数据**: 季节、场合、风格、颜色、材质、品牌等标签

### 运行方式

```go
// 运行所有种子数据
seeds.RunAllSeeders(db)

// 运行指定种子数据
seeds.RunSeeder(db, "categories")
seeds.RunSeeder(db, "tags")
```

## 最佳实践

### 1. 开发环境

```go
// 开发时可以重置数据库
if gin.Mode() == gin.DebugMode {
    database.Reset()
} else {
    database.Initialize()
}
```

### 2. 生产环境

```go
// 生产环境只初始化，不重置
database.Initialize()
```

### 3. 测试环境

```go
func TestSetup(t *testing.T) {
    // 使用内存数据库或测试数据库
    db := setupTestDB()
    
    // 运行迁移和种子数据
    database.AutoMigrate(db)
    seeds.RunAllSeeders(db)
}
```

### 4. 错误处理

```go
if err := database.Initialize(); err != nil {
    log.Fatal("Database initialization failed:", err)
}

// 健康检查
if err := database.HealthCheck(); err != nil {
    log.Error("Database health check failed:", err)
}
```

## 监控和维护

### 健康检查

```go
// 检查数据库连接
err := database.HealthCheck()

// 检查迁移状态
err := database.CheckMigrationStatus(db)
```

### 日志记录

- 所有数据库操作都有详细的日志记录
- 支持SQL查询日志（开发环境）
- 错误日志包含详细的错误信息

### 性能优化

- 连接池配置
- 查询优化
- 索引管理
- 事务处理

## 故障排除

### 常见问题

1. **连接失败**: 检查数据库配置和网络连接
2. **迁移失败**: 检查模型定义和数据库权限
3. **种子数据失败**: 检查数据完整性和约束条件

### 恢复操作

```bash
# 重置数据库
go run cmd/migrate/main.go -action reset

# 只运行迁移
go run cmd/migrate/main.go -action migrate

# 只运行种子数据
go run cmd/migrate/main.go -action seed
```

---

**维护者**: What to Wear 开发团队  
**最后更新**: 2025年8月19日