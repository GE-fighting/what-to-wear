package database

import (
	"fmt"
	"gorm.io/gorm"
	"what-to-wear/server/config"
)

// DB 全局数据库实例
var DB *gorm.DB

// Initialize 初始化数据库（连接 + 迁移 + 种子数据）
func Initialize() error {
	// 1. 加载数据库配置
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		return fmt.Errorf("加载数据库配置失败: %v", err)
	}

	// 2. 连接数据库
	db, err := config.ConnectDatabase(dbConfig)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 3. 运行数据库迁移
	if err := AutoMigrate(db); err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	// 4. 初始化种子数据
	if err := SeedInitialData(db); err != nil {
		return fmt.Errorf("种子数据初始化失败: %v", err)
	}

	// 5. 设置全局数据库实例
	DB = db
	fmt.Println("数据库初始化完成")

	return nil
}

// InitializeWithoutSeeds 初始化数据库（不包含种子数据）
func InitializeWithoutSeeds() error {
	// 1. 加载数据库配置
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		return fmt.Errorf("加载数据库配置失败: %v", err)
	}

	// 2. 连接数据库
	db, err := config.ConnectDatabase(dbConfig)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 3. 运行数据库迁移
	if err := AutoMigrate(db); err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	// 4. 设置全局数据库实例
	DB = db
	fmt.Println("数据库初始化完成（不包含种子数据）")

	return nil
}

// SeedInitialData 初始化种子数据
func SeedInitialData(db *gorm.DB) error {
	fmt.Println("开始初始化种子数据...")

	// 运行所有种子数据
	if err := RunAllSeeders(db); err != nil {
		return fmt.Errorf("初始化种子数据失败: %v", err)
	}

	fmt.Println("种子数据初始化完成")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// Close 关闭数据库连接
func Close() error {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("关闭数据库连接失败: %v", err)
	}

	fmt.Println("数据库连接已关闭")
	return nil
}

// Reset 重置数据库（删除所有表并重新初始化）
func Reset() error {
	if DB == nil {
		return fmt.Errorf("数据库未初始化")
	}

	fmt.Println("开始重置数据库...")

	// 1. 删除所有表
	if err := DropAllTables(DB); err != nil {
		return fmt.Errorf("删除表失败: %v", err)
	}

	// 2. 重新迁移
	if err := AutoMigrate(DB); err != nil {
		return fmt.Errorf("重新迁移失败: %v", err)
	}

	// 3. 重新初始化种子数据
	if err := SeedInitialData(DB); err != nil {
		return fmt.Errorf("重新初始化种子数据失败: %v", err)
	}

	fmt.Println("数据库重置完成")
	return nil
}

// HealthCheck 数据库健康检查
func HealthCheck() error {
	if DB == nil {
		return fmt.Errorf("数据库未初始化")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %v", err)
	}

	return nil
}