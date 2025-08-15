package repositories

import (
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// clothingCategoryRepository 衣物分类仓库实现
type clothingCategoryRepository struct {
	db *gorm.DB
}

// NewClothingCategoryRepository 创建衣物分类仓库实例
func NewClothingCategoryRepository(db *gorm.DB) ClothingCategoryRepository {
	return &clothingCategoryRepository{db: db}
}

// Create 创建分类
func (r *clothingCategoryRepository) Create(category *models.ClothingCategory) error {
	return r.db.Create(category).Error
}

// GetByID 根据ID获取分类
func (r *clothingCategoryRepository) GetByID(id uint) (*models.ClothingCategory, error) {
	var category models.ClothingCategory
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAll 获取所有分类
func (r *clothingCategoryRepository) GetAll() ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.Where("is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// Update 更新分类
func (r *clothingCategoryRepository) Update(category *models.ClothingCategory) error {
	return r.db.Save(category).Error
}

// Delete 删除分类
func (r *clothingCategoryRepository) Delete(id uint) error {
	return r.db.Model(&models.ClothingCategory{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

// GetRootCategories 获取根分类
func (r *clothingCategoryRepository) GetRootCategories() ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.Where("parent_id IS NULL AND is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetChildCategories 获取子分类
func (r *clothingCategoryRepository) GetChildCategories(parentID uint) ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.Where("parent_id = ? AND is_active = ?", parentID, true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetCategoryTree 获取分类树
func (r *clothingCategoryRepository) GetCategoryTree() ([]models.ClothingCategory, error) {
	var categories []models.ClothingCategory
	err := r.db.Where("is_active = ?", true).
		Order("COALESCE(parent_id, 0) ASC, sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}

// GetCategoryItemCount 获取分类下的衣物数量
func (r *clothingCategoryRepository) GetCategoryItemCount(categoryID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.ClothingItem{}).
		Where("category_id = ? AND is_active = ?", categoryID, true).
		Count(&count).Error
	return count, err
}