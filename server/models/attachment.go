package models

import (
	"gorm.io/gorm"
)

// AttachmentType 附件类型枚举
type AttachmentType string

const (
	AttachmentTypeImage AttachmentType = "image" // 图片
	AttachmentTypeVideo AttachmentType = "video" // 视频
	AttachmentTypeFile  AttachmentType = "file"  // 其他文件
)

// EntityType 关联实体类型枚举
type EntityType string

const (
	EntityTypeClothingItem EntityType = "clothing_item" // 衣物
	EntityTypeOutfit       EntityType = "outfit"        // 穿搭
	EntityTypeUser         EntityType = "user"          // 用户
	EntityTypeMaintenance  EntityType = "maintenance"   // 保养记录
	EntityTypeWearRecord   EntityType = "wear_record"   // 穿着记录
	EntityTypePurchase     EntityType = "purchase"      // 购买记录
)

// Attachment 附件模型
type Attachment struct {
	gorm.Model
	
	// 基础信息
	OriginalName   string         `json:"original_name" gorm:"not null"`     // 原始文件名
	FileName       string         `json:"file_name" gorm:"not null"`         // 存储文件名
	FilePath       string         `json:"file_path" gorm:"not null"`         // 文件路径
	FileSize       int64          `json:"file_size" gorm:"not null"`         // 文件大小（字节）
	MimeType       string         `json:"mime_type" gorm:"not null"`         // MIME类型
	Extension      string         `json:"extension" gorm:"not null"`         // 文件扩展名
	AttachmentType AttachmentType `json:"attachment_type" gorm:"not null"`   // 附件类型
	
	// 关联信息
	EntityType EntityType `json:"entity_type" gorm:"not null;index"`     // 关联实体类型
	EntityID   uint       `json:"entity_id" gorm:"not null;index"`       // 关联实体ID
	UserID     uint       `json:"user_id" gorm:"not null;index"`         // 上传用户ID
	
	// 存储信息
	StorageProvider string `json:"storage_provider" gorm:"default:'local'"` // 存储提供商 (local, oss, s3, etc.)
	BucketName      string `json:"bucket_name"`                             // 存储桶名称
	ObjectKey       string `json:"object_key"`                              // 对象键/路径
	PublicURL       string `json:"public_url"`                              // 公开访问URL
	PrivateURL      string `json:"private_url"`                             // 私有访问URL
	
	// 图片/视频特有属性
	Width     *int    `json:"width"`     // 图片/视频宽度
	Height    *int    `json:"height"`    // 图片/视频高度
	Duration  *int    `json:"duration"`  // 视频时长（秒）
	Thumbnail *string `json:"thumbnail"` // 缩略图URL
	
	// 元数据
	Description string            `json:"description"`                       // 描述
	Tags        []string          `json:"tags" gorm:"type:json"`             // 标签
	Metadata    map[string]string `json:"metadata" gorm:"type:json"`         // 额外元数据
	
	// 状态信息
	IsActive  bool `json:"is_active" gorm:"default:true"`  // 是否激活
	IsPublic  bool `json:"is_public" gorm:"default:false"` // 是否公开
	SortOrder int  `json:"sort_order" gorm:"default:0"`    // 排序顺序
}

// TableName 指定表名
func (Attachment) TableName() string {
	return "attachments"
}

// IsImage 检查是否为图片类型
func (a *Attachment) IsImage() bool {
	return a.AttachmentType == AttachmentTypeImage
}

// IsVideo 检查是否为视频类型
func (a *Attachment) IsVideo() bool {
	return a.AttachmentType == AttachmentTypeVideo
}

// GetURL 获取最合适的访问URL
func (a *Attachment) GetURL() string {
	if a.PublicURL != "" {
		return a.PublicURL
	}
	if a.PrivateURL != "" {
		return a.PrivateURL
	}
	return a.FilePath
}

// IsValidAttachmentType 检查附件类型是否有效
func IsValidAttachmentType(attachmentType string) bool {
	validTypes := []AttachmentType{
		AttachmentTypeImage, AttachmentTypeVideo, AttachmentTypeFile,
	}
	
	for _, validType := range validTypes {
		if AttachmentType(attachmentType) == validType {
			return true
		}
	}
	return false
}

// IsValidEntityType 检查实体类型是否有效
func IsValidEntityType(entityType string) bool {
	validTypes := []EntityType{
		EntityTypeClothingItem, EntityTypeOutfit, EntityTypeUser,
		EntityTypeMaintenance, EntityTypeWearRecord, EntityTypePurchase,
	}
	
	for _, validType := range validTypes {
		if EntityType(entityType) == validType {
			return true
		}
	}
	return false
}