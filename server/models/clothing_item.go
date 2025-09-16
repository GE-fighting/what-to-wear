package models

import (
	"math"
	"time"

	"what-to-wear/server/api"

	"gorm.io/gorm"
)

// ClothingSize 衣物尺码
type ClothingSize struct {
	Size   string `json:"size"`   // 尺码标识 (XS, S, M, L, XL, 38, 40, etc.)
	System string `json:"system"` // 尺码系统 (US, EU, CN, UK, etc.)
}

// SpecificAttributes 特定属性结构
type SpecificAttributes struct {
	// 通用属性
	Sleeve    string `json:"sleeve,omitempty"`    // 袖长 (长袖、短袖、无袖)
	Neckline  string `json:"neckline,omitempty"`  // 领型 (圆领、V领、高领)
	Pattern   string `json:"pattern,omitempty"`   // 图案 (纯色、条纹、格子)
	Thickness string `json:"thickness,omitempty"` // 厚度 (薄、中等、厚)

	// 上衣特有
	Fit    string `json:"fit,omitempty"`    // 版型 (修身、宽松、标准)
	Length string `json:"length,omitempty"` // 长度 (短款、中长款、长款)

	// 裤子特有
	Rise   string `json:"rise,omitempty"`   // 腰线 (高腰、中腰、低腰)
	Leg    string `json:"leg,omitempty"`    // 裤型 (直筒、修身、阔腿)
	Inseam int    `json:"inseam,omitempty"` // 内长 (cm)

	// 鞋子特有
	HeelHeight float64 `json:"heel_height,omitempty"` // 跟高 (cm)
	ShoeType   string  `json:"shoe_type,omitempty"`   // 鞋型 (运动鞋、皮鞋、靴子)
	Closure    string  `json:"closure,omitempty"`     // 闭合方式 (系带、拉链、套脚)

	// 配饰特有
	Width    float64 `json:"width,omitempty"`    // 宽度 (cm)
	Diameter float64 `json:"diameter,omitempty"` // 直径 (cm)
}

// ClothingItem 衣物资产模型
type ClothingItem struct {
	gorm.Model
	UserID             uint               `json:"user_id" gorm:"not null;index"`
	CategoryID         uint               `json:"category_id" gorm:"not null;index"`
	Name               string             `json:"name" gorm:"not null"`
	Brand              string             `json:"brand"`
	Color              string             `json:"color" gorm:"not null"`
	Size               ClothingSize       `json:"size" gorm:"type:json"`
	Material           string             `json:"material"`
	Description        string             `json:"description"` // 详细描述
	Price              float64            `json:"price" gorm:"type:decimal(10,2)"`
	PurchaseDate       *time.Time         `json:"purchase_date"`
	Condition          api.ClothingStatus `json:"condition" gorm:"default:'active'"`
	WearCount          int                `json:"wear_count" gorm:"default:0"`
	DurabilityScore    float64            `json:"durability_score" gorm:"default:100.0"`
	LastWornDate       *time.Time         `json:"last_worn_date"`
	SpecificAttributes SpecificAttributes `json:"specific_attributes" gorm:"type:json"`
	Notes              string             `json:"notes"`
	IsActive           bool               `json:"is_active" gorm:"default:true"`
	IsFavorite         bool               `json:"is_favorite" gorm:"default:false"`
}

// TableName 指定表名
func (ClothingItem) TableName() string {
	return "clothing_items"
}

// CalculateDurability 计算耐久度
func (c *ClothingItem) CalculateDurability() float64 {
	baseScore := 100.0

	// 根据穿着次数计算磨损
	wearDamage := float64(c.WearCount) * c.getWearDamageRate()

	// 根据时间计算自然磨损
	timeDamage := c.getTimeDamage()

	// 根据保养记录计算加分
	maintenanceBonus := c.getMaintenanceBonus()

	// 根据材质计算耐久系数
	materialFactor := c.getMaterialDurabilityFactor()

	durability := (baseScore - wearDamage - timeDamage + maintenanceBonus) * materialFactor

	// 确保耐久度在0-100之间
	if durability < 0 {
		durability = 0
	} else if durability > 100 {
		durability = 100
	}

	return math.Round(durability*100) / 100
}

// getWearDamageRate 获取穿着磨损率
func (c *ClothingItem) getWearDamageRate() float64 {
	// 由于移除了Category关联对象，这里使用默认磨损率
	// 具体的分类相关逻辑可以在服务层通过CategoryID查询分类信息后处理
	return 1.0 // 默认磨损率
}

// getTimeDamage 获取时间磨损
func (c *ClothingItem) getTimeDamage() float64 {
	if c.PurchaseDate == nil {
		return 0
	}

	// 计算购买至今的天数
	daysSincePurchase := time.Since(*c.PurchaseDate).Hours() / 24

	// 每年自然磨损5分
	yearsSincePurchase := daysSincePurchase / 365
	return yearsSincePurchase * 5.0
}

// getMaintenanceBonus 获取保养加分
func (c *ClothingItem) getMaintenanceBonus() float64 {
	// 这里简化处理，实际应该从数据库查询保养记录
	// 每次保养可以恢复2-5分耐久度
	// 由于移除了直接关联，这里返回默认值，实际使用时需要传入保养记录数量
	return 0.0
}

// getMaterialDurabilityFactor 获取材质耐久系数
func (c *ClothingItem) getMaterialDurabilityFactor() float64 {
	materialFactors := map[string]float64{
		"真皮":   1.2,
		"羊毛":   1.1,
		"棉":    1.0,
		"聚酯纤维": 0.9,
		"尼龙":   1.1,
		"丝绸":   0.8,
		"麻":    1.0,
	}

	if factor, exists := materialFactors[c.Material]; exists {
		return factor
	}
	return 1.0 // 默认系数
}

// UpdateDurability 更新耐久度
func (c *ClothingItem) UpdateDurability(db *gorm.DB) error {
	c.DurabilityScore = c.CalculateDurability()
	return db.Model(c).Update("durability_score", c.DurabilityScore).Error
}

// IncrementWearCount 增加穿着次数
func (c *ClothingItem) IncrementWearCount(db *gorm.DB) error {
	c.WearCount++
	now := time.Now()
	c.LastWornDate = &now

	// 更新穿着次数和最后穿着时间
	err := db.Model(c).Updates(map[string]interface{}{
		"wear_count":     c.WearCount,
		"last_worn_date": c.LastWornDate,
	}).Error

	if err != nil {
		return err
	}

	// 重新计算耐久度
	return c.UpdateDurability(db)
}

// GetCostPerWear 获取每次穿着成本
func (c *ClothingItem) GetCostPerWear() float64 {
	if c.WearCount == 0 {
		return c.Price
	}
	return c.Price / float64(c.WearCount)
}

// IsValidCondition 检查状态是否有效
func IsValidCondition(condition string) bool {
	return api.ClothingStatus(condition).IsValid()
}
