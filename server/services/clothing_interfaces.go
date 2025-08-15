package services

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

// ClothingItemService 衣物服务接口
type ClothingItemService interface {
	// 基础CRUD操作
	CreateClothingItem(userID uint, req *dto.CreateClothingItemRequest) (*dto.ClothingItemResponse, error)
	GetClothingItem(userID, itemID uint) (*dto.ClothingItemResponse, error)
	GetClothingItems(userID uint, req *dto.ClothingItemListRequest) (*dto.ClothingItemListResponse, error)
	UpdateClothingItem(userID, itemID uint, req *dto.UpdateClothingItemRequest) (*dto.ClothingItemResponse, error)
	DeleteClothingItem(userID, itemID uint) error

	// 批量操作
	BatchDeleteClothingItems(userID uint, itemIDs []uint) (map[string]interface{}, error)

	// 高级功能
	GetClothingStats(userID uint) (*dto.ClothingStatsResponse, error)
	SearchClothingItems(userID uint, query string, limit int) ([]dto.ClothingItemSummary, error)
	GetRecommendations(userID uint, occasion string, weather string) ([]dto.ClothingItemSummary, error)

	// 标签管理
	AddTagsToItem(userID, itemID uint, tagIDs []uint) error
	RemoveTagsFromItem(userID, itemID uint, tagIDs []uint) error

	// 穿着记录
	RecordWear(userID, itemID uint, req *dto.CreateWearRecordRequest) error
	GetWearHistory(userID, itemID uint, limit int) ([]models.WearRecord, error)
	GetWearRecords(userID, itemID uint, page, pageSize int) ([]dto.WearRecordResponse, error)

	// 保养记录
	AddMaintenanceRecord(userID, itemID uint, req *dto.CreateMaintenanceRecordRequest) error
	GetMaintenanceHistory(userID, itemID uint) ([]models.MaintenanceRecord, error)
	GetMaintenanceRecords(userID, itemID uint) ([]dto.MaintenanceRecordResponse, error)
	CreateMaintenanceRecord(userID, itemID uint, req *dto.CreateMaintenanceRecordRequest) (*dto.MaintenanceRecordResponse, error)
	GetMaintenanceReminders(userID uint, days int) ([]models.MaintenanceRecord, error)

	// 收藏管理
	ToggleFavorite(userID, itemID uint) error
	GetFavorites(userID uint, limit int) ([]dto.ClothingItemSummary, error)

	// 耐久度管理
	UpdateDurability(userID, itemID uint) error
	GetLowDurabilityItems(userID uint, threshold float64) ([]dto.ClothingItemSummary, error)
}

// ClothingCategoryService 衣物分类服务接口
type ClothingCategoryService interface {
	// 基础CRUD操作
	CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetCategory(categoryID uint) (*dto.CategoryResponse, error)
	GetAllCategories() ([]dto.CategoryResponse, error)
	UpdateCategory(categoryID uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
	DeleteCategory(categoryID uint) error

	// 层级管理
	GetCategoryTree() ([]dto.CategoryTreeNode, error)
	GetRootCategories() ([]dto.CategoryResponse, error)
	GetChildCategories(parentID uint) ([]dto.CategoryResponse, error)

	// 统计
	GetCategoryStats() ([]dto.CategoryStatsItem, error)
}

// ClothingTagService 衣物标签服务接口
type ClothingTagService interface {
	// 基础CRUD操作
	CreateTag(userID uint, req *dto.CreateTagRequest) (*dto.TagResponse, error)
	GetTag(tagID uint) (*dto.TagResponse, error)
	GetAllTags(userID uint) ([]dto.TagResponse, error)
	UpdateTag(userID, tagID uint, req *dto.UpdateTagRequest) (*dto.TagResponse, error)
	DeleteTag(userID, tagID uint) error

	// 按类型查询
	GetTagsByType(tagType models.TagType, userID *uint) ([]dto.TagResponse, error)
	GetSystemTags() ([]dto.TagResponse, error)
	GetUserTags(userID uint) ([]dto.TagResponse, error)

	// 统计
	GetPopularTags(userID uint, limit int) ([]dto.TagResponse, error)
	GetTagStats(userID uint) ([]dto.TagStatsItem, error)
}

// PurchaseRecordService 购买记录服务接口
type PurchaseRecordService interface {
	// 基础CRUD操作
	CreatePurchaseRecord(userID, itemID uint, req *dto.CreatePurchaseRecordRequest) (*dto.PurchaseRecordResponse, error)
	GetPurchaseRecord(userID, recordID uint) (*dto.PurchaseRecordResponse, error)
	GetPurchaseRecords(userID uint, limit int) ([]dto.PurchaseRecordResponse, error)
	UpdatePurchaseRecord(userID, recordID uint, req *dto.UpdatePurchaseRecordRequest) (*dto.PurchaseRecordResponse, error)
	DeletePurchaseRecord(userID, recordID uint) error

	// 查询
	GetPurchasesByDateRange(userID uint, startDate, endDate string) ([]dto.PurchaseRecordResponse, error)
	GetPurchasesByStore(userID uint, storeName string) ([]dto.PurchaseRecordResponse, error)

	// 统计
	GetSpendingStats(userID uint) (*dto.SpendingStatsResponse, error)
	GetSpendingByMonth(userID uint, year int) (map[string]float64, error)
	GetSpendingByCategory(userID uint) (map[string]float64, error)
}

// MaintenanceService 保养服务接口
type MaintenanceService interface {
	// 基础CRUD操作
	CreateMaintenanceRecord(userID, itemID uint, req *dto.CreateMaintenanceRecordRequest) (*dto.MaintenanceRecordResponse, error)
	GetMaintenanceRecord(userID, recordID uint) (*dto.MaintenanceRecordResponse, error)
	GetMaintenanceRecords(userID uint, limit int) ([]dto.MaintenanceRecordResponse, error)
	UpdateMaintenanceRecord(userID, recordID uint, req *dto.UpdateMaintenanceRecordRequest) (*dto.MaintenanceRecordResponse, error)
	DeleteMaintenanceRecord(userID, recordID uint) error

	// 提醒功能
	GetUpcomingMaintenance(userID uint, days int) ([]dto.MaintenanceReminderResponse, error)
	GetOverdueMaintenance(userID uint) ([]dto.MaintenanceReminderResponse, error)
	MarkReminderSent(recordID uint) error

	// 统计
	GetMaintenanceStats(userID uint) (*dto.MaintenanceStatsResponse, error)
	GetMaintenanceCostByType(userID uint) (map[string]float64, error)
}

// WearRecordService 穿着记录服务接口
type WearRecordService interface {
	// 基础CRUD操作
	CreateWearRecord(userID, itemID uint, req *dto.CreateWearRecordRequest) (*dto.WearRecordResponse, error)
	GetWearRecord(userID, recordID uint) (*dto.WearRecordResponse, error)
	GetWearRecords(userID uint, limit int) ([]dto.WearRecordResponse, error)
	UpdateWearRecord(userID, recordID uint, req *dto.UpdateWearRecordRequest) (*dto.WearRecordResponse, error)
	DeleteWearRecord(userID, recordID uint) error

	// 查询
	GetWearRecordsByItem(userID, itemID uint, limit int) ([]dto.WearRecordResponse, error)
	GetWearRecordsByDateRange(userID uint, startDate, endDate string) ([]dto.WearRecordResponse, error)
	GetWearRecordsByOccasion(userID uint, occasion string) ([]dto.WearRecordResponse, error)

	// 统计
	GetWearStats(userID uint) (*dto.WearStatsResponse, error)
	GetWearFrequency(userID uint) (map[string]int64, error)
	GetComfortAnalysis(userID uint) (*dto.ComfortAnalysisResponse, error)
}
