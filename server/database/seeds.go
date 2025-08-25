package database

import (
	"fmt"
	"gorm.io/gorm"
)

// Seeder 种子数据接口
type Seeder interface {
	Seed(db *gorm.DB) error
	GetName() string
}

// CategorySeeder 分类种子数据
type CategorySeeder struct{}

func (s CategorySeeder) Seed(db *gorm.DB) error {
	return SeedCategories(db)
}

func (s CategorySeeder) GetName() string {
	return "Categories"
}

// TagSeeder 标签种子数据
type TagSeeder struct{}

func (s TagSeeder) Seed(db *gorm.DB) error {
	return SeedTags(db)
}

func (s TagSeeder) GetName() string {
	return "Tags"
}

// RunAllSeeders 运行所有种子数据
func RunAllSeeders(db *gorm.DB) error {
	seeders := []Seeder{
		CategorySeeder{},
		TagSeeder{},
	}

	for _, seeder := range seeders {
		fmt.Printf("正在运行种子数据: %s...\n", seeder.GetName())
		if err := seeder.Seed(db); err != nil {
			return fmt.Errorf("运行种子数据 %s 失败: %v", seeder.GetName(), err)
		}
		fmt.Printf("种子数据 %s 运行成功\n", seeder.GetName())
	}

	return nil
}

// RunSeeder 运行指定的种子数据
func RunSeeder(db *gorm.DB, seederName string) error {
	seeders := map[string]Seeder{
		"categories": CategorySeeder{},
		"tags":       TagSeeder{},
	}

	seeder, exists := seeders[seederName]
	if !exists {
		return fmt.Errorf("种子数据 %s 不存在", seederName)
	}

	fmt.Printf("正在运行种子数据: %s...\n", seeder.GetName())
	if err := seeder.Seed(db); err != nil {
		return fmt.Errorf("运行种子数据 %s 失败: %v", seeder.GetName(), err)
	}
	fmt.Printf("种子数据 %s 运行成功\n", seeder.GetName())

	return nil
}