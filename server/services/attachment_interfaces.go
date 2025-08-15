package services

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

type AttachmentServiceInterface interface {
	// 上传单个附件
	UploadAttachment(req *dto.UploadAttachmentRequest) (*dto.AttachmentResponse, error)

	// 批量上传附件
	BatchUploadAttachments(req *dto.BatchUploadRequest) (*dto.BatchUploadResponse, error)

	// 根据实体获取附件列表
	GetAttachmentsByEntity(entityType models.EntityType, entityID uint) ([]dto.AttachmentResponse, error)

	// 根据用户获取附件列表
	GetAttachmentsByUser(userID uint, limit, offset int) (*dto.AttachmentListResponse, error)

	// 根据类型获取附件列表
	GetAttachmentsByType(attachmentType models.AttachmentType, limit, offset int) (*dto.AttachmentListResponse, error)

	// 获取单个附件信息
	GetAttachment(id uint) (*dto.AttachmentResponse, error)

	// 获取附件详细信息
	GetAttachmentInfo(attachmentID uint, userID uint) (*dto.AttachmentInfoResponse, error)

	// 删除附件
	DeleteAttachment(id uint, userID uint) error

	// 批量删除附件
	BatchDeleteAttachments(attachmentIDs []uint, userID uint) (*dto.BatchUploadResponse, error)

	// 更新附件排序
	UpdateAttachmentOrder(attachmentID uint, sortOrder int, userID uint) error

	// 更新附件信息
	UpdateAttachment(id uint, userID uint, req *dto.UpdateAttachmentRequest) (*dto.AttachmentResponse, error)

	// 更新附件详细信息
	UpdateAttachmentInfo(attachmentID uint, userID uint, req *dto.UpdateAttachmentInfoRequest) (*dto.AttachmentInfoResponse, error)

	// 获取附件统计信息
	GetAttachmentStats(userID uint) (*dto.AttachmentStatsResponse, error)
}
