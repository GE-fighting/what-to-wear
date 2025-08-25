package dto

import (
	"mime/multipart"
	"time"
	"what-to-wear/server/api"
)

// UploadAttachmentDTO 上传附件DTO
type UploadAttachmentDTO struct {
	File        *multipart.FileHeader `form:"file" binding:"required"`
	EntityType  api.EntityType        `form:"entity_type" binding:"required"`
	EntityID    uint                  `form:"entity_id" binding:"required"`
	UserID      uint                  `form:"user_id" binding:"required"`
	Description string                `form:"description"`
	Tags        []string              `form:"tags"`
	IsPublic    bool                  `form:"is_public"`
	SortOrder   int                   `form:"sort_order"`
}

// BatchUploadDTO 批量上传附件DTO
type BatchUploadDTO struct {
	Files       []*multipart.FileHeader `form:"files" binding:"required"`
	EntityType  api.EntityType          `form:"entity_type" binding:"required"`
	EntityID    uint                    `form:"entity_id" binding:"required"`
	UserID      uint                    `form:"user_id" binding:"required"`
	Description string                  `form:"description"`
	Tags        []string                `form:"tags"`
	IsPublic    bool                    `form:"is_public"`
}

// UpdateAttachmentDTO 更新附件DTO
type UpdateAttachmentDTO struct {
	Description *string  `json:"description"`
	Tags        []string `json:"tags"`
	IsPublic    *bool    `json:"is_public"`
	SortOrder   *int     `json:"sort_order"`
}

// AttachmentDTO 附件DTO
type AttachmentDTO struct {
	ID             uint               `json:"id"`
	OriginalName   string             `json:"original_name"`
	FileName       string             `json:"file_name"`
	FileSize       int64              `json:"file_size"`
	MimeType       string             `json:"mime_type"`
	AttachmentType api.AttachmentType `json:"attachment_type"`
	EntityType     api.EntityType     `json:"entity_type"`
	EntityID       uint               `json:"entity_id"`
	PublicURL      string             `json:"public_url"`
	Width          *int               `json:"width,omitempty"`
	Height         *int               `json:"height,omitempty"`
	Duration       *int               `json:"duration,omitempty"`
	Thumbnail      string             `json:"thumbnail,omitempty"`
	Description    string             `json:"description"`
	Tags           []string           `json:"tags"`
	SortOrder      int                `json:"sort_order"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

// AttachmentListDTO 附件列表DTO
type AttachmentListDTO struct {
	EntityType     *api.EntityType     `form:"entity_type"`
	EntityID       *uint               `form:"entity_id"`
	AttachmentType *api.AttachmentType `form:"attachment_type"`
	SearchRequest
}

// AttachmentStatsDTO 附件统计DTO
type AttachmentStatsDTO struct {
	TotalAttachments int64                        `json:"total_attachments"`
	TotalSize        int64                        `json:"total_size"`
	ByType           map[api.AttachmentType]int64 `json:"by_type"`
	ByEntity         map[api.EntityType]int64     `json:"by_entity"`
	StorageUsage     StorageUsageStats            `json:"storage_usage"`
}

// StorageUsageStats 存储使用统计
type StorageUsageStats struct {
	Images int64 `json:"images"`
	Videos int64 `json:"videos"`
	Files  int64 `json:"files"`
}
