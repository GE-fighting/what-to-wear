package database

import (
	"fmt"
	"gorm.io/gorm"
	"what-to-wear/server/models"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	fmt.Println("开始数据库迁移...")

	// 迁移所有模型
	err := db.AutoMigrate(
		&models.User{},
		&models.ClothingCategory{},
		&models.ClothingTag{},
		&models.ClothingItem{},
		&models.Outfit{},
		&models.OutfitItem{},
		&models.WearRecord{},
		&models.MaintenanceRecord{},
		&models.PurchaseRecord{},
		&models.Attachment{},
	)

	if err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	fmt.Println("数据库迁移完成")
	return nil
}

// MigrateSpecificModels 迁移指定的模型
func MigrateSpecificModels(db *gorm.DB, models ...interface{}) error {
	fmt.Printf("开始迁移指定模型 (%d个)...\n", len(models))

	err := db.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("指定模型迁移失败: %v", err)
	}

	fmt.Println("指定模型迁移完成")
	return nil
}

// DropAllTables 删除所有表（危险操作，仅用于开发环境）
func DropAllTables(db *gorm.DB) error {
	fmt.Println("警告: 正在删除所有表...")

	// 按依赖关系逆序删除表
	tables := []interface{}{
		&models.Attachment{},
		&models.PurchaseRecord{},
		&models.MaintenanceRecord{},
		&models.WearRecord{},
		&models.OutfitItem{},
		&models.Outfit{},
		&models.ClothingItem{},
		&models.ClothingTag{},
		&models.ClothingCategory{},
		&models.User{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("删除表失败: %v", err)
		}
	}

	fmt.Println("所有表删除完成")
	return nil
}

// CheckMigrationStatus 检查迁移状态
func CheckMigrationStatus(db *gorm.DB) error {
	fmt.Println("检查数据库迁移状态...")

	models := []interface{}{
		&models.User{},
		&models.ClothingCategory{},
		&models.ClothingTag{},
		&models.ClothingItem{},
		&models.Outfit{},
		&models.OutfitItem{},
		&models.WearRecord{},
		&models.MaintenanceRecord{},
		&models.PurchaseRecord{},
		&models.Attachment{},
	}

	for _, model := range models {
		if !db.Migrator().HasTable(model) {
			return fmt.Errorf("表 %T 不存在，需要运行迁移", model)
		}
	}

	fmt.Println("所有表都存在，迁移状态正常")
	return nil
}