package repositories

import (
	"context"
	"fmt"
	"strings"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// ClothingItemRepository 衣物仓库接口
type ClothingItemRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, item *models.ClothingItem) error
	GetByID(ctx context.Context, id uint) (*models.ClothingItem, error)
	GetByUserID(ctx context.Context, userID uint, req *dto.ClothingItemListDTO) ([]models.ClothingItem, int64, error)
	Update(ctx context.Context, item *models.ClothingItem) error
	Delete(ctx context.Context, id uint) error

	// 高级查询
	GetByCategory(ctx context.Context, userID, categoryID uint, limit int) ([]models.ClothingItem, error)
	GetByTags(ctx context.Context, userID uint, tagIDs []uint, limit int) ([]models.ClothingItem, error)
	GetFavorites(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error)
	GetRecentlyAdded(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error)
	GetMostWorn(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error)
	GetLeastWorn(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error)

	// 统计查询
	GetCategoryStats(ctx context.Context, userID uint) ([]dto.CategoryStatsItem, error)
	GetBrandStats(ctx context.Context, userID uint) ([]dto.BrandStatsItem, error)
	GetColorStats(ctx context.Context, userID uint) ([]dto.ColorStatsItem, error)

	// 搜索
	Search(ctx context.Context, userID uint, query string, limit int) ([]models.ClothingItem, error)

	// 标签关联
	AddTags(ctx context.Context, itemID uint, tagIDs []uint) error
	RemoveTags(ctx context.Context, itemID uint, tagIDs []uint) error
	GetItemTags(ctx context.Context, itemID uint) ([]models.ClothingTag, error)

	// 穿着记录
	IncrementWearCount(ctx context.Context, itemID uint) error
	UpdateDurability(ctx context.Context, itemID uint, score float64) error
}

// clothingItemRepository 衣物仓库实现
type clothingItemRepository struct {
	db *gorm.DB
}

// NewClothingItemRepository 创建衣物仓库实例
func NewClothingItemRepository(db *gorm.DB) ClothingItemRepository {
	return &clothingItemRepository{db: db}
}

// Create 创建衣物
func (r *clothingItemRepository) Create(ctx context.Context, item *models.ClothingItem) error {
	return r.db.Create(item).Error
}

// GetByID 根据ID获取衣物
func (r *clothingItemRepository) GetByID(ctx context.Context, id uint) (*models.ClothingItem, error) {
	var item models.ClothingItem
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetByUserID 根据用户ID获取衣物列表
func (r *clothingItemRepository) GetByUserID(ctx context.Context, userID uint, req *dto.ClothingItemListDTO) ([]models.ClothingItem, int64, error) {
	var items []models.ClothingItem
	var total int64
	
	query := r.db.WithContext(ctx).Model(&models.ClothingItem{}).Where("user_id = ? AND is_active = ?", userID, true)

	// 应用过滤条件
	if req.CategoryIDs != nil && len(req.CategoryIDs) > 0 {
		query = query.Where("category_id IN ?", req.CategoryIDs)
	}
	if req.Color != "" {
		query = query.Where("color LIKE ?", "%"+req.Color+"%")
	}
	if req.Brand != "" {
		query = query.Where("brand LIKE ?", "%"+req.Brand+"%")
	}
	if req.Material != "" {
		query = query.Where("material LIKE ?", "%"+req.Material+"%")
	}
	if req.Condition != "" {
		query = query.Where("condition = ?", req.Condition)
	}
	if req.MinPrice != nil {
		query = query.Where("price >= ?", req.MinPrice)
	}
	if req.MaxPrice != nil {
		query = query.Where("price <= ?", req.MaxPrice)
	}
	if req.IsFavorite != nil {
		query = query.Where("is_favorite = ?", *req.IsFavorite)
	}
	if req.Search != "" {
		searchTerm := "%" + req.Search + "%"
		query = query.Where("name LIKE ? OR brand LIKE ? OR notes LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	// 标签过滤
	if len(req.TagIDs) > 0 {		
		query = query.Joins("JOIN clothing_item_tags ON clothing_items.id = clothing_item_tags.clothing_item_id").
			Where("clothing_item_tags.clothing_tag_id IN ?", req.TagIDs)
	}

	// 获取总数
	err := query.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 应用排序
	orderBy := "created_at DESC"
	if req.SortBy != "" {
		direction := "ASC"
		if req.SortOrder == "desc" {
			direction = "DESC"
		}
		orderBy = fmt.Sprintf("%s %s", req.SortBy, direction)
	}
	query = query.Order(orderBy)

	// 应用分页
	offset := (req.Page - 1) * req.PageSize
	err = query.Offset(offset).Limit(req.PageSize).Find(&items).Error

	return items, total, err
}

// Update 更新衣物
func (r *clothingItemRepository) Update(ctx context.Context, item *models.ClothingItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

// Delete 删除衣物
func (r *clothingItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

// GetByCategory 根据分类获取衣物
func (r *clothingItemRepository) GetByCategory(ctx context.Context, userID, categoryID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND category_id = ? AND is_active = ?", userID, categoryID, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetByTags 根据标签获取衣物
func (r *clothingItemRepository) GetByTags(ctx context.Context, userID uint, tagIDs []uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Joins("JOIN clothing_item_tags ON clothing_items.id = clothing_item_tags.clothing_item_id").
		Where("clothing_item_tags.clothing_tag_id IN ?", tagIDs).
		Order("clothing_items.created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetFavorites 获取收藏的衣物
func (r *clothingItemRepository) GetFavorites(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_favorite = ? AND is_active = ?", userID, true, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetRecentlyAdded 获取最近添加的衣物
func (r *clothingItemRepository) GetRecentlyAdded(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetMostWorn 获取最常穿的衣物
func (r *clothingItemRepository) GetMostWorn(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Order("wear_count DESC, created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetLeastWorn 获取最少穿的衣物
func (r *clothingItemRepository) GetLeastWorn(ctx context.Context, userID uint, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Order("wear_count ASC, created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetCategoryStats 获取分类统计
func (r *clothingItemRepository) GetCategoryStats(ctx context.Context, userID uint) ([]dto.CategoryStatsItem, error) {
	var stats []dto.CategoryStatsItem

	err := r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Select("clothing_categories.name as category_name, COUNT(*) as count, COALESCE(SUM(clothing_items.price), 0) as total_value, COALESCE(AVG(clothing_items.wear_count), 0) as avg_wear_count").
		Joins("JOIN clothing_categories ON clothing_items.category_id = clothing_categories.id").
		Where("clothing_items.user_id = ? AND clothing_items.is_active = ?", userID, true).
		Group("clothing_categories.id, clothing_categories.name").
		Scan(&stats).Error

	return stats, err
}

// GetBrandStats 获取品牌统计
func (r *clothingItemRepository) GetBrandStats(ctx context.Context, userID uint) ([]dto.BrandStatsItem, error) {
	var stats []dto.BrandStatsItem

	err := r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Select("brand as brand_name, COUNT(*) as count, COALESCE(SUM(price), 0) as total_value, COALESCE(AVG(wear_count), 0) as avg_wear_count").
		Where("user_id = ? AND is_active = ? AND brand != ''", userID, true).
		Group("brand").
		Order("count DESC").
		Scan(&stats).Error

	return stats, err
}

// GetColorStats 获取颜色统计
func (r *clothingItemRepository) GetColorStats(ctx context.Context, userID uint) ([]dto.ColorStatsItem, error) {
	var stats []dto.ColorStatsItem

	err := r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Select("color, COUNT(*) as count, COALESCE(SUM(price), 0) as total_value, COALESCE(AVG(wear_count), 0) as avg_wear_count").
		Where("user_id = ? AND is_active = ?", userID, true).
		Group("color").
		Order("count DESC").
		Scan(&stats).Error

	return stats, err
}

// Search 搜索衣物
func (r *clothingItemRepository) Search(ctx context.Context, userID uint, query string, limit int) ([]models.ClothingItem, error) {
	var items []models.ClothingItem
	searchTerm := "%" + strings.ToLower(query) + "%"

	dbQuery := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Where("LOWER(name) LIKE ? OR LOWER(brand) LIKE ? OR LOWER(color) LIKE ? OR LOWER(material) LIKE ? OR LOWER(notes) LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm, searchTerm).
		Order("created_at DESC")

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}

	err := dbQuery.WithContext(ctx).Find(&items).Error
	return items, err
}

// AddTags 为衣物添加标签
func (r *clothingItemRepository) AddTags(ctx context.Context, itemID uint, tagIDs []uint) error {
	for _, tagID := range tagIDs {
		// 检查关联是否已存在
		var count int64
		r.db.WithContext(ctx).Model(&models.ClothingItemTag{}).
			Where("clothing_item_id = ? AND clothing_tag_id = ?", itemID, tagID).
			Count(&count)

		if count == 0 {
			// 创建新的关联
			itemTag := models.ClothingItemTag{
				ClothingItemID: itemID,
				ClothingTagID:  tagID,
			}
			if err := r.db.WithContext(ctx).Create(&itemTag).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveTags 移除衣物标签
func (r *clothingItemRepository) RemoveTags(ctx context.Context, itemID uint, tagIDs []uint) error {
	return r.db.WithContext(ctx).Where("clothing_item_id = ? AND clothing_tag_id IN ?", itemID, tagIDs).
		Delete(&models.ClothingItemTag{}).Error
}

// GetItemTags 获取衣物的标签
func (r *clothingItemRepository) GetItemTags(ctx context.Context, itemID uint) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	err := r.db.WithContext(ctx).Model(&models.ClothingTag{}).
		Joins("JOIN clothing_item_tags ON clothing_tags.id = clothing_item_tags.clothing_tag_id").
		Where("clothing_item_tags.clothing_item_id = ?", itemID).
		Find(&tags).Error
	return tags, err
}

// IncrementWearCount 增加穿着次数
func (r *clothingItemRepository) IncrementWearCount(ctx context.Context, itemID uint) error {
	return r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Where("id = ?", itemID).
		UpdateColumn("wear_count", gorm.Expr("wear_count + ?", 1)).Error
}

// UpdateDurability 更新耐久度
func (r *clothingItemRepository) UpdateDurability(ctx context.Context, itemID uint, score float64) error {
	return r.db.WithContext(ctx).Model(&models.ClothingItem{}).
		Where("id = ?", itemID).
		Update("durability_score", score).Error
}
