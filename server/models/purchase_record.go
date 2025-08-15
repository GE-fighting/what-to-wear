package models

import (
	"time"

	"gorm.io/gorm"
)

// PurchaseRecord 购买记录模型
type PurchaseRecord struct {
	gorm.Model
	ClothingItemID uint       `json:"clothing_item_id" gorm:"not null;uniqueIndex"`
	PurchasePrice  float64    `json:"purchase_price" gorm:"type:decimal(10,2);not null"`
	OriginalPrice  float64    `json:"original_price" gorm:"type:decimal(10,2)"`    // 原价
	Discount       float64    `json:"discount" gorm:"type:decimal(5,2);default:0"` // 折扣百分比
	StoreName      string     `json:"store_name"`
	StoreLocation  string     `json:"store_location"`
	OnlineStore    string     `json:"online_store"` // 网店名称
	OrderNumber    string     `json:"order_number"` // 订单号
	PurchaseDate   time.Time  `json:"purchase_date" gorm:"not null"`
	PaymentMethod  string     `json:"payment_method"` // 支付方式
	Currency       string     `json:"currency" gorm:"default:'CNY'"`
	ReceiptURL     string     `json:"receipt_url"`     // 收据图片URL
	WarrantyPeriod int        `json:"warranty_period"` // 保修期（月）
	WarrantyExpiry *time.Time `json:"warranty_expiry"` // 保修到期时间
	ReturnPolicy   string     `json:"return_policy"`   // 退换货政策
	Notes          string     `json:"notes"`
	Tags           []string   `json:"tags" gorm:"type:json"` // 购买标签（如：打折、新品、限量等）
}

// TableName 指定表名
func (PurchaseRecord) TableName() string {
	return "purchase_records"
}

// CalculateDiscount 计算折扣
func (p *PurchaseRecord) CalculateDiscount() float64 {
	if p.OriginalPrice <= 0 {
		return 0
	}
	return ((p.OriginalPrice - p.PurchasePrice) / p.OriginalPrice) * 100
}

// IsUnderWarranty 检查是否在保修期内
func (p *PurchaseRecord) IsUnderWarranty() bool {
	if p.WarrantyExpiry == nil {
		return false
	}
	return time.Now().Before(*p.WarrantyExpiry)
}

// CalculateWarrantyExpiry 计算保修到期时间
func (p *PurchaseRecord) CalculateWarrantyExpiry() {
	if p.WarrantyPeriod > 0 {
		expiry := p.PurchaseDate.AddDate(0, p.WarrantyPeriod, 0)
		p.WarrantyExpiry = &expiry
	}
}

// BeforeCreate GORM钩子：创建前自动计算相关字段
func (p *PurchaseRecord) BeforeCreate(tx *gorm.DB) error {
	// 自动计算折扣
	if p.OriginalPrice > 0 && p.Discount == 0 {
		p.Discount = p.CalculateDiscount()
	}

	// 自动计算保修到期时间
	p.CalculateWarrantyExpiry()

	return nil
}

// BeforeUpdate GORM钩子：更新前自动计算相关字段
func (p *PurchaseRecord) BeforeUpdate(tx *gorm.DB) error {
	// 重新计算折扣
	if p.OriginalPrice > 0 {
		p.Discount = p.CalculateDiscount()
	}

	// 重新计算保修到期时间
	p.CalculateWarrantyExpiry()

	return nil
}
