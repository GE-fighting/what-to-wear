package dto

import (
	"time"
	"what-to-wear/server/api"
)

// CreateClothingItemDTO 创建衣物DTO
type CreateClothingItemDTO struct {
	CategoryID         uint                     `json:"category_id" binding:"required"`
	CategoryName       string                   `json:"category_name" binding:"required"`
	Name               string                   `json:"name" binding:"required"`
	Brand              string                   `json:"brand"`
	Color              string                   `json:"color"`
	Size               string                   `json:"size"`
	Material           string                   `json:"material"`
	Style              string                   `json:"style"`
	Description        string                   `json:"description"`
	TagNames           []string                 `json:"tag_names"`
	Status             api.ClothingStatus       `json:"status"`
	IsFavorite         bool                     `json:"is_favorite"`
	SpecificAttributes map[string]interface{}   `json:"specific_attributes"`
	PurchaseInfo       *CreatePurchaseRecordDTO `json:"purchase_info,omitempty"`
	Tags               []uint                   `json:"tags"`
}

// UpdateClothingItemDTO 更新衣物DTO
type UpdateClothingItemDTO struct {
	CategoryID  *uint               `json:"category_id"`
	Name        *string             `json:"name"`
	Brand       *string             `json:"brand"`
	Color       *string             `json:"color"`
	Size        *string             `json:"size"`
	Material    *string             `json:"material"`
	Season      []string            `json:"season"`      // 通过标签系统管理
	Occasion    []string            `json:"occasion"`    // 通过标签系统管理
	Style       []string            `json:"style"`       // 改为标签管理，支持多风格
	Description *string             `json:"description"` // 直接字段
	Tags        []uint              `json:"tags"`
	Status      *api.ClothingStatus `json:"status"`
	IsFavorite  *bool               `json:"is_favorite"`
}

// ClothingItemDTO 衣物DTO
type ClothingItemDTO struct {
	ID                 uint                   `json:"id"`
	UserID             uint                   `json:"user_id"`
	CategoryID         uint                   `json:"category_id"`
	CategoryName       string                 `json:"category_name"`
	Name               string                 `json:"name"`
	Brand              string                 `json:"brand"`
	Color              string                 `json:"color"`
	Size               string                 `json:"size"`
	Material           string                 `json:"material"`
	Season             []string               `json:"season"`
	Occasion           []string               `json:"occasion"`
	Style              []string               `json:"style"` // 改为数组，支持多风格
	Description        string                 `json:"description"`
	Status             api.ClothingStatus     `json:"status"`
	Tags               []TagDTO               `json:"tags"`
	Attachments        []AttachmentDTO        `json:"attachments"`
	PurchaseRecord     *PurchaseRecordDTO     `json:"purchase_record,omitempty"`
	MaintenanceRecords []MaintenanceRecordDTO `json:"maintenance_records"`
	WearRecords        []WearRecordDTO        `json:"wear_records"`
	WearCount          int                    `json:"wear_count"`
	LastWornDate       *time.Time             `json:"last_worn_date"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

// ClothingItemListDTO 衣物列表DTO
type ClothingItemListDTO struct {
	CategoryIDs []uint              `form:"category_ids"`
	TagIDs      []uint              `form:"tag_ids"`
	Status      *api.ClothingStatus `form:"status"`
	Brand       string              `form:"brand"`
	Color       string              `form:"color"`
	Season      string              `form:"season"`
	Occasion    string              `form:"occasion"`
	Material    string              `form:"material"`
	Condition   string              `form:"condition"`
	MinPrice    *float32            `form:"min_price"`
	MaxPrice    *float32            `form:"max_price"`
	IsFavorite  *bool               `form:"is_favorite"`
	Search      string              `form:"search"`
	SearchRequest
}

// ClothingItemListResponse 衣物列表响应DTO
type ClothingItemListResponseDTO struct {
	Items      []ClothingItemDTO `json:"items"`
	TotalCount int64             `json:"total_count"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalPages int               `json:"total_pages"`
}

// CreatePurchaseRecordDTO 创建购买记录DTO - 简化版
type CreatePurchaseRecordDTO struct {
	Price        float64   `json:"price" binding:"required"`         // 实际购买价格
	Store        string    `json:"store"`                            // 商店名称（线上或线下）
	PurchaseDate time.Time `json:"purchase_date" binding:"required"` // 购买日期
	Notes        string    `json:"notes"`                            // 备注信息（可包含折扣、原价等信息）
}

// UpdatePurchaseRecordDTO 更新购买记录DTO - 简化版
type UpdatePurchaseRecordDTO struct {
	Price        *float64   `json:"price,omitempty"`         // 实际购买价格
	Store        *string    `json:"store,omitempty"`         // 商店名称
	PurchaseDate *time.Time `json:"purchase_date,omitempty"` // 购买日期
	Notes        *string    `json:"notes,omitempty"`         // 备注信息
}

// PurchaseRecordDTO 购买记录DTO - 简化版
type PurchaseRecordDTO struct {
	ID             uint      `json:"id"`
	ClothingItemID uint      `json:"clothing_item_id"`
	Price          float64   `json:"price"`         // 实际购买价格
	Store          string    `json:"store"`         // 商店名称
	PurchaseDate   time.Time `json:"purchase_date"` // 购买日期
	Notes          string    `json:"notes"`         // 备注信息
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CreateMaintenanceRecordDTO 创建保养记录DTO
type CreateMaintenanceRecordDTO struct {
	MaintenanceType     string     `json:"maintenance_type" binding:"required"`
	Cost                float64    `json:"cost"`
	MaintenanceDate     time.Time  `json:"maintenance_date" binding:"required"`
	ServiceProvider     string     `json:"service_provider"`
	Description         string     `json:"description"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date"`
	Notes               string     `json:"notes"`
}

// UpdateMaintenanceRecordDTO 更新保养记录DTO
type UpdateMaintenanceRecordDTO struct {
	MaintenanceType     *string    `json:"maintenance_type,omitempty"`
	Cost                *float64   `json:"cost,omitempty"`
	MaintenanceDate     *time.Time `json:"maintenance_date,omitempty"`
	ServiceProvider     *string    `json:"service_provider,omitempty"`
	Description         *string    `json:"description,omitempty"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date,omitempty"`
	Notes               *string    `json:"notes,omitempty"`
}

// MaintenanceRecordDTO 保养记录DTO
type MaintenanceRecordDTO struct {
	ID                  uint       `json:"id"`
	ClothingItemID      uint       `json:"clothing_item_id"`
	MaintenanceType     string     `json:"maintenance_type"`
	Cost                float64    `json:"cost"`
	MaintenanceDate     time.Time  `json:"maintenance_date"`
	ServiceProvider     string     `json:"service_provider"`
	Description         string     `json:"description"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date"`
	Notes               string     `json:"notes"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// CreateWearRecordDTO 创建穿着记录DTO - 简化版
type CreateWearRecordDTO struct {
	WearDate time.Time `json:"wear_date" binding:"required"`
	Notes    string    `json:"notes"`
}

// UpdateWearRecordDTO 更新穿着记录DTO - 简化版
type UpdateWearRecordDTO struct {
	WearDate *time.Time `json:"wear_date,omitempty"`
	Notes    *string    `json:"notes,omitempty"`
}

// WearRecordDTO 穿着记录DTO - 简化版
type WearRecordDTO struct {
	ID             uint      `json:"id"`
	ClothingItemID uint      `json:"clothing_item_id"`
	WearDate       time.Time `json:"wear_date"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CreateTagDTO 创建标签DTO
type CreateTagDTO struct {
	Name        string `json:"name" binding:"required,min=1,max=20"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"max=100"`
	Color       string `json:"color" binding:"max=7"` // HEX颜色代码，如 #FF0000
}

// UpdateTagDTO 更新标签DTO
type UpdateTagDTO struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=1,max=20"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=100"`
	Color       *string `json:"color,omitempty" binding:"omitempty,max=7"`
}

// TagDTO 标签DTO
type TagDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// TagStatsItem 标签统计项
type TagStatsItem struct {
	TagName    string  `json:"tag_name"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}

// ClothingStatsDTO 衣物统计DTO
type ClothingStatsDTO struct {
	TotalItems    int64                        `json:"total_items"`
	ByCategory    map[string]int64             `json:"by_category"`
	ByStatus      map[api.ClothingStatus]int64 `json:"by_status"`
	BySeason      map[string]int64             `json:"by_season"`
	ByOccasion    map[string]int64             `json:"by_occasion"`
	ByBrand       map[string]int64             `json:"by_brand"`
	ByColor       map[string]int64             `json:"by_color"`
	TotalValue    float64                      `json:"total_value"`
	AveragePrice  float64                      `json:"average_price"`
	MostWornItems []ClothingItemSummary        `json:"most_worn_items"`
	RecentlyAdded []ClothingItemSummary        `json:"recently_added"`
	LastUpdated   time.Time                    `json:"last_updated"`
}

// MaintenanceStatsDTO 保养统计DTO
type MaintenanceStatsDTO struct {
	TotalCost        float64            `json:"total_cost"`
	MaintenanceCount int64              `json:"maintenance_count"`
	CostByType       map[string]float64 `json:"cost_by_type"`
	LastMaintenance  *time.Time         `json:"last_maintenance"`
}

// WearStatsDTO 穿着统计DTO
type WearStatsDTO struct {
	TotalWears      int64                     `json:"total_wears"`
	AveragePerItem  float64                   `json:"average_per_item"`
	WearsByCategory map[string]int64          `json:"wears_by_category"`
	WearsByOccasion map[string]int64          `json:"wears_by_occasion"`
	WearsByWeather  map[api.WeatherType]int64 `json:"wears_by_weather"`
	MostWornItems   []ClothingItemSummary     `json:"most_worn_items"`
	RecentWears     []WearRecordDTO           `json:"recent_wears"`
	LastWearDate    *time.Time                `json:"last_wear_date"`
}
