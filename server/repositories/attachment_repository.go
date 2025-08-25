package repositories

import (
	"context"
	"what-to-wear/server/api"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

type AttachmentRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, attachment *models.Attachment) error
	GetByID(ctx context.Context, id uint) (*models.Attachment, error)
	Update(ctx context.Context, attachment *models.Attachment) error
	Delete(ctx context.Context, id uint) error

	// 查询操作
	GetByEntityID(ctx context.Context, entityType api.EntityType, entityID uint) ([]models.Attachment, error)
	GetByUserID(ctx context.Context, userID uint, limit int) ([]models.Attachment, error)
	GetByType(ctx context.Context, attachmentType api.AttachmentType, limit int) ([]models.Attachment, error)

	// 统计操作
	GetTotalSize(ctx context.Context, userID uint) (int64, error)
	GetCountByType(ctx context.Context, userID uint) (map[api.AttachmentType]int64, error)
}

// attachmentRepository 附件仓库实现
type attachmentRepository struct {
	db *gorm.DB
}

// NewAttachmentRepository 创建附件仓库实例
func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

// Create 创建附件
func (r *attachmentRepository) Create(ctx context.Context, attachment *models.Attachment) error {
	return r.db.WithContext(ctx).Create(attachment).Error
}

// GetByID 根据ID获取附件
func (r *attachmentRepository) GetByID(ctx context.Context, id uint) (*models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.WithContext(ctx).First(&attachment, id).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

// Update 更新附件
func (r *attachmentRepository) Update(ctx context.Context, attachment *models.Attachment) error {
	return r.db.WithContext(ctx).Save(attachment).Error
}

// Delete 删除附件
func (r *attachmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Attachment{}, id).Error
}

// GetByEntityID 根据关联实体获取附件
func (r *attachmentRepository) GetByEntityID(ctx context.Context, entityType api.EntityType, entityID uint) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ? AND is_active = ?", entityType, entityID, true).
		Order("sort_order ASC, created_at ASC").
		Find(&attachments).Error
	return attachments, err
}

// GetByUserID 根据用户ID获取附件
func (r *attachmentRepository) GetByUserID(ctx context.Context, userID uint, limit int) ([]models.Attachment, error) {
	var attachments []models.Attachment
	query := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&attachments).Error
	return attachments, err
}

// GetByType 根据附件类型获取附件
func (r *attachmentRepository) GetByType(ctx context.Context, attachmentType api.AttachmentType, limit int) ([]models.Attachment, error) {
	var attachments []models.Attachment
	query := r.db.WithContext(ctx).Where("attachment_type = ? AND is_active = ?", attachmentType, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&attachments).Error
	return attachments, err
}

// GetTotalSize 获取用户附件总大小
func (r *attachmentRepository) GetTotalSize(ctx context.Context, userID uint) (int64, error) {
	var totalSize int64
	err := r.db.WithContext(ctx).Model(&models.Attachment{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Select("COALESCE(SUM(file_size), 0)").
		Scan(&totalSize).Error
	return totalSize, err
}

// GetCountByType 获取用户各类型附件数量统计
func (r *attachmentRepository) GetCountByType(ctx context.Context, userID uint) (map[api.AttachmentType]int64, error) {
	var results []struct {
		AttachmentType api.AttachmentType `json:"attachment_type"`
		Count          int64              `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.Attachment{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Select("attachment_type, COUNT(*) as count").
		Group("attachment_type").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	countMap := make(map[api.AttachmentType]int64)
	for _, result := range results {
		countMap[result.AttachmentType] = result.Count
	}

	return countMap, nil
}
