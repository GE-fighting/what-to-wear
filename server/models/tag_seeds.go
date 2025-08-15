package models

import (
	"gorm.io/gorm"
)

// SystemTag 系统预设标签结构
type SystemTag struct {
	Name        string
	Type        TagType
	Description string
	Color       string
	Icon        string
	SortOrder   int
}

// GetSystemTags 获取系统预设标签
func GetSystemTags() []SystemTag {
	return []SystemTag{
		// 季节标签
		{Name: "春季", Type: TagTypeSeason, Description: "适合春季穿着", Color: "#90EE90", Icon: "🌸", SortOrder: 1},
		{Name: "夏季", Type: TagTypeSeason, Description: "适合夏季穿着", Color: "#FFD700", Icon: "☀️", SortOrder: 2},
		{Name: "秋季", Type: TagTypeSeason, Description: "适合秋季穿着", Color: "#DEB887", Icon: "🍂", SortOrder: 3},
		{Name: "冬季", Type: TagTypeSeason, Description: "适合冬季穿着", Color: "#87CEEB", Icon: "❄️", SortOrder: 4},

		// 场合标签
		{Name: "休闲", Type: TagTypeOccasion, Description: "日常休闲场合", Color: "#98FB98", Icon: "🏠", SortOrder: 1},
		{Name: "正式", Type: TagTypeOccasion, Description: "正式商务场合", Color: "#4682B4", Icon: "💼", SortOrder: 2},
		{Name: "运动", Type: TagTypeOccasion, Description: "运动健身场合", Color: "#FF6347", Icon: "🏃", SortOrder: 3},
		{Name: "派对", Type: TagTypeOccasion, Description: "聚会派对场合", Color: "#FF69B4", Icon: "🎉", SortOrder: 4},
		{Name: "工作", Type: TagTypeOccasion, Description: "工作办公场合", Color: "#708090", Icon: "🏢", SortOrder: 5},
		{Name: "约会", Type: TagTypeOccasion, Description: "约会浪漫场合", Color: "#FFB6C1", Icon: "💕", SortOrder: 6},
		{Name: "旅行", Type: TagTypeOccasion, Description: "旅行出游场合", Color: "#20B2AA", Icon: "✈️", SortOrder: 7},

		// 风格标签
		{Name: "简约", Type: TagTypeStyle, Description: "简约现代风格", Color: "#F5F5F5", Icon: "⚪", SortOrder: 1},
		{Name: "复古", Type: TagTypeStyle, Description: "复古怀旧风格", Color: "#D2691E", Icon: "📻", SortOrder: 2},
		{Name: "街头", Type: TagTypeStyle, Description: "街头潮流风格", Color: "#FF4500", Icon: "🛹", SortOrder: 3},
		{Name: "优雅", Type: TagTypeStyle, Description: "优雅精致风格", Color: "#DDA0DD", Icon: "👑", SortOrder: 4},
		{Name: "运动风", Type: TagTypeStyle, Description: "运动休闲风格", Color: "#32CD32", Icon: "⚽", SortOrder: 5},
		{Name: "商务", Type: TagTypeStyle, Description: "商务专业风格", Color: "#2F4F4F", Icon: "📊", SortOrder: 6},
		{Name: "甜美", Type: TagTypeStyle, Description: "甜美可爱风格", Color: "#FFB6C1", Icon: "🎀", SortOrder: 7},
		{Name: "朋克", Type: TagTypeStyle, Description: "朋克摇滚风格", Color: "#000000", Icon: "🎸", SortOrder: 8},

		// 颜色标签
		{Name: "基础色", Type: TagTypeColor, Description: "黑白灰等基础色", Color: "#808080", Icon: "⚫", SortOrder: 1},
		{Name: "亮色", Type: TagTypeColor, Description: "鲜艳明亮色彩", Color: "#FF0000", Icon: "🔴", SortOrder: 2},
		{Name: "暗色", Type: TagTypeColor, Description: "深沉暗淡色彩", Color: "#2F2F2F", Icon: "⚫", SortOrder: 3},
		{Name: "中性色", Type: TagTypeColor, Description: "中性自然色彩", Color: "#8FBC8F", Icon: "🟤", SortOrder: 4},

		// 材质标签
		{Name: "棉质", Type: TagTypeMaterial, Description: "纯棉材质", Color: "#F0E68C", Icon: "🌿", SortOrder: 1},
		{Name: "丝质", Type: TagTypeMaterial, Description: "真丝材质", Color: "#DDA0DD", Icon: "✨", SortOrder: 2},
		{Name: "羊毛", Type: TagTypeMaterial, Description: "羊毛材质", Color: "#F5DEB3", Icon: "🐑", SortOrder: 3},
		{Name: "皮革", Type: TagTypeMaterial, Description: "真皮材质", Color: "#8B4513", Icon: "🦬", SortOrder: 4},
		{Name: "牛仔", Type: TagTypeMaterial, Description: "牛仔布料", Color: "#4169E1", Icon: "👖", SortOrder: 5},
		{Name: "聚酯纤维", Type: TagTypeMaterial, Description: "化纤材质", Color: "#B0C4DE", Icon: "🧵", SortOrder: 6},

		// 品牌标签（示例）
		{Name: "奢侈品牌", Type: TagTypeBrand, Description: "高端奢侈品牌", Color: "#FFD700", Icon: "💎", SortOrder: 1},
		{Name: "快时尚", Type: TagTypeBrand, Description: "快时尚品牌", Color: "#FF69B4", Icon: "⚡", SortOrder: 2},
		{Name: "设计师品牌", Type: TagTypeBrand, Description: "独立设计师品牌", Color: "#9370DB", Icon: "🎨", SortOrder: 3},
		{Name: "运动品牌", Type: TagTypeBrand, Description: "专业运动品牌", Color: "#32CD32", Icon: "🏃", SortOrder: 4},
	}
}

// SeedSystemTags 初始化系统预设标签
func SeedSystemTags(db *gorm.DB) error {
	systemTags := GetSystemTags()

	for _, tag := range systemTags {
		// 检查标签是否已存在
		var existingTag ClothingTag
		err := db.Where("name = ? AND type = ? AND is_system = ?", tag.Name, tag.Type, true).First(&existingTag).Error

		if err == gorm.ErrRecordNotFound {
			// 标签不存在，创建新标签
			newTag := ClothingTag{
				Name:        tag.Name,
				Type:        tag.Type,
				Description: tag.Description,
				Color:       tag.Color,
				Icon:        tag.Icon,
				IsSystem:    true,
				IsActive:    true,
				SortOrder:   tag.SortOrder,
			}

			if err := db.Create(&newTag).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// GetSystemTagsByType 根据类型获取系统预设标签
func GetSystemTagsByType(tagType TagType) []SystemTag {
	allTags := GetSystemTags()
	var filteredTags []SystemTag

	for _, tag := range allTags {
		if tag.Type == tagType {
			filteredTags = append(filteredTags, tag)
		}
	}

	return filteredTags
}
