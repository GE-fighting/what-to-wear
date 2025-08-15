package dto

import (
	"time"
	"what-to-wear/server/models"
)

// CreateClothingItemRequest 创建衣物请求
type CreateClothingItemRequest struct {
	CategoryID         uint                         `json:"category_id" binding:"required"`
	Name               string                       `json:"name" binding:"required"`
	Brand              string                       `json:"brand"`
	Model              string                       `json:"model"`
	Color              string                       `json:"color" binding:"required"`
	Size               models.ClothingSize          `json:"size"`
	Material           string                       `json:"material"`
	Price              float64                      `json:"price"`
	PurchaseDate       *time.Time                   `json:"purchase_date"`
	Condition          string                       `json:"condition"`
	SpecificAttributes models.SpecificAttributes    `json:"specific_attributes"`
	ImageURLs          []string                     `json:"image_urls"`
	Notes              string                       `json:"notes"`
	TagIDs             []uint                       `json:"tag_ids"`
	PurchaseInfo       *CreatePurchaseRecordRequest `json:"purchase_info,omitempty"`
}

// UpdateClothingItemRequest 更新衣物请求
type UpdateClothingItemRequest struct {
	CategoryID         *uint                      `json:"category_id"`
	Name               *string                    `json:"name"`
	Brand              *string                    `json:"brand"`
	Model              *string                    `json:"model"`
	Color              *string                    `json:"color"`
	Size               *models.ClothingSize       `json:"size"`
	Material           *string                    `json:"material"`
	Price              *float64                   `json:"price"`
	PurchaseDate       *time.Time                 `json:"purchase_date"`
	Condition          *string                    `json:"condition"`
	SpecificAttributes *models.SpecificAttributes `json:"specific_attributes"`
	ImageURLs          []string                   `json:"image_urls"`
	Notes              *string                    `json:"notes"`
	IsFavorite         *bool                      `json:"is_favorite"`
	TagIDs             []uint                     `json:"tag_ids"`
}

// ClothingItemResponse 衣物响应
type ClothingItemResponse struct {
	ID                 uint                      `json:"id"`
	UserID             uint                      `json:"user_id"`
	CategoryID         uint                      `json:"category_id"`
	Category           CategorySummary           `json:"category"`
	Name               string                    `json:"name"`
	Brand              string                    `json:"brand"`
	Model              string                    `json:"model"`
	Color              string                    `json:"color"`
	Size               models.ClothingSize       `json:"size"`
	Material           string                    `json:"material"`
	Price              float64                   `json:"price"`
	PurchaseDate       *time.Time                `json:"purchase_date"`
	Condition          string                    `json:"condition"`
	WearCount          int                       `json:"wear_count"`
	DurabilityScore    float64                   `json:"durability_score"`
	LastWornDate       *time.Time                `json:"last_worn_date"`
	SpecificAttributes models.SpecificAttributes `json:"specific_attributes"`
	ImageURLs          []string                  `json:"image_urls"`
	Notes              string                    `json:"notes"`
	IsActive           bool                      `json:"is_active"`
	IsFavorite         bool                      `json:"is_favorite"`
	Tags               []TagSummary              `json:"tags"`
	CostPerWear        float64                   `json:"cost_per_wear"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

// ClothingItemSummary 衣物摘要
type ClothingItemSummary struct {
	ID              uint         `json:"id"`
	Name            string       `json:"name"`
	Brand           string       `json:"brand"`
	Color           string       `json:"color"`
	CategoryName    string       `json:"category_name"`
	ImageURL        string       `json:"image_url"` // 主图片
	DurabilityScore float64      `json:"durability_score"`
	WearCount       int          `json:"wear_count"`
	IsFavorite      bool         `json:"is_favorite"`
	Tags            []TagSummary `json:"tags"`
}

// ClothingItemListRequest 衣物列表请求
type ClothingItemListRequest struct {
	Page       int      `form:"page" binding:"min=1"`
	PageSize   int      `form:"page_size" binding:"min=1,max=100"`
	CategoryID *uint    `form:"category_id"`
	TagIDs     []uint   `form:"tag_ids"`
	Color      string   `form:"color"`
	Brand      string   `form:"brand"`
	Material   string   `form:"material"`
	Condition  string   `form:"condition"`
	MinPrice   *float64 `form:"min_price"`
	MaxPrice   *float64 `form:"max_price"`
	IsFavorite *bool    `form:"is_favorite"`
	Search     string   `form:"search"`
	SortBy     string   `form:"sort_by"`    // name, price, wear_count, durability_score, created_at
	SortOrder  string   `form:"sort_order"` // asc, desc
}

// ClothingItemListResponse 衣物列表响应
type ClothingItemListResponse struct {
	Items      []ClothingItemSummary `json:"items"`
	Total      int64                 `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	TotalPages int                   `json:"total_pages"`
}

// CreatePurchaseRecordRequest 创建购买记录请求
type CreatePurchaseRecordRequest struct {
	PurchasePrice  float64   `json:"purchase_price" binding:"required"`
	OriginalPrice  float64   `json:"original_price"`
	StoreName      string    `json:"store_name"`
	StoreLocation  string    `json:"store_location"`
	OnlineStore    string    `json:"online_store"`
	OrderNumber    string    `json:"order_number"`
	PurchaseDate   time.Time `json:"purchase_date" binding:"required"`
	PaymentMethod  string    `json:"payment_method"`
	Currency       string    `json:"currency"`
	ReceiptURL     string    `json:"receipt_url"`
	WarrantyPeriod int       `json:"warranty_period"`
	ReturnPolicy   string    `json:"return_policy"`
	Notes          string    `json:"notes"`
	Tags           []string  `json:"tags"`
}

// CreateMaintenanceRecordRequest 创建保养记录请求
type CreateMaintenanceRecordRequest struct {
	MaintenanceType    string    `json:"maintenance_type" binding:"required"`
	Cost               float64   `json:"cost"`
	MaintenanceDate    time.Time `json:"maintenance_date" binding:"required"`
	ServiceProvider    string    `json:"service_provider"`
	ServiceLocation    string    `json:"service_location"`
	BeforeCondition    string    `json:"before_condition"`
	AfterCondition     string    `json:"after_condition"`
	EffectivenessScore int       `json:"effectiveness_score" binding:"min=1,max=10"`
	Notes              string    `json:"notes"`
	Images             []string  `json:"images"`
}

// CreateWearRecordRequest 创建穿着记录请求
type CreateWearRecordRequest struct {
	WearDate         time.Time `json:"wear_date" binding:"required"`
	DurationHours    int       `json:"duration_hours" binding:"min=1,max=24"`
	Occasion         string    `json:"occasion"`
	WeatherCondition string    `json:"weather_condition"`
	Temperature      float64   `json:"temperature"`
	Activity         string    `json:"activity"`
	ComfortRating    int       `json:"comfort_rating" binding:"min=1,max=10"`
	StyleRating      int       `json:"style_rating" binding:"min=1,max=10"`
	Notes            string    `json:"notes"`
	Photos           []string  `json:"photos"`
	Location         string    `json:"location"`
	Companions       []string  `json:"companions"`
	Mood             string    `json:"mood"`
	WearIntensity    string    `json:"wear_intensity"`
}

// CategorySummary 分类摘要
type CategorySummary struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	ParentID    *uint  `json:"parent_id"`
	ParentName  string `json:"parent_name,omitempty"`
}

// TagSummary 标签摘要
type TagSummary struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
}

// ClothingStatsResponse 衣物统计响应
type ClothingStatsResponse struct {
	TotalItems        int64                 `json:"total_items"`
	TotalValue        float64               `json:"total_value"`
	AverageDurability float64               `json:"average_durability"`
	MostWornItem      *ClothingItemSummary  `json:"most_worn_item"`
	RecentlyAdded     []ClothingItemSummary `json:"recently_added"`
	CategoryStats     []CategoryStatsItem   `json:"category_stats"`
	BrandStats        []BrandStatsItem      `json:"brand_stats"`
	ColorStats        []ColorStatsItem      `json:"color_stats"`
}

// CategoryStatsItem 分类统计项
type CategoryStatsItem struct {
	CategoryName string  `json:"category_name"`
	Count        int64   `json:"count"`
	TotalValue   float64 `json:"total_value"`
	AvgWearCount float64 `json:"avg_wear_count"`
}

// 分类相关DTO
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parent_id"`
	Icon        string `json:"icon"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ParentID    *uint   `json:"parent_id"`
	Icon        *string `json:"icon"`
	SortOrder   *int    `json:"sort_order"`
	IsActive    *bool   `json:"is_active"`
}

type CategoryResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	ParentID    *uint              `json:"parent_id"`
	ParentName  string             `json:"parent_name,omitempty"`
	Icon        string             `json:"icon"`
	SortOrder   int                `json:"sort_order"`
	IsActive    bool               `json:"is_active"`
	ItemCount   int64              `json:"item_count"`
	Children    []CategoryResponse `json:"children,omitempty"`
}

type CategoryTreeNode struct {
	CategoryResponse
	Children []CategoryTreeNode `json:"children"`
}

// 标签相关DTO
type CreateTagRequest struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
}

type UpdateTagRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
	Icon        *string `json:"icon"`
	IsActive    *bool   `json:"is_active"`
}

type TagResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	IsSystem    bool   `json:"is_system"`
	IsActive    bool   `json:"is_active"`
	ItemCount   int64  `json:"item_count"`
	UserID      *uint  `json:"user_id,omitempty"`
}

type TagStatsItem struct {
	TagName   string  `json:"tag_name"`
	TagType   string  `json:"tag_type"`
	Count     int64   `json:"count"`
	UsageRate float64 `json:"usage_rate"`
}

// 购买记录相关DTO
type UpdatePurchaseRecordRequest struct {
	PurchasePrice  *float64   `json:"purchase_price"`
	OriginalPrice  *float64   `json:"original_price"`
	StoreName      *string    `json:"store_name"`
	StoreLocation  *string    `json:"store_location"`
	OnlineStore    *string    `json:"online_store"`
	OrderNumber    *string    `json:"order_number"`
	PurchaseDate   *time.Time `json:"purchase_date"`
	PaymentMethod  *string    `json:"payment_method"`
	Currency       *string    `json:"currency"`
	ReceiptURL     *string    `json:"receipt_url"`
	WarrantyPeriod *int       `json:"warranty_period"`
	ReturnPolicy   *string    `json:"return_policy"`
	Notes          *string    `json:"notes"`
	Tags           []string   `json:"tags"`
}

type PurchaseRecordResponse struct {
	ID             uint       `json:"id"`
	ClothingItemID uint       `json:"clothing_item_id"`
	ItemName       string     `json:"item_name"`
	PurchasePrice  float64    `json:"purchase_price"`
	OriginalPrice  float64    `json:"original_price"`
	Discount       float64    `json:"discount"`
	StoreName      string     `json:"store_name"`
	StoreLocation  string     `json:"store_location"`
	OnlineStore    string     `json:"online_store"`
	OrderNumber    string     `json:"order_number"`
	PurchaseDate   time.Time  `json:"purchase_date"`
	PaymentMethod  string     `json:"payment_method"`
	Currency       string     `json:"currency"`
	ReceiptURL     string     `json:"receipt_url"`
	WarrantyPeriod int        `json:"warranty_period"`
	WarrantyExpiry *time.Time `json:"warranty_expiry"`
	ReturnPolicy   string     `json:"return_policy"`
	Notes          string     `json:"notes"`
	Tags           []string   `json:"tags"`
	CreatedAt      time.Time  `json:"created_at"`
}

type SpendingStatsResponse struct {
	TotalSpent        float64              `json:"total_spent"`
	MonthlySpending   map[string]float64   `json:"monthly_spending"`
	CategorySpending  map[string]float64   `json:"category_spending"`
	BrandSpending     map[string]float64   `json:"brand_spending"`
	AverageItemPrice  float64              `json:"average_item_price"`
	MostExpensiveItem *ClothingItemSummary `json:"most_expensive_item"`
}

// 保养记录相关DTO
type UpdateMaintenanceRecordRequest struct {
	MaintenanceType    *string    `json:"maintenance_type"`
	Cost               *float64   `json:"cost"`
	MaintenanceDate    *time.Time `json:"maintenance_date"`
	ServiceProvider    *string    `json:"service_provider"`
	ServiceLocation    *string    `json:"service_location"`
	BeforeCondition    *string    `json:"before_condition"`
	AfterCondition     *string    `json:"after_condition"`
	EffectivenessScore *int       `json:"effectiveness_score"`
	Notes              *string    `json:"notes"`
	Images             []string   `json:"images"`
}

type MaintenanceRecordResponse struct {
	ID                  uint       `json:"id"`
	ClothingItemID      uint       `json:"clothing_item_id"`
	ItemName            string     `json:"item_name"`
	MaintenanceType     string     `json:"maintenance_type"`
	Cost                float64    `json:"cost"`
	MaintenanceDate     time.Time  `json:"maintenance_date"`
	ServiceProvider     string     `json:"service_provider"`
	ServiceLocation     string     `json:"service_location"`
	BeforeCondition     string     `json:"before_condition"`
	AfterCondition      string     `json:"after_condition"`
	EffectivenessScore  int        `json:"effectiveness_score"`
	Notes               string     `json:"notes"`
	Images              []string   `json:"images"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date"`
	CreatedAt           time.Time  `json:"created_at"`
}

type MaintenanceReminderResponse struct {
	ID                  uint       `json:"id"`
	ClothingItemID      uint       `json:"clothing_item_id"`
	ItemName            string     `json:"item_name"`
	ItemImageURL        string     `json:"item_image_url"`
	MaintenanceType     string     `json:"maintenance_type"`
	NextMaintenanceDate time.Time  `json:"next_maintenance_date"`
	DaysOverdue         int        `json:"days_overdue"`
	LastMaintenanceDate *time.Time `json:"last_maintenance_date"`
}

type MaintenanceStatsResponse struct {
	TotalCost            float64            `json:"total_cost"`
	MaintenanceCount     int64              `json:"maintenance_count"`
	CostByType           map[string]float64 `json:"cost_by_type"`
	FrequencyByType      map[string]int64   `json:"frequency_by_type"`
	AverageEffectiveness float64            `json:"average_effectiveness"`
	UpcomingCount        int64              `json:"upcoming_count"`
	OverdueCount         int64              `json:"overdue_count"`
}

// 穿着记录相关DTO
type UpdateWearRecordRequest struct {
	WearDate         *time.Time `json:"wear_date"`
	DurationHours    *int       `json:"duration_hours"`
	Occasion         *string    `json:"occasion"`
	WeatherCondition *string    `json:"weather_condition"`
	Temperature      *float64   `json:"temperature"`
	Activity         *string    `json:"activity"`
	ComfortRating    *int       `json:"comfort_rating"`
	StyleRating      *int       `json:"style_rating"`
	Notes            *string    `json:"notes"`
	Photos           []string   `json:"photos"`
	Location         *string    `json:"location"`
	Companions       []string   `json:"companions"`
	Mood             *string    `json:"mood"`
	WearIntensity    *string    `json:"wear_intensity"`
}

type WearRecordResponse struct {
	ID               uint      `json:"id"`
	ClothingItemID   uint      `json:"clothing_item_id"`
	ItemName         string    `json:"item_name"`
	ItemImageURL     string    `json:"item_image_url"`
	WearDate         time.Time `json:"wear_date"`
	DurationHours    int       `json:"duration_hours"`
	Occasion         string    `json:"occasion"`
	WeatherCondition string    `json:"weather_condition"`
	Temperature      float64   `json:"temperature"`
	Activity         string    `json:"activity"`
	ComfortRating    int       `json:"comfort_rating"`
	StyleRating      int       `json:"style_rating"`
	Notes            string    `json:"notes"`
	Photos           []string  `json:"photos"`
	Location         string    `json:"location"`
	Companions       []string  `json:"companions"`
	Mood             string    `json:"mood"`
	WearIntensity    string    `json:"wear_intensity"`
	CreatedAt        time.Time `json:"created_at"`
}

type WearStatsResponse struct {
	TotalWears      int64                 `json:"total_wears"`
	TotalHours      int64                 `json:"total_hours"`
	AverageComfort  float64               `json:"average_comfort"`
	AverageStyle    float64               `json:"average_style"`
	WearsByOccasion map[string]int64      `json:"wears_by_occasion"`
	WearsByWeather  map[string]int64      `json:"wears_by_weather"`
	WearsByMonth    map[string]int64      `json:"wears_by_month"`
	MostWornItems   []ClothingItemSummary `json:"most_worn_items"`
	LeastWornItems  []ClothingItemSummary `json:"least_worn_items"`
}

type ComfortAnalysisResponse struct {
	OverallComfort     float64               `json:"overall_comfort"`
	ComfortByCategory  map[string]float64    `json:"comfort_by_category"`
	ComfortByBrand     map[string]float64    `json:"comfort_by_brand"`
	ComfortByMaterial  map[string]float64    `json:"comfort_by_material"`
	ComfortByOccasion  map[string]float64    `json:"comfort_by_occasion"`
	ComfortableItems   []ClothingItemSummary `json:"comfortable_items"`
	UncomfortableItems []ClothingItemSummary `json:"uncomfortable_items"`
}

// BrandStatsItem 品牌统计项
type BrandStatsItem struct {
	BrandName    string  `json:"brand_name"`
	Count        int64   `json:"count"`
	TotalValue   float64 `json:"total_value"`
	AvgWearCount float64 `json:"avg_wear_count"`
}

// ColorStatsItem 颜色统计项
type ColorStatsItem struct {
	Color        string  `json:"color"`
	Count        int64   `json:"count"`
	TotalValue   float64 `json:"total_value"`
	AvgWearCount float64 `json:"avg_wear_count"`
}
