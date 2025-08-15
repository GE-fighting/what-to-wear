package dto

import "time"

// DashboardStatsResponse 仪表板统计响应
type DashboardStatsResponse struct {
	UserStats     UserStatsResponse     `json:"user_stats"`
	ClothingStats ClothingStatsResponse `json:"clothing_stats"`
	WearStats     WearStatsResponse     `json:"wear_stats"`
	SpendingStats SpendingStatsResponse `json:"spending_stats"`
	RecentActivity []ActivityItem       `json:"recent_activity"`
}

// ActivityItem 活动项
type ActivityItem struct {
	ID          uint      `json:"id"`
	Type        string    `json:"type"` // wear, purchase, maintenance, outfit
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// AnalyticsRequest 分析请求
type AnalyticsRequest struct {
	StartDate string   `json:"start_date" form:"start_date"`
	EndDate   string   `json:"end_date" form:"end_date"`
	Metrics   []string `json:"metrics" form:"metrics"`
	GroupBy   string   `json:"group_by" form:"group_by"` // day, week, month, year
}

// AnalyticsResponse 分析响应
type AnalyticsResponse struct {
	Period    string                 `json:"period"`
	StartDate string                 `json:"start_date"`
	EndDate   string                 `json:"end_date"`
	Metrics   map[string]interface{} `json:"metrics"`
	Charts    []ChartData           `json:"charts"`
	Summary   AnalyticsSummary      `json:"summary"`
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

// TrendAnalysisResponse 趋势分析响应
type TrendAnalysisResponse struct {
	WearTrends      []TrendItem `json:"wear_trends"`
	SpendingTrends  []TrendItem `json:"spending_trends"`
	CategoryTrends  []TrendItem `json:"category_trends"`
	SeasonalTrends  []TrendItem `json:"seasonal_trends"`
	Predictions     []Prediction `json:"predictions"`
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
	Type        string  `json:"type"`        // wear, spending, maintenance
	Period      string  `json:"period"`      // next_week, next_month, next_season
	Value       float64 `json:"value"`
	Confidence  float64 `json:"confidence"`
	Description string  `json:"description"`
}

// ReportRequest 报告请求
type ReportRequest struct {
	Type      string `json:"type" binding:"required"` // monthly, yearly, custom
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Format    string `json:"format"` // pdf, excel, json
	Sections  []string `json:"sections"` // overview, clothing, spending, wear, maintenance
}

// ReportResponse 报告响应
type ReportResponse struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Period      string    `json:"period"`
	GeneratedAt time.Time `json:"generated_at"`
	FileURL     string    `json:"file_url,omitempty"`
	Data        ReportData `json:"data,omitempty"`
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
	CategoryBreakdown []CategoryStatsItem `json:"category_breakdown"`
	BrandBreakdown    []BrandStatsItem    `json:"brand_breakdown"`
	ColorBreakdown    []ColorStatsItem    `json:"color_breakdown"`
	MostWorn          []ClothingItemSummary `json:"most_worn"`
	LeastWorn         []ClothingItemSummary `json:"least_worn"`
	NewAdditions      []ClothingItemSummary `json:"new_additions"`
}

// ReportSpending 报告支出部分
type ReportSpending struct {
	TotalSpent       float64            `json:"total_spent"`
	MonthlyBreakdown map[string]float64 `json:"monthly_breakdown"`
	CategorySpending map[string]float64 `json:"category_spending"`
	BrandSpending    map[string]float64 `json:"brand_spending"`
	AverageItemPrice float64            `json:"average_item_price"`
	MostExpensive    *ClothingItemSummary `json:"most_expensive"`
}

// ReportWear 报告穿着部分
type ReportWear struct {
	TotalWears       int64                 `json:"total_wears"`
	AveragePerItem   float64               `json:"average_per_item"`
	WearsByCategory  map[string]int64      `json:"wears_by_category"`
	WearsByOccasion  map[string]int64      `json:"wears_by_occasion"`
	WearsByWeather   map[string]int64      `json:"wears_by_weather"`
	ComfortAnalysis  ComfortAnalysisResponse `json:"comfort_analysis"`
}

// ReportMaintenance 报告保养部分
type ReportMaintenance struct {
	TotalCost        float64            `json:"total_cost"`
	MaintenanceCount int64              `json:"maintenance_count"`
	CostByType       map[string]float64 `json:"cost_by_type"`
	UpcomingTasks    []MaintenanceReminderResponse `json:"upcoming_tasks"`
	OverdueTasks     []MaintenanceReminderResponse `json:"overdue_tasks"`
}