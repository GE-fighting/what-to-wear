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

// ItemRole 单品角色枚举
type ItemRole string

const (
	ItemRoleMain      ItemRole = "main"      // 主要单品
	ItemRoleAccent    ItemRole = "accent"    // 点缀单品
	ItemRoleBase      ItemRole = "base"      // 基础单品
	ItemRoleOuter     ItemRole = "outer"     // 外层单品
	ItemRoleInner     ItemRole = "inner"     // 内层单品
	ItemRoleBottom    ItemRole = "bottom"    // 下装
	ItemRoleShoes     ItemRole = "shoes"     // 鞋子
	ItemRoleAccessory ItemRole = "accessory" // 配饰
)

// IsValidItemRole 检查单品角色是否有效
func IsValidItemRole(role string) bool {
	validRoles := []ItemRole{
		ItemRoleMain, ItemRoleAccent, ItemRoleBase, ItemRoleOuter,
		ItemRoleInner, ItemRoleBottom, ItemRoleShoes, ItemRoleAccessory,
	}

	for _, validRole := range validRoles {
		if ItemRole(role) == validRole {
			return true
		}
	}
	return false
}

// GetOutfitItemsByOutfit 获取指定穿搭的所有单品
func GetOutfitItemsByOutfit(db *gorm.DB, outfitID uint) ([]OutfitItem, error) {
	var items []OutfitItem
	err := db.Where("outfit_id = ?", outfitID).
		Preload("ClothingItem").
		Preload("ClothingItem.Category").
		Preload("ClothingItem.Tags").
		Order("layer_order ASC").
		Find(&items).Error
	return items, err
}

// GetOutfitsByClothingItem 获取包含指定衣物的所有穿搭
func GetOutfitsByClothingItem(db *gorm.DB, clothingItemID uint) ([]Outfit, error) {
	var outfits []Outfit
	err := db.Joins("JOIN outfit_items ON outfits.id = outfit_items.outfit_id").
		Where("outfit_items.clothing_item_id = ?", clothingItemID).
		Preload("User").
		Order("outfits.date DESC").
		Find(&outfits).Error
	return outfits, err
}
