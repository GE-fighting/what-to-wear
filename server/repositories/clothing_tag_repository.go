package repositories

import (
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// clothingTagRepository 衣物标签仓库实现
type clothingTagRepository struct {
	db *gorm.DB
}

// NewClothingTagRepository 创建衣物标签仓库实例
func NewClothingTagRepository(db *gorm.DB) ClothingTagRepository {
	return &clothingTagRepository{db: db}
}

// Create 创建标签
func (r *clothingTagRepository) Create(tag *models.ClothingTag) error {
	return r.db.Create(tag).Error
}

// GetByID 根据ID获取标签
func (r *clothingTagRepository) GetByID(id uint) (*models.ClothingTag, error) {
	var tag models.ClothingTag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetAll 获取所有标签
func (r *clothingTagRepository) GetAll() ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	err := r.db.Where("is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&tags).Error
	return tags, err
}

// GetByUserID 根据用户ID获取标签（包括系统标签和用户自定义标签）
func (r *clothingTagRepository) GetByUserID(userID uint) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	err := r.db.Where("is_active = ? AND (is_system = ? OR user_id = ?)", true, true, userID).
		Order("is_system DESC, sort_order ASC, name ASC").
		Find(&tags).Error
	return tags, err
}

// Update 更新标签
func (r *clothingTagRepository) Update(tag *models.ClothingTag) error {
	return r.db.Save(tag).Error
}

// Delete 删除标签
func (r *clothingTagRepository) Delete(id uint) error {
	return r.db.Model(&models.ClothingTag{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

// GetByType 根据类型获取标签
func (r *clothingTagRepository) GetByType(tagType models.TagType, userID *uint) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	query := r.db.Where("type = ? AND is_active = ?", tagType, true)

	// 包含系统标签和用户自定义标签
	if userID != nil {
		query = query.Where("is_system = ? OR user_id = ?", true, *userID)
	} else {
		query = query.Where("is_system = ?", true)
	}

	err := query.Order("is_system DESC, sort_order ASC, name ASC").Find(&tags).Error
	return tags, err
}

// GetSystemTags 获取系统标签
func (r *clothingTagRepository) GetSystemTags() ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	err := r.db.Where("is_system = ? AND is_active = ?", true, true).
		Order("type ASC, sort_order ASC, name ASC").
		Find(&tags).Error
	return tags, err
}

// GetUserTags 获取用户自定义标签
func (r *clothingTagRepository) GetUserTags(userID uint) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).
		Order("type ASC, sort_order ASC, name ASC").
		Find(&tags).Error
	return tags, err
}

// GetTagItemCount 获取标签关联的衣物数量
func (r *clothingTagRepository) GetTagItemCount(tagID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.ClothingItemTag{}).
		Joins("JOIN clothing_items ON clothing_item_tags.clothing_item_id = clothing_items.id").
		Where("clothing_item_tags.clothing_tag_id = ? AND clothing_items.is_active = ?", tagID, true).
		Count(&count).Error
	return count, err
}

// GetPopularTags 获取热门标签（按使用频率排序）
func (r *clothingTagRepository) GetPopularTags(userID uint, limit int) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag

	query := r.db.Model(&models.ClothingTag{}).
		Select("clothing_tags.*, COUNT(clothing_item_tags.clothing_tag_id) as usage_count").
		Joins("LEFT JOIN clothing_item_tags ON clothing_tags.id = clothing_item_tags.clothing_tag_id").
		Joins("LEFT JOIN clothing_items ON clothing_item_tags.clothing_item_id = clothing_items.id").
		Where("clothing_tags.is_active = ? AND (clothing_tags.is_system = ? OR clothing_tags.user_id = ?)", true, true, userID).
		Where("clothing_items.user_id = ? OR clothing_items.user_id IS NULL", userID).
		Group("clothing_tags.id").
		Order("usage_count DESC, clothing_tags.name ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&tags).Error
	return tags, err
}

// GetTagsByNames 根据标签名称批量获取标签
func (r *clothingTagRepository) GetTagsByNames(names []string, userID *uint) ([]models.ClothingTag, error) {
	var tags []models.ClothingTag
	query := r.db.Where("name IN ? AND is_active = ?", names, true)

	if userID != nil {
		query = query.Where("is_system = ? OR user_id = ?", true, *userID)
	} else {
		query = query.Where("is_system = ?", true)
	}

	err := query.Find(&tags).Error
	return tags, err
}

// CreateSystemTagIfNotExists 创建系统标签（如果不存在）
func (r *clothingTagRepository) CreateSystemTagIfNotExists(tag *models.ClothingTag) error {
	var existingTag models.ClothingTag
	err := r.db.Where("name = ? AND type = ? AND is_system = ?", tag.Name, tag.Type, true).
		First(&existingTag).Error

	if err == gorm.ErrRecordNotFound {
		// 标签不存在，创建新标签
		tag.IsSystem = true
		return r.db.Create(tag).Error
	}

	return err
}

// GetTagUsageStats 获取标签使用统计
func (r *clothingTagRepository) GetTagUsageStats(userID uint) (map[uint]int64, error) {
	var results []struct {
		TagID uint  `json:"tag_id"`
		Count int64 `json:"count"`
	}

	err := r.db.Model(&models.ClothingItemTag{}).
		Select("clothing_tag_id as tag_id, COUNT(*) as count").
		Joins("JOIN clothing_items ON clothing_item_tags.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND clothing_items.is_active = ?", userID, true).
		Group("clothing_tag_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	statsMap := make(map[uint]int64)
	for _, result := range results {
		statsMap[result.TagID] = result.Count
	}

	return statsMap, nil
}