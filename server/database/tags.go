package database

import (
	"what-to-wear/server/api"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// TagSeed 标签种子数据结构
type TagSeed struct {
	Name        string
	Type        string
	Color       string
	Description string
	SortOrder   int
}

// GetTagSeeds 获取标签种子数据
func GetTagSeeds() []TagSeed {
	return []TagSeed{
		// 季节标签
		{Name: "春季", Type: "season", Description: "适合春季穿着", SortOrder: 1},
		{Name: "夏季", Type: "season", Description: "适合夏季穿着", SortOrder: 2},
		{Name: "秋季", Type: "season", Description: "适合秋季穿着", SortOrder: 3},
		{Name: "冬季", Type: "season", Description: "适合冬季穿着", SortOrder: 4},

		// 场合标签
		{Name: "正式", Type: "occasion", Description: "正式场合", SortOrder: 1},
		{Name: "休闲", Type: "occasion", Description: "休闲场合", SortOrder: 2},
		{Name: "运动", Type: "occasion", Description: "运动健身", SortOrder: 3},
		{Name: "约会", Type: "occasion", Description: "约会场合", SortOrder: 4},
		{Name: "工作", Type: "occasion", Description: "工作场合", SortOrder: 5},
		{Name: "聚会", Type: "occasion", Description: "聚会场合", SortOrder: 6},
		{Name: "居家", Type: "occasion", Description: "居家休息", SortOrder: 7},

		// 风格标签
		{Name: "简约", Type: "style", Description: "简约风格", SortOrder: 1},
		{Name: "甜美", Type: "style", Description: "甜美风格", SortOrder: 2},
		{Name: "酷帅", Type: "style", Description: "酷帅风格", SortOrder: 3},
		{Name: "优雅", Type: "style", Description: "优雅风格", SortOrder: 4},
		{Name: "复古", Type: "style", Description: "复古风格", SortOrder: 5},
		{Name: "街头", Type: "style", Description: "街头风格", SortOrder: 6},
		{Name: "文艺", Type: "style", Description: "文艺风格", SortOrder: 7},
	}
}

// SeedTags 初始化标签数据
func SeedTags(db *gorm.DB) error {
	tagSeeds := GetTagSeeds()

	for _, seed := range tagSeeds {
		var existingTag models.ClothingTag
		err := db.Where("name = ? AND type = ?", seed.Name, seed.Type).First(&existingTag).Error

		if err == gorm.ErrRecordNotFound {
			newTag := models.ClothingTag{
				Name:        seed.Name,
				Type:        api.TagType(seed.Type),
				Description: seed.Description,
				SortOrder:   seed.SortOrder,
				IsSystem:    true,
				IsActive:    true,
			}

			if err := db.Create(&newTag).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
