package dto

import "time"

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `form:"page" binding:"min=1" json:"page"`
	PageSize int `form:"page_size" binding:"min=1,max=100" json:"page_size"`
}

// SearchRequest 通用搜索请求
type SearchRequest struct {
	Query     string            `json:"query" form:"query"`
	Filters   map[string]string `json:"filters" form:"filters"`
	SortBy    string            `json:"sort_by" form:"sort_by"`
	SortOrder string            `json:"sort_order" form:"sort_order"` // asc, desc
	PaginationRequest
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
	Errors       []string `json:"errors,omitempty"`
	Results      []uint   `json:"results,omitempty"`
}

// ExportRequest 导出请求
type ExportRequest struct {
	Type     string            `json:"type" binding:"required"` // csv, excel, json, pdf
	Format   string            `json:"format"`                  // detailed, summary
	Filters  map[string]string `json:"filters,omitempty"`
	DateFrom *time.Time        `json:"date_from,omitempty"`
	DateTo   *time.Time        `json:"date_to,omitempty"`
}

// ImportRequest 导入请求
type ImportRequest struct {
	Type        string            `json:"type" binding:"required"` // csv, excel, json
	FileURL     string            `json:"file_url"`
	Options     map[string]string `json:"options,omitempty"`
	DryRun      bool              `json:"dry_run"`
	SkipErrors  bool              `json:"skip_errors"`
}