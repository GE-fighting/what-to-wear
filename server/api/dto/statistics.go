package dto

import (
	"time"
	"what-to-wear/server/api"
)

// DashboardStatsDTO 仪表板统计DTO
type DashboardStatsDTO struct {
	UserStats      UserStatsDTO     `json:"user_stats"`
	ClothingStats  ClothingStatsDTO `json:"clothing_stats"`
	OutfitStats    OutfitStatsDTO   `json:"outfit_stats"`
	SpendingStats  SpendingStatsDTO `json:"spending_stats"`
	RecentActivity []ActivityItem   `json:"recent_activity"`
}

// ActivityItem 活动项
type ActivityItem struct {
	ID          uint           `json:"id"`
	Type        api.EntityType `json:"type"` // clothing_item, outfit, purchase, maintenance, wear_record
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ImageURL    string         `json:"image_url,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
}

// AnalyticsDTO 分析DTO
type AnalyticsDTO struct {
	Period    string                 `json:"period"`
	StartDate time.Time              `json:"start_date"`
	EndDate   time.Time              `json:"end_date"`
	Metrics   map[string]interface{} `json:"metrics"`
	Charts    []ChartData            `json:"charts"`
	Summary   AnalyticsSummary       `json:"summary"`
}

// ChartData 图表数据
type ChartData struct {
	Type   string                   `json:"type"` // line, bar, pie, doughnut
	Title  string                   `json:"title"`
	Labels []string                 `json:"labels"`
	Data   []map[string]interface{} `json:"data"`
}

// AnalyticsSummary 分析摘要
type AnalyticsSummary struct {
	TotalItems      int64   `json:"total_items"`
	TotalWears      int64   `json:"total_wears"`
	TotalSpent      float64 `json:"total_spent"`
	AverageWearRate float64 `json:"average_wear_rate"`
	TopCategory     string  `json:"top_category"`
	TopBrand        string  `json:"top_brand"`
	TopColor        string  `json:"top_color"`
}

// TrendAnalysisDTO 趋势分析DTO
type TrendAnalysisDTO struct {
	WearTrends     []TrendItem  `json:"wear_trends"`
	SpendingTrends []TrendItem  `json:"spending_trends"`
	CategoryTrends []TrendItem  `json:"category_trends"`
	SeasonalTrends []TrendItem  `json:"seasonal_trends"`
	Predictions    []Prediction `json:"predictions"`
}

// TrendItem 趋势项
type TrendItem struct {
	Period    string  `json:"period"`
	Value     float64 `json:"value"`
	Change    float64 `json:"change"`    // 变化百分比
	Direction string  `json:"direction"` // up, down, stable
}

// Prediction 预测
type Prediction struct {
	Type        string  `json:"type"`   // wear, spending, maintenance
	Period      string  `json:"period"` // next_week, next_month, next_season
	Value       float64 `json:"value"`
	Confidence  float64 `json:"confidence"`
	Description string  `json:"description"`
}

// ReportDTO 报告DTO
type ReportDTO struct {
	Type      string    `json:"type" binding:"required"` // monthly, yearly, custom
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Format    string    `json:"format"`   // pdf, excel, json
	Sections  []string  `json:"sections"` // overview, clothing, spending, wear, maintenance
}

// ReportData 报告数据
type ReportData struct {
	Overview    ReportOverview    `json:"overview"`
	Clothing    ReportClothing    `json:"clothing"`
	Spending    ReportSpending    `json:"spending"`
	Wear        ReportWear        `json:"wear"`
	Maintenance ReportMaintenance `json:"maintenance"`
}

// ReportOverview 报告概览
type ReportOverview struct {
	Period          string  `json:"period"`
	TotalItems      int64   `json:"total_items"`
	NewItems        int64   `json:"new_items"`
	TotalWears      int64   `json:"total_wears"`
	TotalSpent      float64 `json:"total_spent"`
	MaintenanceCost float64 `json:"maintenance_cost"`
	CostPerWear     float64 `json:"cost_per_wear"`
}

// ReportClothing 报告衣物部分
type ReportClothing struct {
	CategoryBreakdown []CategoryStatsItem   `json:"category_breakdown"`
	BrandBreakdown    []BrandStatsItem      `json:"brand_breakdown"`
	ColorBreakdown    []ColorStatsItem      `json:"color_breakdown"`
	MostWorn          []ClothingItemSummary `json:"most_worn"`
	LeastWorn         []ClothingItemSummary `json:"least_worn"`
	NewAdditions      []ClothingItemSummary `json:"new_additions"`
}

// ReportSpending 报告支出部分
type ReportSpending struct {
	TotalSpent       float64              `json:"total_spent"`
	MonthlyBreakdown map[string]float64   `json:"monthly_breakdown"`
	CategorySpending map[string]float64   `json:"category_spending"`
	BrandSpending    map[string]float64   `json:"brand_spending"`
	AverageItemPrice float64              `json:"average_item_price"`
	MostExpensive    *ClothingItemSummary `json:"most_expensive"`
}

// ReportWear 报告穿着部分
type ReportWear struct {
	TotalWears      int64                     `json:"total_wears"`
	AveragePerItem  float64                   `json:"average_per_item"`
	WearsByCategory map[string]int64          `json:"wears_by_category"`
	WearsByOccasion map[string]int64          `json:"wears_by_occasion"`
	WearsByWeather  map[api.WeatherType]int64 `json:"wears_by_weather"`
	ComfortAnalysis ComfortAnalysisDTO        `json:"comfort_analysis"`
}

// ReportMaintenance 报告保养部分
type ReportMaintenance struct {
	TotalCost        float64                  `json:"total_cost"`
	MaintenanceCount int64                    `json:"maintenance_count"`
	CostByType       map[string]float64       `json:"cost_by_type"`
	UpcomingTasks    []MaintenanceReminderDTO `json:"upcoming_tasks"`
	OverdueTasks     []MaintenanceReminderDTO `json:"overdue_tasks"`
}

// CategoryStatsItem 分类统计项
type CategoryStatsItem struct {
	CategoryName string  `json:"category_name"`
	Count        int64   `json:"count"`
	Percentage   float64 `json:"percentage"`
}

// BrandStatsItem 品牌统计项
type BrandStatsItem struct {
	BrandName  string  `json:"brand_name"`
	Count      int64   `json:"count"`
	TotalSpent float64 `json:"total_spent"`
	Percentage float64 `json:"percentage"`
}

// ColorStatsItem 颜色统计项
type ColorStatsItem struct {
	ColorName  string  `json:"color_name"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}

// ClothingItemSummary 衣物摘要
type ClothingItemSummary struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name"`
	Brand         string             `json:"brand"`
	Color         string             `json:"color"`
	CategoryName  string             `json:"category_name"`
	ImageURL      string             `json:"image_url"`
	Status        api.ClothingStatus `json:"status"`
	WearCount     int                `json:"wear_count"`
	LastWornDate  *time.Time         `json:"last_worn_date"`
	PurchasePrice *float64           `json:"purchase_price"`
}

// ComfortAnalysisDTO 舒适度分析DTO
type ComfortAnalysisDTO struct {
	AverageComfort         float64                     `json:"average_comfort"`
	AverageStyle           float64                     `json:"average_style"`
	AverageAppropriateness float64                     `json:"average_appropriateness"`
	ComfortByCategory      map[string]float64          `json:"comfort_by_category"`
	ComfortByWeather       map[api.WeatherType]float64 `json:"comfort_by_weather"`
}

// MaintenanceReminderDTO 保养提醒DTO
type MaintenanceReminderDTO struct {
	ID                  uint      `json:"id"`
	ClothingItemID      uint      `json:"clothing_item_id"`
	ClothingItemName    string    `json:"clothing_item_name"`
	MaintenanceType     string    `json:"maintenance_type"`
	NextMaintenanceDate time.Time `json:"next_maintenance_date"`
	DaysOverdue         int       `json:"days_overdue"`
	Priority            string    `json:"priority"` // low, medium, high, urgent
}

// SpendingStatsDTO 支出统计DTO
type SpendingStatsDTO struct {
	TotalSpent        float64               `json:"total_spent"`
	MonthlySpending   map[string]float64    `json:"monthly_spending"`
	CategorySpending  map[string]float64    `json:"category_spending"`
	BrandSpending     map[string]float64    `json:"brand_spending"`
	AverageItemPrice  float64               `json:"average_item_price"`
	CostPerWear       float64               `json:"cost_per_wear"`
	MostExpensiveItem *ClothingItemSummary  `json:"most_expensive_item"`
	BestValueItems    []ClothingItemSummary `json:"best_value_items"`
}
