package models

import (
	"time"

	"gorm.io/gorm"
)

// TagType 标签类型枚举
type TagType string

const (
	TagTypeSeason   TagType = "season"   // 季节标签
	TagTypeOccasion TagType = "occasion" // 场合标签
	TagTypeStyle    TagType = "style"    // 风格标签
	TagTypeColor    TagType = "color"    // 颜色标签
	TagTypeMaterial TagType = "material" // 材质标签
	TagTypeBrand    TagType = "brand"    // 品牌标签
	TagTypeCustom   TagType = "custom"   // 自定义标签
)

// ClothingTag 衣物标签模型
type ClothingTag struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null;index"`
	Type        TagType `json:"type" gorm:"not null;index"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	Icon        string  `json:"icon"`
	IsSystem    bool    `json:"is_system" gorm:"default:false"` // 是否为系统预设标签
	IsActive    bool    `json:"is_active" gorm:"default:true"`
	SortOrder   int     `json:"sort_order" gorm:"default:0"`
	UserID      *uint   `json:"user_id" gorm:"index"` // 自定义标签的创建者
}

// TableName 指定表名
func (ClothingTag) TableName() string {
	return "clothing_tags"
}

// ClothingItemTag 衣物标签关联表
type ClothingItemTag struct {
	ClothingItemID uint `json:"clothing_item_id" gorm:"primaryKey"`
	ClothingTagID  uint `json:"clothing_tag_id" gorm:"primaryKey"`
	// 移除直接关联以避免循环依赖
	CreatedAt time.Time `json:"created_at"`
}

// TableName 指定表名
func (ClothingItemTag) TableName() string {
	return "clothing_item_tags"
}

// GetTagsByType 根据类型获取标签
func GetTagsByType(db *gorm.DB, tagType TagType, userID *uint) ([]ClothingTag, error) {
	var tags []ClothingTag
	query := db.Where("type = ? AND is_active = ?", tagType, true)

	// 包含系统标签和用户自定义标签
	if userID != nil {
		query = query.Where("is_system = ? OR user_id = ?", true, *userID)
	} else {
		query = query.Where("is_system = ?", true)
	}

	err := query.Order("sort_order ASC, name ASC").Find(&tags).Error
	return tags, err
}

// IsValidTagType 检查标签类型是否有效
func IsValidTagType(tagType string) bool {
	validTypes := []TagType{
		TagTypeSeason, TagTypeOccasion, TagTypeStyle,
		TagTypeColor, TagTypeMaterial, TagTypeBrand, TagTypeCustom,
	}

	for _, validType := range validTypes {
		if TagType(tagType) == validType {
			return true
		}
	}
	return false
}
