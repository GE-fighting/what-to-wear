package models

import (
	"what-to-wear/server/api"

	"gorm.io/gorm"
)

// ClothingTag 衣物标签模型
type ClothingTag struct {
	gorm.Model
	Name        string      `json:"name" gorm:"not null;index"`
	Type        api.TagType `json:"type" gorm:"not null;index"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
	IsSystem    bool        `json:"is_system" gorm:"default:false"` // 是否为系统预设标签
	IsActive    bool        `json:"is_active" gorm:"default:true"`
	SortOrder   int         `json:"sort_order" gorm:"default:0"`
	UserID      *uint       `json:"user_id" gorm:"index"` // 自定义标签的创建者
}

// TableName 指定表名
func (ClothingTag) TableName() string {
	return "clothing_tags"
}

// ClothingItemTag 衣物标签关联表
type ClothingItemTag struct {
	gorm.Model
	ClothingItemID uint `json:"clothing_item_id" gorm:"primaryKey"`
	ClothingTagID  uint `json:"clothing_tag_id" gorm:"primaryKey"`
}

// TableName 指定表名
func (ClothingItemTag) TableName() string {
	return "clothing_item_tags"
}
