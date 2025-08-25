package repositories

import (
	"context"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// ClothingCategoryRepository 衣物分类仓库接口
type ClothingCategoryRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, category *models.ClothingCategory) error
	GetByID(ctx context.Context, id uint) (*models.ClothingCategory, error)
	GetAll(ctx context.Context) ([]models.ClothingCategory, error)
	Update(ctx context.Context, category *models.ClothingCategory) error
	Delete(ctx context.Context, id uint) error

	// 层级查询
	GetRootCategories(ctx context.Context) ([]models.ClothingCategory, error)
	GetChildCategories(ctx context.Context, parentID uint) ([]models.ClothingCategory, error)
	GetCategoryTree(ctx context.Context) ([]models.ClothingCategory, error)

	// 统计
	GetCategoryItemCount(ctx context.Context, categoryID uint) (int64, error)
}

// clothingCategoryRepository 衣物分类仓库实现
type clothingCategoryRepository struct {
	db *gorm.DB
}

// NewClothingCategoryRepository 创建衣物分类仓库实例
func NewClothingCategoryRepository(db *gorm.DB) ClothingCategoryRepository {
	return &clothingCategoryRepository{db: db}
}

// Create 创建分类
func (r *clothingCategoryRepository) Create(ctx context.Context, category *models.ClothingCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// GetByID 根据ID获取分类
func (r *clothingCategoryRepository) GetByID(ctx context.Context, id uint) (*models.ClothingCategory, error) {
	var category models.ClothingCategory
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAll 获取所有分类
func (r *clothingCategoryRepository) GetAll(ctx context.Context) ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.WithContext(ctx).Where("is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// Update 更新分类
func (r *clothingCategoryRepository) Update(ctx context.Context, category *models.ClothingCategory) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete 删除分类
func (r *clothingCategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&models.ClothingCategory{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

// GetRootCategories 获取根分类
func (r *clothingCategoryRepository) GetRootCategories(ctx context.Context) ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.WithContext(ctx).Where("parent_id IS NULL AND is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetChildCategories 获取子分类
func (r *clothingCategoryRepository) GetChildCategories(ctx context.Context, parentID uint) ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.WithContext(ctx).Where("parent_id = ? AND is_active = ?", parentID, true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetCategoryTree 获取分类树
func (r *clothingCategoryRepository) GetCategoryTree(ctx context.Context) ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.WithContext(ctx).Where("is_active = ?", true).
		Order("COALESCE(parent_id, 0) ASC, sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetCategoryItemCount 获取分类下的衣物数量
func (r *clothingCategoryRepository) GetCategoryItemCount(ctx context.Context, categoryID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Where(&models.ClothingItem{}).
		Where("category_id = ? AND is_active = ?", categoryID, true).
		Count(&count).Error
	return count, err
}
