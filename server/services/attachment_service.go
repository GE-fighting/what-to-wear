package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

type AttachmentService struct {
	attachmentRepo repositories.AttachmentRepository
}

func NewAttachmentService(attachmentRepo repositories.AttachmentRepository) AttachmentServiceInterface {
	return &AttachmentService{
		attachmentRepo: attachmentRepo,
	}
}

func (s *AttachmentService) UploadAttachment(req *dto.UploadAttachmentRequest) (*dto.AttachmentResponse, error) {
	// 验证文件类型
	if !s.isValidFileType(req.File) {
		return nil, fmt.Errorf("不支持的文件类型")
	}

	// 生成文件名和路径
	fileName := s.generateFileName(req.File.Filename)
	filePath := s.generateFilePath(req.EntityType, fileName)

	// 这里应该实现实际的文件上传逻辑到对象存储服务
	// 目前作为示例，使用本地存储路径
	publicURL := fmt.Sprintf("/uploads/%s/%s", req.EntityType, fileName)

	// 获取文件信息
	fileSize := req.File.Size
	mimeType := req.File.Header.Get("Content-Type")
	extension := strings.ToLower(filepath.Ext(req.File.Filename))

	// 确定附件类型
	attachmentType := s.determineAttachmentType(mimeType)

	// 创建附件记录
	attachment := &models.Attachment{
		OriginalName:    req.File.Filename,
		FileName:        fileName,
		FilePath:        filePath,
		FileSize:        fileSize,
		MimeType:        mimeType,
		Extension:       extension,
		AttachmentType:  attachmentType,
		EntityType:      req.EntityType,
		EntityID:        req.EntityID,
		UserID:          req.UserID,
		StorageProvider: "local", // 默认本地存储
		PublicURL:       publicURL,
		Description:     req.Description,
		Tags:            req.Tags,
		IsPublic:        req.IsPublic,
		SortOrder:       req.SortOrder,
	}

	// 保存到数据库
	if err := s.attachmentRepo.Create(attachment); err != nil {
		return nil, fmt.Errorf("保存附件记录失败: %v", err)
	}

	// 转换为响应DTO
	return s.convertToAttachmentResponse(attachment), nil
}

func (s *AttachmentService) BatchUploadAttachments(req *dto.BatchUploadRequest) (*dto.BatchUploadResponse, error) {
	var attachments []dto.AttachmentResponse
	var errors []string
	successCount := 0

	for _, file := range req.Files {
		uploadReq := &dto.UploadAttachmentRequest{
			File:        file,
			EntityType:  req.EntityType,
			EntityID:    req.EntityID,
			UserID:      req.UserID,
			Description: req.Description,
			Tags:        req.Tags,
			IsPublic:    req.IsPublic,
		}

		attachment, err := s.UploadAttachment(uploadReq)
		if err != nil {
			errors = append(errors, fmt.Sprintf("文件 %s 上传失败: %v", file.Filename, err))
		} else {
			attachments = append(attachments, *attachment)
			successCount++
		}
	}

	return &dto.BatchUploadResponse{
		SuccessCount: successCount,
		FailureCount: len(req.Files) - successCount,
		Attachments:  attachments,
		Errors:       errors,
	}, nil
}

func (s *AttachmentService) GetAttachmentsByEntity(entityType models.EntityType, entityID uint) ([]dto.AttachmentResponse, error) {
	attachments, err := s.attachmentRepo.GetByEntityID(entityType, entityID)
	if err != nil {
		return nil, fmt.Errorf("获取附件列表失败: %v", err)
	}

	var responses []dto.AttachmentResponse
	for _, attachment := range attachments {
		responses = append(responses, *s.convertToAttachmentResponse(&attachment))
	}

	return responses, nil
}

func (s *AttachmentService) GetAttachmentsByUser(userID uint, limit, offset int) (*dto.AttachmentListResponse, error) {
	attachments, err := s.attachmentRepo.GetByUserID(userID, limit)
	if err != nil {
		return nil, fmt.Errorf("获取用户附件失败: %v", err)
	}

	var responses []dto.AttachmentResponse
	for _, attachment := range attachments {
		responses = append(responses, *s.convertToAttachmentResponse(&attachment))
	}

	return &dto.AttachmentListResponse{
		Attachments: responses,
		Total:       len(responses),
		Page:        offset/limit + 1,
		PageSize:    limit,
	}, nil
}

func (s *AttachmentService) GetAttachmentsByType(attachmentType models.AttachmentType, limit, offset int) (*dto.AttachmentListResponse, error) {
	attachments, err := s.attachmentRepo.GetByType(attachmentType, limit)
	if err != nil {
		return nil, fmt.Errorf("获取附件列表失败: %v", err)
	}

	var responses []dto.AttachmentResponse
	for _, attachment := range attachments {
		responses = append(responses, *s.convertToAttachmentResponse(&attachment))
	}

	return &dto.AttachmentListResponse{
		Attachments: responses,
		Total:       len(responses),
		Page:        offset/limit + 1,
		PageSize:    limit,
	}, nil
}

func (s *AttachmentService) GetAttachment(id uint) (*dto.AttachmentResponse, error) {
	attachment, err := s.attachmentRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取附件信息失败: %v", err)
	}

	return s.convertToAttachmentResponse(attachment), nil
}

func (s *AttachmentService) DeleteAttachment(id uint, userID uint) error {
	// 获取附件信息
	attachment, err := s.attachmentRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("附件不存在: %v", err)
	}

	// 检查权限（只有上传者可以删除）
	if attachment.UserID != userID {
		return fmt.Errorf("没有权限删除此附件")
	}

	// 软删除附件记录
	if err := s.attachmentRepo.Delete(id); err != nil {
		return fmt.Errorf("删除附件失败: %v", err)
	}

	// 这里应该实现从对象存储服务删除实际文件的逻辑
	// TODO: 实现文件删除逻辑

	return nil
}

func (s *AttachmentService) UpdateAttachmentOrder(attachmentID uint, sortOrder int, userID uint) error {
	// 获取附件信息
	attachment, err := s.attachmentRepo.GetByID(attachmentID)
	if err != nil {
		return fmt.Errorf("附件不存在: %v", err)
	}

	// 检查权限
	if attachment.UserID != userID {
		return fmt.Errorf("没有权限修改此附件")
	}

	// 更新排序字段
	attachment.SortOrder = sortOrder
	return s.attachmentRepo.Update(attachment)
}

func (s *AttachmentService) UpdateAttachment(id uint, userID uint, req *dto.UpdateAttachmentRequest) (*dto.AttachmentResponse, error) {
	attachment, err := s.attachmentRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("附件不存在: %v", err)
	}

	// 检查权限
	if attachment.UserID != userID {
		return nil, fmt.Errorf("没有权限修改此附件")
	}

	// 更新字段
	if req.Description != nil {
		attachment.Description = *req.Description
	}
	if req.Tags != nil {
		attachment.Tags = req.Tags
	}
	if req.IsPublic != nil {
		attachment.IsPublic = *req.IsPublic
	}
	if req.SortOrder != nil {
		attachment.SortOrder = *req.SortOrder
	}

	if err := s.attachmentRepo.Update(attachment); err != nil {
		return nil, fmt.Errorf("更新附件失败: %v", err)
	}

	return s.convertToAttachmentResponse(attachment), nil
}

// GetAttachmentInfo 获取附件详细信息
func (s *AttachmentService) GetAttachmentInfo(attachmentID uint, userID uint) (*dto.AttachmentInfoResponse, error) {
	attachment, err := s.attachmentRepo.GetByID(attachmentID)
	if err != nil {
		return nil, fmt.Errorf("附件不存在: %v", err)
	}

	// 检查权限
	if attachment.UserID != userID {
		return nil, fmt.Errorf("没有权限访问此附件")
	}

	return &dto.AttachmentInfoResponse{
		AttachmentResponse: *s.convertToAttachmentResponse(attachment),
		Metadata:           attachment.Metadata,
		UploadedBy:         fmt.Sprintf("用户%d", attachment.UserID),
		StorageInfo: dto.StorageInfo{
			Provider:   attachment.StorageProvider,
			BucketName: attachment.BucketName,
			ObjectKey:  attachment.ObjectKey,
			PrivateURL: attachment.PrivateURL,
		},
	}, nil
}

// UpdateAttachmentInfo 更新附件详细信息
func (s *AttachmentService) UpdateAttachmentInfo(attachmentID uint, userID uint, req *dto.UpdateAttachmentInfoRequest) (*dto.AttachmentInfoResponse, error) {
	attachment, err := s.attachmentRepo.GetByID(attachmentID)
	if err != nil {
		return nil, fmt.Errorf("附件不存在: %v", err)
	}

	// 检查权限
	if attachment.UserID != userID {
		return nil, fmt.Errorf("没有权限修改此附件")
	}

	// 更新字段
	if req.FileName != nil {
		attachment.OriginalName = *req.FileName
	}
	if req.Description != nil {
		attachment.Description = *req.Description
	}
	if req.Tags != nil {
		attachment.Tags = req.Tags
	}

	if err := s.attachmentRepo.Update(attachment); err != nil {
		return nil, fmt.Errorf("更新附件失败: %v", err)
	}

	return s.GetAttachmentInfo(attachmentID, userID)
}

// BatchDeleteAttachments 批量删除附件
func (s *AttachmentService) BatchDeleteAttachments(attachmentIDs []uint, userID uint) (*dto.BatchUploadResponse, error) {
	var errors []string
	successCount := 0

	for _, id := range attachmentIDs {
		if err := s.DeleteAttachment(id, userID); err != nil {
			errors = append(errors, fmt.Sprintf("删除附件 %d 失败: %v", id, err))
		} else {
			successCount++
		}
	}

	return &dto.BatchUploadResponse{
		SuccessCount: successCount,
		FailureCount: len(attachmentIDs) - successCount,
		Attachments:  []dto.AttachmentResponse{}, // 删除操作不返回附件列表
		Errors:       errors,
	}, nil
}

// GetAttachmentStats 获取附件统计信息
func (s *AttachmentService) GetAttachmentStats(userID uint) (*dto.AttachmentStatsResponse, error) {
	// 这里应该实现实际的统计逻辑
	// 目前返回基础统计信息
	attachments, err := s.attachmentRepo.GetByUserID(userID, 1000) // 获取用户所有附件
	if err != nil {
		return nil, fmt.Errorf("获取附件统计失败: %v", err)
	}

	stats := &dto.AttachmentStatsResponse{
		TotalAttachments: int64(len(attachments)),
		TotalSize:        0,
		ByType:           make(map[string]int64),
		ByEntity:         make(map[string]int64),
		StorageUsage: dto.StorageUsageStats{
			Images: 0,
			Videos: 0,
			Files:  0,
		},
	}

	// 计算统计信息
	for _, attachment := range attachments {
		stats.TotalSize += attachment.FileSize
		stats.ByType[string(attachment.AttachmentType)]++
		stats.ByEntity[string(attachment.EntityType)]++

		switch attachment.AttachmentType {
		case models.AttachmentTypeImage:
			stats.StorageUsage.Images += attachment.FileSize
		case models.AttachmentTypeVideo:
			stats.StorageUsage.Videos += attachment.FileSize
		default:
			stats.StorageUsage.Files += attachment.FileSize
		}
	}

	return stats, nil
}

// 辅助方法
func (s *AttachmentService) generateFileName(originalName string) string {
	// 生成唯一文件名，避免重复
	ext := filepath.Ext(originalName)
	baseName := strings.TrimSuffix(originalName, ext)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d%s", baseName, timestamp, ext)
}

func (s *AttachmentService) generateFilePath(entityType models.EntityType, fileName string) string {
	return fmt.Sprintf("uploads/%s/%s", entityType, fileName)
}

func (s *AttachmentService) isValidFileType(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"video/mp4":  true,
		"video/avi":  true,
		"video/mov":  true,
	}

	mimeType := file.Header.Get("Content-Type")
	return allowedTypes[mimeType]
}

func (s *AttachmentService) determineAttachmentType(mimeType string) models.AttachmentType {
	if strings.HasPrefix(mimeType, "image/") {
		return models.AttachmentTypeImage
	}
	if strings.HasPrefix(mimeType, "video/") {
		return models.AttachmentTypeVideo
	}
	return models.AttachmentTypeFile
}

func (s *AttachmentService) convertToAttachmentResponse(attachment *models.Attachment) *dto.AttachmentResponse {
	return &dto.AttachmentResponse{
		ID:             attachment.ID,
		OriginalName:   attachment.OriginalName,
		FileName:       attachment.FileName,
		FileSize:       attachment.FileSize,
		MimeType:       attachment.MimeType,
		AttachmentType: string(attachment.AttachmentType),
		EntityType:     string(attachment.EntityType),
		EntityID:       attachment.EntityID,
		PublicURL:      attachment.GetURL(),
		Width:          attachment.Width,
		Height:         attachment.Height,
		Duration:       attachment.Duration,
		Thumbnail: func() string {
			if attachment.Thumbnail != nil {
				return *attachment.Thumbnail
			}
			return ""
		}(),
		Description: attachment.Description,
		Tags:        attachment.Tags,
		SortOrder:   attachment.SortOrder,
		CreatedAt:   attachment.CreatedAt,
		UpdatedAt:   attachment.UpdatedAt,
	}
}
