package dto

// SearchRequest 通用搜索请求
type SearchRequest struct {
	Query     string            `json:"query" form:"query"`
	Filters   map[string]string `json:"filters" form:"filters"`
	SortBy    string            `json:"sort_by" form:"sort_by"`
	SortOrder string            `json:"sort_order" form:"sort_order"` // asc, desc
	Page      int               `json:"page" form:"page" binding:"min=1"`
	PageSize  int               `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// SearchResponse 通用搜索响应
type SearchResponse struct {
	Results    []SearchResultItem `json:"results"`
	Total      int64              `json:"total"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
	TotalPages int                `json:"total_pages"`
	Facets     []SearchFacet      `json:"facets,omitempty"`
	Suggestions []string          `json:"suggestions,omitempty"`
}

// SearchResultItem 搜索结果项
type SearchResultItem struct {
	ID          uint                   `json:"id"`
	Type        string                 `json:"type"` // clothing, outfit, tag, category
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	ImageURL    string                 `json:"image_url,omitempty"`
	Score       float64                `json:"score"`
	Highlights  map[string][]string    `json:"highlights,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// SearchFacet 搜索分面
type SearchFacet struct {
	Name   string            `json:"name"`
	Values []SearchFacetValue `json:"values"`
}

// SearchFacetValue 搜索分面值
type SearchFacetValue struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

// FilterRequest 过滤请求
type FilterRequest struct {
	CategoryIDs   []uint    `json:"category_ids" form:"category_ids"`
	TagIDs        []uint    `json:"tag_ids" form:"tag_ids"`
	Brands        []string  `json:"brands" form:"brands"`
	Colors        []string  `json:"colors" form:"colors"`
	Materials     []string  `json:"materials" form:"materials"`
	Sizes         []string  `json:"sizes" form:"sizes"`
	Conditions    []string  `json:"conditions" form:"conditions"`
	MinPrice      *float64  `json:"min_price" form:"min_price"`
	MaxPrice      *float64  `json:"max_price" form:"max_price"`
	MinWearCount  *int      `json:"min_wear_count" form:"min_wear_count"`
	MaxWearCount  *int      `json:"max_wear_count" form:"max_wear_count"`
	IsFavorite    *bool     `json:"is_favorite" form:"is_favorite"`
	IsActive      *bool     `json:"is_active" form:"is_active"`
	DateFrom      string    `json:"date_from" form:"date_from"`
	DateTo        string    `json:"date_to" form:"date_to"`
}

// SortOption 排序选项
type SortOption struct {
	Field       string `json:"field"`
	Order       string `json:"order"` // asc, desc
	DisplayName string `json:"display_name"`
}

// BulkOperationRequest 批量操作请求
type BulkOperationRequest struct {
	IDs       []uint                 `json:"ids" binding:"required"`
	Operation string                 `json:"operation" binding:"required"` // delete, update, tag, untag
	Data      map[string]interface{} `json:"data,omitempty"`
}

// BulkOperationResponse 批量操作响应
type BulkOperationResponse struct {
	SuccessCount int      `json:"success_count"`
	FailureCount int      `json:"failure_count"`
	SuccessIDs   []uint   `json:"success_ids"`
	FailureIDs   []uint   `json:"failure_ids"`
	Errors       []string `json:"errors,omitempty"`
}

// ExportRequest 导出请求
type ExportRequest struct {
	Type     string            `json:"type" binding:"required"` // csv, excel, json, pdf
	Format   string            `json:"format"`                  // detailed, summary
	Filters  FilterRequest     `json:"filters"`
	Fields   []string          `json:"fields"`
	Options  map[string]string `json:"options"`
}

// ExportResponse 导出响应
type ExportResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"` // pending, processing, completed, failed
	FileURL     string `json:"file_url,omitempty"`
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size,omitempty"`
	RecordCount int64  `json:"record_count"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at,omitempty"`
}

// ImportRequest 导入请求
type ImportRequest struct {
	Type        string            `json:"type" binding:"required"` // csv, excel, json
	FileURL     string            `json:"file_url"`
	Mapping     map[string]string `json:"mapping"`     // 字段映射
	Options     map[string]string `json:"options"`     // 导入选项
	DryRun      bool              `json:"dry_run"`     // 是否为试运行
	SkipErrors  bool              `json:"skip_errors"` // 是否跳过错误
}

// ImportResponse 导入响应
type ImportResponse struct {
	ID            string        `json:"id"`
	Status        string        `json:"status"` // pending, processing, completed, failed
	TotalRecords  int64         `json:"total_records"`
	ProcessedRecords int64      `json:"processed_records"`
	SuccessCount  int64         `json:"success_count"`
	ErrorCount    int64         `json:"error_count"`
	Errors        []ImportError `json:"errors,omitempty"`
	CreatedAt     string        `json:"created_at"`
	CompletedAt   string        `json:"completed_at,omitempty"`
}

// ImportError 导入错误
type ImportError struct {
	Row     int    `json:"row"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// ValidationRequest 验证请求
type ValidationRequest struct {
	Data   map[string]interface{} `json:"data" binding:"required"`
	Rules  map[string]string      `json:"rules"`
	Fields []string               `json:"fields"`
}

// ValidationResponse 验证响应
type ValidationResponse struct {
	Valid   bool              `json:"valid"`
	Errors  map[string]string `json:"errors,omitempty"`
	Warnings map[string]string `json:"warnings,omitempty"`
}

// HealthCheckResponse 健康检查响应
type HealthCheckResponse struct {
	Status    string            `json:"status"` // healthy, unhealthy, degraded
	Version   string            `json:"version"`
	Timestamp string            `json:"timestamp"`
	Services  map[string]string `json:"services"`
	Uptime    string            `json:"uptime"`
}