package models

import (
	"gorm.io/gorm"
)

// OutfitItem 穿搭单品模型
type OutfitItem struct {
	gorm.Model
	OutfitID       uint   `json:"outfit_id" gorm:"not null;index"`
	ClothingItemID uint   `json:"clothing_item_id" gorm:"not null;index"`
	ItemRole       string `json:"item_role"`                        // 在穿搭中的角色 (main, accent, base, etc.)
	LayerOrder     int    `json:"layer_order" gorm:"default:1"`     // 穿着层次顺序
	IsOptional     bool   `json:"is_optional" gorm:"default:false"` // 是否为可选单品
	Notes          string `json:"notes"`
}

// TableName 指定表名
func (OutfitItem) TableName() string {
	return "outfit_items"
}
