package repositories

import (
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// attachmentRepository 附件仓库实现
type attachmentRepository struct {
	db *gorm.DB
}

// NewAttachmentRepository 创建附件仓库实例
func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

// Create 创建附件
func (r *attachmentRepository) Create(attachment *models.Attachment) error {
	return r.db.Create(attachment).Error
}

// GetByID 根据ID获取附件
func (r *attachmentRepository) GetByID(id uint) (*models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.First(&attachment, id).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

// Update 更新附件
func (r *attachmentRepository) Update(attachment *models.Attachment) error {
	return r.db.Save(attachment).Error
}

// Delete 删除附件
func (r *attachmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Attachment{}, id).Error
}

// GetByEntityID 根据关联实体获取附件
func (r *attachmentRepository) GetByEntityID(entityType models.EntityType, entityID uint) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.Where("entity_type = ? AND entity_id = ? AND is_active = ?", entityType, entityID, true).
		Order("sort_order ASC, created_at ASC").
		Find(&attachments).Error
	return attachments, err
}

// GetByUserID 根据用户ID获取附件
func (r *attachmentRepository) GetByUserID(userID uint, limit int) ([]models.Attachment, error) {
	var attachments []models.Attachment
	query := r.db.Where("user_id = ? AND is_active = ?", userID, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&attachments).Error
	return attachments, err
}

// GetByType 根据附件类型获取附件
func (r *attachmentRepository) GetByType(attachmentType models.AttachmentType, limit int) ([]models.Attachment, error) {
	var attachments []models.Attachment
	query := r.db.Where("attachment_type = ? AND is_active = ?", attachmentType, true).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&attachments).Error
	return attachments, err
}

// GetTotalSize 获取用户附件总大小
func (r *attachmentRepository) GetTotalSize(userID uint) (int64, error) {
	var totalSize int64
	err := r.db.Model(&models.Attachment{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Select("COALESCE(SUM(file_size), 0)").
		Scan(&totalSize).Error
	return totalSize, err
}

// GetCountByType 获取用户各类型附件数量统计
func (r *attachmentRepository) GetCountByType(userID uint) (map[models.AttachmentType]int64, error) {
	var results []struct {
		AttachmentType models.AttachmentType `json:"attachment_type"`
		Count          int64                 `json:"count"`
	}

	err := r.db.Model(&models.Attachment{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Select("attachment_type, COUNT(*) as count").
		Group("attachment_type").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	countMap := make(map[models.AttachmentType]int64)
	for _, result := range results {
		countMap[result.AttachmentType] = result.Count
	}

	return countMap, nil
}
