package repositories

import (
	"context"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// OutfitItemRepository 穿搭单品仓库接口
type OutfitItemRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, item *models.OutfitItem) error
	GetByID(ctx context.Context, id uint) (*models.OutfitItem, error)
	Update(ctx context.Context, item *models.OutfitItem) error
	Delete(ctx context.Context, id uint) error

	// 查询操作
	GetByOutfitID(ctx context.Context, outfitID uint) ([]models.OutfitItem, error)
	GetByClothingItemID(ctx context.Context, clothingItemID uint) ([]models.OutfitItem, error)
	GetByRole(ctx context.Context, outfitID uint, role string) ([]models.OutfitItem, error)

	// 批量操作
	CreateBatch(ctx context.Context, items []models.OutfitItem) error
	DeleteByOutfitID(ctx context.Context, outfitID uint) error
	UpdateLayerOrders(ctx context.Context, outfitID uint, items []models.OutfitItem) error

	// 统计操作
	GetItemUsageCount(ctx context.Context, clothingItemID uint) (int64, error)
	GetPopularItems(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error)
}

// outfitItemRepository 穿搭单品仓库实现
type outfitItemRepository struct {
	db *gorm.DB
}

// NewOutfitItemRepository 创建穿搭单品仓库实例
func NewOutfitItemRepository(db *gorm.DB) OutfitItemRepository {
	return &outfitItemRepository{db: db}
}

// Create 创建穿搭单品
func (r *outfitItemRepository) Create(ctx context.Context, item *models.OutfitItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// GetByID 根据ID获取穿搭单品
func (r *outfitItemRepository) GetByID(ctx context.Context, id uint) (*models.OutfitItem, error) {
	var item models.OutfitItem
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Update 更新穿搭单品
func (r *outfitItemRepository) Update(ctx context.Context, item *models.OutfitItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

// Delete 删除穿搭单品
func (r *outfitItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.OutfitItem{}, id).Error
}

// GetByOutfitID 根据穿搭ID获取所有单品
func (r *outfitItemRepository) GetByOutfitID(ctx context.Context, outfitID uint) ([]models.OutfitItem, error) {
	var items []models.OutfitItem
	err := r.db.WithContext(ctx).Where("outfit_id = ?", outfitID).
		Order("layer_order ASC, created_at ASC").
		Find(&items).Error
	return items, err
}

// GetByClothingItemID 根据衣物ID获取相关穿搭单品
func (r *outfitItemRepository) GetByClothingItemID(ctx context.Context, clothingItemID uint) ([]models.OutfitItem, error) {
	var items []models.OutfitItem
	err := r.db.WithContext(ctx).Where("clothing_item_id = ?", clothingItemID).
		Order("created_at DESC").
		Find(&items).Error
	return items, err
}

// GetByRole 根据角色获取穿搭单品
func (r *outfitItemRepository) GetByRole(ctx context.Context, outfitID uint, role string) ([]models.OutfitItem, error) {
	var items []models.OutfitItem
	err := r.db.WithContext(ctx).Where("outfit_id = ? AND item_role = ?", outfitID, role).
		Order("layer_order ASC").
		Find(&items).Error
	return items, err
}

// CreateBatch 批量创建穿搭单品
func (r *outfitItemRepository) CreateBatch(ctx context.Context, items []models.OutfitItem) error {
	if len(items) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&items).Error
}

// DeleteByOutfitID 删除指定穿搭的所有单品
func (r *outfitItemRepository) DeleteByOutfitID(ctx context.Context, outfitID uint) error {
	return r.db.WithContext(ctx).Where("outfit_id = ?", outfitID).Delete(&models.OutfitItem{}).Error
}

// UpdateLayerOrders 更新穿搭单品的层次顺序
func (r *outfitItemRepository) UpdateLayerOrders(ctx context.Context, outfitID uint, items []models.OutfitItem) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if item.OutfitID != outfitID {
				continue
			}
			err := tx.Model(&models.OutfitItem{}).
				Where("id = ? AND outfit_id = ?", item.ID, outfitID).
				Update("layer_order", item.LayerOrder).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetItemUsageCount 获取衣物在穿搭中的使用次数
func (r *outfitItemRepository) GetItemUsageCount(ctx context.Context, clothingItemID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.OutfitItem{}).
		Where("clothing_item_id = ?", clothingItemID).
		Count(&count).Error
	return count, err
}

// GetPopularItems 获取热门衣物（按穿搭使用频率排序）
func (r *outfitItemRepository) GetPopularItems(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem

	query := r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Select("clothing_items.*, COUNT(outfit_items.clothing_item_id) as usage_count").
		Joins("JOIN outfit_items ON clothing_items.id = outfit_items.clothing_item_id").
		Joins("JOIN outfits ON outfit_items.outfit_id = outfits.id").
		Where("clothing_items.user_id = ? AND clothing_items.is_active = ?", userID, true).
		Group("clothing_items.id").
		Order("usage_count DESC, clothing_items.created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&items).Error
	return items, err
}

// GetOutfitItemsWithDetails 获取穿搭单品及其详细信息
func (r *outfitItemRepository) GetOutfitItemsWithDetails(ctx context.Context, outfitID uint) ([]models.OutfitItem, error) {
	var items []models.OutfitItem
	err := r.db.WithContext(ctx).Where("outfit_id = ?", outfitID).
		Preload("ClothingItem").
		Preload("ClothingItem.Category").
		Order("layer_order ASC, created_at ASC").
		Find(&items).Error
	return items, err
}

// GetItemsByRole 获取指定角色的所有单品
func (r *outfitItemRepository) GetItemsByRole(ctx context.Context, userID uint, role string, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem

	query := r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Joins("JOIN outfit_items ON clothing_items.id = outfit_items.clothing_item_id").
		Joins("JOIN outfits ON outfit_items.outfit_id = outfits.id").
		Where("outfits.user_id = ? AND outfit_items.item_role = ? AND clothing_items.is_active = ?", userID, role, true).
		Group("clothing_items.id").
		Order("clothing_items.created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&items).Error
	return items, err
}

// GetRoleStats 获取角色统计信息
func (r *outfitItemRepository) GetRoleStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		ItemRole string `json:"item_role"`
		Count    int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.OutfitItem{}).
		Select("item_role, COUNT(*) as count").
		Joins("JOIN outfits ON outfit_items.outfit_id = outfits.id").
		Where("outfits.user_id = ? AND outfit_items.item_role != ''", userID).
		Group("item_role").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	roleStats := make(map[string]int64)
	for _, result := range results {
		roleStats[result.ItemRole] = result.Count
	}

	return roleStats, nil
}

// GetLayerStats 获取层次统计信息
func (r *outfitItemRepository) GetLayerStats(ctx context.Context, userID uint) (map[int]int64, error) {
	var results []struct {
		LayerOrder int   `json:"layer_order"`
		Count      int64 `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.OutfitItem{}).
		Select("layer_order, COUNT(*) as count").
		Joins("JOIN outfits ON outfit_items.outfit_id = outfits.id").
		Where("outfits.user_id = ?", userID).
		Group("layer_order").
		Order("layer_order ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	layerStats := make(map[int]int64)
	for _, result := range results {
		layerStats[result.LayerOrder] = result.Count
	}

	return layerStats, nil
}

// CheckItemInOutfit 检查衣物是否已在指定穿搭中
func (r *outfitItemRepository) CheckItemInOutfit(ctx context.Context, outfitID, clothingItemID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.OutfitItem{}).
		Where("outfit_id = ? AND clothing_item_id = ?", outfitID, clothingItemID).
		Count(&count).Error
	return count > 0, err
}

// GetMaxLayerOrder 获取穿搭中的最大层次顺序
func (r *outfitItemRepository) GetMaxLayerOrder(ctx context.Context, outfitID uint) (int, error) {
	var maxOrder int
	err := r.db.WithContext(ctx).Model(&models.OutfitItem{}).
		Where("outfit_id = ?", outfitID).
		Select("COALESCE(MAX(layer_order), 0)").
		Scan(&maxOrder).Error
	return maxOrder, err
}
