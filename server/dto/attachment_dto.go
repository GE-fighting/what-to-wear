package dto

import (
	"mime/multipart"
	"time"
	"what-to-wear/server/models"
)

// UploadAttachmentRequest 上传附件请求
type UploadAttachmentRequest struct {
	File        *multipart.FileHeader `form:"file" binding:"required"`
	EntityType  models.EntityType     `form:"entity_type" binding:"required"`
	EntityID    uint                  `form:"entity_id" binding:"required"`
	UserID      uint                  `form:"user_id" binding:"required"`
	Description string                `form:"description"`
	Tags        []string              `form:"tags"`
	IsPublic    bool                  `form:"is_public"`
	SortOrder   int                   `form:"sort_order"`
}

// BatchUploadRequest 批量上传附件请求
type BatchUploadRequest struct {
	Files       []*multipart.FileHeader `form:"files" binding:"required"`
	EntityType  models.EntityType       `form:"entity_type" binding:"required"`
	EntityID    uint                    `form:"entity_id" binding:"required"`
	UserID      uint                    `form:"user_id" binding:"required"`
	Description string                  `form:"description"`
	Tags        []string                `form:"tags"`
	IsPublic    bool                    `form:"is_public"`
}

// UpdateAttachmentRequest 更新附件请求
type UpdateAttachmentRequest struct {
	Description *string  `json:"description"`
	Tags        []string `json:"tags"`
	IsPublic    *bool    `json:"is_public"`
	SortOrder   *int     `json:"sort_order"`
}

// AttachmentResponse 附件响应
type AttachmentResponse struct {
	ID             uint      `json:"id"`
	OriginalName   string    `json:"original_name"`
	FileName       string    `json:"file_name"`
	FileSize       int64     `json:"file_size"`
	MimeType       string    `json:"mime_type"`
	AttachmentType string    `json:"attachment_type"`
	EntityType     string    `json:"entity_type"`
	EntityID       uint      `json:"entity_id"`
	PublicURL      string    `json:"public_url"`
	Width          *int      `json:"width,omitempty"`
	Height         *int      `json:"height,omitempty"`
	Duration       *int      `json:"duration,omitempty"`
	Thumbnail      string    `json:"thumbnail,omitempty"`
	Description    string    `json:"description"`
	Tags           []string  `json:"tags"`
	SortOrder      int       `json:"sort_order"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// BatchUploadResponse 批量上传响应
type BatchUploadResponse struct {
	SuccessCount int                  `json:"success_count"`
	FailureCount int                  `json:"failure_count"`
	Attachments  []AttachmentResponse `json:"attachments"`
	Errors       []string             `json:"errors"`
}

// AttachmentListResponse 附件列表响应
type AttachmentListResponse struct {
	Attachments []AttachmentResponse `json:"attachments"`
	Total       int                  `json:"total"`
	Page        int                  `json:"page"`
	PageSize    int                  `json:"page_size"`
}

// AttachmentSummary 附件摘要
type AttachmentSummary struct {
	ID             uint   `json:"id"`
	OriginalName   string `json:"original_name"`
	FileSize       int64  `json:"file_size"`
	AttachmentType string `json:"attachment_type"`
	PublicURL      string `json:"public_url"`
	Thumbnail      string `json:"thumbnail,omitempty"`
}

// UpdateAttachmentOrderRequest 更新附件排序请求
type UpdateAttachmentOrderRequest struct {
	AttachmentID uint `json:"attachment_id" binding:"required"`
	SortOrder    int  `json:"sort_order" binding:"required"`
}

// AttachmentInfoResponse 附件详细信息响应
type AttachmentInfoResponse struct {
	AttachmentResponse
	Metadata    map[string]string `json:"metadata"`
	UploadedBy  string            `json:"uploaded_by"`
	StorageInfo StorageInfo       `json:"storage_info"`
}

// StorageInfo 存储信息
type StorageInfo struct {
	Provider   string `json:"provider"`
	BucketName string `json:"bucket_name,omitempty"`
	ObjectKey  string `json:"object_key,omitempty"`
	PrivateURL string `json:"private_url,omitempty"`
}

// UpdateAttachmentInfoRequest 更新附件信息请求
type UpdateAttachmentInfoRequest struct {
	FileName    *string  `json:"file_name"`
	Description *string  `json:"description"`
	Tags        []string `json:"tags"`
}

// BatchDeleteAttachmentRequest 批量删除附件请求
type BatchDeleteAttachmentRequest struct {
	AttachmentIDs []uint `json:"attachment_ids" binding:"required,min=1"`
}

// AttachmentStatsResponse 附件统计响应
type AttachmentStatsResponse struct {
	TotalAttachments int64             `json:"total_attachments"`
	TotalSize        int64             `json:"total_size"`
	ByType           map[string]int64  `json:"by_type"`
	ByEntity         map[string]int64  `json:"by_entity"`
	StorageUsage     StorageUsageStats `json:"storage_usage"`
}

// StorageUsageStats 存储使用统计
type StorageUsageStats struct {
	Images int64 `json:"images"`
	Videos int64 `json:"videos"`
	Files  int64 `json:"files"`
}
