package models

import (
	"gorm.io/gorm"
)

// ClothingCategory 衣物分类模型
type ClothingCategory struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;uniqueIndex"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parent_id" gorm:"index"`
	Icon        string `json:"icon"`
	SortOrder   int    `json:"sort_order" gorm:"default:0"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

// TableName 指定表名
func (ClothingCategory) TableName() string {
	return "clothing_categories"
}

// IsRootCategory 检查是否为根分类
func (c *ClothingCategory) IsRootCategory() bool {
	return c.ParentID == nil
}

// HasParent 检查是否有父分类
func (c *ClothingCategory) HasParent() bool {
	return c.ParentID != nil
}
