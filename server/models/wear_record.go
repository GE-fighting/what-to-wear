package models

import (
	"time"

	"gorm.io/gorm"
)

// WearRecord 穿着记录模型 - 简化版
type WearRecord struct {
	gorm.Model
	ClothingItemID uint      `json:"clothing_item_id" gorm:"not null;index"`
	WearDate       time.Time `json:"wear_date" gorm:"not null"`
	Notes          string    `json:"notes"` // 备注信息（可包含场合、天气、评分等）
}

// TableName 指定表名
func (WearRecord) TableName() string {
	return "wear_records"
}
