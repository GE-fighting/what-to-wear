package repositories

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

// ClothingItemRepository 衣物仓库接口
type ClothingItemRepository interface {
	// 基础CRUD操作
	Create(item *models.ClothingItem) error
	GetByID(id uint) (*models.ClothingItem, error)
	GetByUserID(userID uint, req *dto.ClothingItemListRequest) ([]models.ClothingItem, int64, error)
	Update(item *models.ClothingItem) error
	Delete(id uint) error

	// 高级查询
	GetByCategory(userID, categoryID uint, limit int) ([]models.ClothingItem, error)
	GetByTags(userID uint, tagIDs []uint, limit int) ([]models.ClothingItem, error)
	GetFavorites(userID uint, limit int) ([]models.ClothingItem, error)
	GetRecentlyAdded(userID uint, limit int) ([]models.ClothingItem, error)
	GetMostWorn(userID uint, limit int) ([]models.ClothingItem, error)
	GetLeastWorn(userID uint, limit int) ([]models.ClothingItem, error)

	// 统计查询
	GetStats(userID uint) (*dto.ClothingStatsResponse, error)
	GetCategoryStats(userID uint) ([]dto.CategoryStatsItem, error)
	GetBrandStats(userID uint) ([]dto.BrandStatsItem, error)
	GetColorStats(userID uint) ([]dto.ColorStatsItem, error)

	// 搜索
	Search(userID uint, query string, limit int) ([]models.ClothingItem, error)

	// 标签关联
	AddTags(itemID uint, tagIDs []uint) error
	RemoveTags(itemID uint, tagIDs []uint) error
	GetItemTags(itemID uint) ([]models.ClothingTag, error)

	// 穿着记录
	IncrementWearCount(itemID uint) error
	UpdateDurability(itemID uint, score float64) error
}

// ClothingCategoryRepository 衣物分类仓库接口
type ClothingCategoryRepository interface {
	// 基础CRUD操作
	Create(category *models.ClothingCategory) error
	GetByID(id uint) (*models.ClothingCategory, error)
	GetAll() ([]models.ClothingCategory, error)
	Update(category *models.ClothingCategory) error
	Delete(id uint) error

	// 层级查询
	GetRootCategories() ([]models.ClothingCategory, error)
	GetChildCategories(parentID uint) ([]models.ClothingCategory, error)
	GetCategoryTree() ([]models.ClothingCategory, error)

	// 统计
	GetCategoryItemCount(categoryID uint) (int64, error)
}

// ClothingTagRepository 衣物标签仓库接口
type ClothingTagRepository interface {
	// 基础CRUD操作
	Create(tag *models.ClothingTag) error
	GetByID(id uint) (*models.ClothingTag, error)
	GetAll() ([]models.ClothingTag, error)
	GetByUserID(userID uint) ([]models.ClothingTag, error)
	Update(tag *models.ClothingTag) error
	Delete(id uint) error

	// 按类型查询
	GetByType(tagType models.TagType, userID *uint) ([]models.ClothingTag, error)
	GetSystemTags() ([]models.ClothingTag, error)
	GetUserTags(userID uint) ([]models.ClothingTag, error)

	// 统计
	GetTagItemCount(tagID uint) (int64, error)
	GetPopularTags(userID uint, limit int) ([]models.ClothingTag, error)
}

// PurchaseRecordRepository 购买记录仓库接口
type PurchaseRecordRepository interface {
	// 基础CRUD操作
	Create(record *models.PurchaseRecord) error
	GetByID(id uint) (*models.PurchaseRecord, error)
	GetByClothingItemID(clothingItemID uint) (*models.PurchaseRecord, error)
	Update(record *models.PurchaseRecord) error
	Delete(id uint) error

	// 查询
	GetByUserID(userID uint, limit int) ([]models.PurchaseRecord, error)
	GetByDateRange(userID uint, startDate, endDate string) ([]models.PurchaseRecord, error)
	GetByStore(userID uint, storeName string) ([]models.PurchaseRecord, error)

	// 统计
	GetTotalSpent(userID uint) (float64, error)
	GetSpentByMonth(userID uint, year int) (map[string]float64, error)
	GetSpentByCategory(userID uint) (map[string]float64, error)
}

// MaintenanceRecordRepository 保养记录仓库接口
type MaintenanceRecordRepository interface {
	// 基础CRUD操作
	Create(record *models.MaintenanceRecord) error
	GetByID(id uint) (*models.MaintenanceRecord, error)
	GetByClothingItemID(clothingItemID uint) ([]models.MaintenanceRecord, error)
	Update(record *models.MaintenanceRecord) error
	Delete(id uint) error

	// 查询
	GetByUserID(userID uint, limit int) ([]models.MaintenanceRecord, error)
	GetByType(userID uint, maintenanceType models.MaintenanceType) ([]models.MaintenanceRecord, error)
	GetUpcoming(userID uint, days int) ([]models.MaintenanceRecord, error)
	GetOverdue(userID uint) ([]models.MaintenanceRecord, error)

	// 统计
	GetMaintenanceCost(userID uint) (float64, error)
	GetMaintenanceFrequency(userID uint) (map[string]int64, error)
}

// WearRecordRepository 穿着记录仓库接口
type WearRecordRepository interface {
	// 基础CRUD操作
	Create(record *models.WearRecord) error
	GetByID(id uint) (*models.WearRecord, error)
	GetByClothingItemID(clothingItemID uint, limit int) ([]models.WearRecord, error)
	Update(record *models.WearRecord) error
	Delete(id uint) error

	// 查询
	GetByUserID(userID uint, limit int) ([]models.WearRecord, error)
	GetByDateRange(userID uint, startDate, endDate string) ([]models.WearRecord, error)
	GetByOccasion(userID uint, occasion string) ([]models.WearRecord, error)
	GetByWeather(userID uint, weather string) ([]models.WearRecord, error)

	// 统计
	GetWearStats(clothingItemID uint) (map[string]interface{}, error)
	GetWearFrequency(userID uint) (map[string]int64, error)
	GetComfortRatings(userID uint) (map[uint]float64, error)
}
