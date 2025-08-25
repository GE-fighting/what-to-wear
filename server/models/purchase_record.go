package models

import (
	"time"

	"gorm.io/gorm"
)

// PurchaseRecord 购买记录模型 - 简化版
type PurchaseRecord struct {
	gorm.Model
	ClothingItemID uint      `json:"clothing_item_id" gorm:"not null;uniqueIndex"`
	Price          float64   `json:"price" gorm:"type:decimal(10,2);not null"` // 实际购买价格
	Store          string    `json:"store"`                                    // 商店名称（线上或线下）
	PurchaseDate   time.Time `json:"purchase_date" gorm:"not null"`
	Notes          string    `json:"notes"` // 备注信息（可包含折扣、原价等信息）
}

// TableName 指定表名
func (PurchaseRecord) TableName() string {
	return "purchase_records"
}
