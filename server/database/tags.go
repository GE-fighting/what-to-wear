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
		{Name: "春季", Type: "season", Color: "#90EE90", Description: "适合春季穿着", SortOrder: 1},
		{Name: "夏季", Type: "season", Color: "#FFD700", Description: "适合夏季穿着", SortOrder: 2},
		{Name: "秋季", Type: "season", Color: "#FFA500", Description: "适合秋季穿着", SortOrder: 3},
		{Name: "冬季", Type: "season", Color: "#87CEEB", Description: "适合冬季穿着", SortOrder: 4},

		// 场合标签
		{Name: "正式", Type: "occasion", Color: "#2F4F4F", Description: "正式场合", SortOrder: 1},
		{Name: "休闲", Type: "occasion", Color: "#32CD32", Description: "休闲场合", SortOrder: 2},
		{Name: "运动", Type: "occasion", Color: "#FF6347", Description: "运动健身", SortOrder: 3},
		{Name: "约会", Type: "occasion", Color: "#FF69B4", Description: "约会场合", SortOrder: 4},
		{Name: "工作", Type: "occasion", Color: "#4682B4", Description: "工作场合", SortOrder: 5},
		{Name: "聚会", Type: "occasion", Color: "#9370DB", Description: "聚会场合", SortOrder: 6},
		{Name: "居家", Type: "occasion", Color: "#DDA0DD", Description: "居家休息", SortOrder: 7},

		// 风格标签
		{Name: "简约", Type: "style", Color: "#708090", Description: "简约风格", SortOrder: 1},
		{Name: "甜美", Type: "style", Color: "#FFB6C1", Description: "甜美风格", SortOrder: 2},
		{Name: "酷帅", Type: "style", Color: "#2F4F4F", Description: "酷帅风格", SortOrder: 3},
		{Name: "优雅", Type: "style", Color: "#9370DB", Description: "优雅风格", SortOrder: 4},
		{Name: "复古", Type: "style", Color: "#CD853F", Description: "复古风格", SortOrder: 5},
		{Name: "街头", Type: "style", Color: "#FF4500", Description: "街头风格", SortOrder: 6},
		{Name: "文艺", Type: "style", Color: "#8FBC8F", Description: "文艺风格", SortOrder: 7},

		// 颜色标签
		{Name: "黑色", Type: "color", Color: "#000000", Description: "黑色系", SortOrder: 1},
		{Name: "白色", Type: "color", Color: "#FFFFFF", Description: "白色系", SortOrder: 2},
		{Name: "灰色", Type: "color", Color: "#808080", Description: "灰色系", SortOrder: 3},
		{Name: "红色", Type: "color", Color: "#FF0000", Description: "红色系", SortOrder: 4},
		{Name: "蓝色", Type: "color", Color: "#0000FF", Description: "蓝色系", SortOrder: 5},
		{Name: "绿色", Type: "color", Color: "#008000", Description: "绿色系", SortOrder: 6},
		{Name: "黄色", Type: "color", Color: "#FFFF00", Description: "黄色系", SortOrder: 7},
		{Name: "粉色", Type: "color", Color: "#FFC0CB", Description: "粉色系", SortOrder: 8},
		{Name: "紫色", Type: "color", Color: "#800080", Description: "紫色系", SortOrder: 9},
		{Name: "棕色", Type: "color", Color: "#A52A2A", Description: "棕色系", SortOrder: 10},

		// 材质标签
		{Name: "棉质", Type: "material", Color: "#F5F5DC", Description: "棉质面料", SortOrder: 1},
		{Name: "丝绸", Type: "material", Color: "#FFE4E1", Description: "丝绸面料", SortOrder: 2},
		{Name: "羊毛", Type: "material", Color: "#F0E68C", Description: "羊毛面料", SortOrder: 3},
		{Name: "牛仔", Type: "material", Color: "#4169E1", Description: "牛仔面料", SortOrder: 4},
		{Name: "皮革", Type: "material", Color: "#8B4513", Description: "皮革材质", SortOrder: 5},
		{Name: "雪纺", Type: "material", Color: "#F8F8FF", Description: "雪纺面料", SortOrder: 6},
		{Name: "针织", Type: "material", Color: "#DDA0DD", Description: "针织面料", SortOrder: 7},
		{Name: "涤纶", Type: "material", Color: "#E6E6FA", Description: "涤纶面料", SortOrder: 8},

		// 品牌标签（示例）
		{Name: "优衣库", Type: "brand", Color: "#FF0000", Description: "UNIQLO品牌", SortOrder: 1},
		{Name: "ZARA", Type: "brand", Color: "#000000", Description: "ZARA品牌", SortOrder: 2},
		{Name: "H&M", Type: "brand", Color: "#DC143C", Description: "H&M品牌", SortOrder: 3},
		{Name: "Nike", Type: "brand", Color: "#FF6347", Description: "Nike品牌", SortOrder: 4},
		{Name: "Adidas", Type: "brand", Color: "#000000", Description: "Adidas品牌", SortOrder: 5},
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
				Color:       seed.Color,
				Description: seed.Description,
				SortOrder:   seed.SortOrder,
				IsActive:    true,
			}

			if err := db.Create(&newTag).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
