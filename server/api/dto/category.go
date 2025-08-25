package dto

import "time"

// CategoryResponse 分类响应
type CategoryDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    *uint     `json:"parent_id,omitempty"`
	ParentName  string    `json:"parent_name,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	ItemCount   int64     `json:"item_count,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryTreeNode 分类树节点
type CategoryTreeNode struct {
	CategoryDTO
	Children []CategoryTreeNode `json:"children"`
}

// CreateCategoryRequest 创建分类请求
type CreateCategoryDTO struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	Description string `json:"description" binding:"max=200"`
	ParentID    *uint  `json:"parent_id,omitempty"`
	Icon        string `json:"icon,omitempty" binding:"max=50"`
	SortOrder   int    `json:"sort_order" binding:"min=0"`
}

// UpdateCategoryRequest 更新分类请求
type UpdateCategoryDTO struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=1,max=50"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=200"`
	ParentID    *uint   `json:"parent_id,omitempty"`
	Icon        *string `json:"icon,omitempty" binding:"omitempty,max=50"`
	SortOrder   *int    `json:"sort_order,omitempty" binding:"omitempty,min=0"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CategoryListRequest 分类列表请求
type CategoryListRequest struct {
	ParentID   *uint  `form:"parent_id"`
	IsActive   *bool  `form:"is_active"`
	Search     string `form:"search"`
	WithCounts bool   `form:"with_counts"` // 是否包含衣物数量统计
	TreeView   bool   `form:"tree_view"`   // 是否返回树形结构
	SearchRequest
}

// CategoryBatchRequest 分类批量操作请求
type CategoryBatchRequest struct {
	CategoryIDs []uint `json:"category_ids" binding:"required,min=1"`
	Operation   string `json:"operation" binding:"required,oneof=activate deactivate delete"`
}

// CategoryMoveRequest 分类移动请求
type CategoryMoveRequest struct {
	CategoryID   uint  `json:"category_id" binding:"required"`
	NewParentID  *uint `json:"new_parent_id,omitempty"`
	NewSortOrder int   `json:"new_sort_order" binding:"min=0"`
}

// CategoryMergeRequest 分类合并请求
type CategoryMergeRequest struct {
	SourceCategoryIDs []uint `json:"source_category_ids" binding:"required,min=1"`
	TargetCategoryID  uint   `json:"target_category_id" binding:"required"`
	DeleteSources     bool   `json:"delete_sources"` // 合并后是否删除源分类
}

// CategoryValidationResponse 分类验证响应
type CategoryValidationResponse struct {
	IsValid     bool     `json:"is_valid"`
	Errors      []string `json:"errors,omitempty"`
	Warnings    []string `json:"warnings,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

// CategorySummaryResponse 分类摘要响应
type CategorySummaryResponse struct {
	TotalCategories   int64               `json:"total_categories"`
	ActiveCategories  int64               `json:"active_categories"`
	RootCategories    int64               `json:"root_categories"`
	MaxDepth          int                 `json:"max_depth"`
	TotalItems        int64               `json:"total_items"`
	CategoryBreakdown []CategoryStatsItem `json:"category_breakdown"`
	PopularCategories []CategoryDTO       `json:"popular_categories"`
	EmptyCategories   []CategoryDTO       `json:"empty_categories"`
}

// CategoryImportRequest 分类导入请求
type CategoryImportRequest struct {
	Categories []CreateCategoryDTO `json:"categories" binding:"required,min=1"`
	Mode       string              `json:"mode" binding:"required,oneof=merge replace append"`
	DryRun     bool                `json:"dry_run"` // 是否只验证不实际导入
}

// CategoryExportRequest 分类导出请求
type CategoryExportRequest struct {
	Format       string `json:"format" binding:"required,oneof=json csv excel"`
	IncludeItems bool   `json:"include_items"` // 是否包含分类下的衣物
	IncludeStats bool   `json:"include_stats"` // 是否包含统计信息
}
