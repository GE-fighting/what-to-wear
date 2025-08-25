package services

import (
	"context"
	"fmt"
	"time"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

// MaintenanceService 保养服务接口
type MaintenanceService interface {
	// 基础CRUD操作
	CreateMaintenanceRecord(ctx context.Context, itemID uint, req *dto.CreateMaintenanceRecordDTO) (*dto.MaintenanceRecordDTO, error)
	GetMaintenanceRecord(ctx context.Context, userID, recordID uint) (*dto.MaintenanceRecordDTO, error)
	GetMaintenanceRecords(ctx context.Context, userID uint, limit int) ([]dto.MaintenanceRecordDTO, error)
	UpdateMaintenanceRecord(ctx context.Context, userID, recordID uint, req *dto.UpdateMaintenanceRecordDTO) (*dto.MaintenanceRecordDTO, error)
	DeleteMaintenanceRecord(ctx context.Context, userID, recordID uint) error

	// 提醒功能
	GetUpcomingMaintenance(ctx context.Context, userID uint, days int) ([]dto.MaintenanceReminderDTO, error)
	GetOverdueMaintenance(ctx context.Context, userID uint) ([]dto.MaintenanceReminderDTO, error)
	MarkReminderSent(ctx context.Context, recordID uint) error

	// 统计
	GetMaintenanceCostByType(ctx context.Context, userID uint) (map[string]float64, error)
}

// maintenanceService 保养服务实现
type maintenanceService struct {
	maintenanceRepo repositories.MaintenanceRecordRepository
	clothingRepo   repositories.ClothingItemRepository
}

// NewMaintenanceService 创建保养服务实例
func NewMaintenanceService(
	maintenanceRepo repositories.MaintenanceRecordRepository,
	clothingRepo repositories.ClothingItemRepository,
) MaintenanceService {
	return &maintenanceService{
		maintenanceRepo: maintenanceRepo,
		clothingRepo:   clothingRepo,
	}
}

// CreateMaintenanceRecord 创建保养记录
func (s *maintenanceService) CreateMaintenanceRecord(ctx context.Context, itemID uint, req *dto.CreateMaintenanceRecordDTO) (*dto.MaintenanceRecordDTO, error) {
	// 验证保养类型
	if !api.IsValidMaintenanceType(req.MaintenanceType) {
		return nil, fmt.Errorf("invalid maintenance type: %s", req.MaintenanceType)
	}

	// 创建保养记录模型
	record := &models.MaintenanceRecord{
		ClothingItemID:  itemID,
		MaintenanceType: api.MaintenanceType(req.MaintenanceType),
		Cost:           req.Cost,
		MaintenanceDate: req.MaintenanceDate,
		ServiceProvider: req.ServiceProvider,
		Notes:          req.Notes,
	}

	// 如果指定了下一次保养日期，使用它
	if req.NextMaintenanceDate != nil {
		record.NextMaintenanceDate = req.NextMaintenanceDate
	}

	// 创建记录
	err := s.maintenanceRepo.Create(ctx, record)
	if err != nil {
		return nil, fmt.Errorf("failed to create maintenance record: %w", err)
	}

	// 转换为 DTO 返回
	return s.convertToMaintenanceRecordDTO(record), nil
}

// GetMaintenanceRecord 获取保养记录
func (s *maintenanceService) GetMaintenanceRecord(ctx context.Context, userID, recordID uint) (*dto.MaintenanceRecordDTO, error) {
	// 获取记录
	record, err := s.maintenanceRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("failed to get maintenance record: %w", err)
	}

	// 验证记录属于该用户
	if err := s.validateRecordOwnership(ctx, userID, record.ClothingItemID); err != nil {
		return nil, err
	}

	// 转换为 DTO 返回
	return s.convertToMaintenanceRecordDTO(record), nil
}

// GetMaintenanceRecords 获取用户的保养记录列表
func (s *maintenanceService) GetMaintenanceRecords(ctx context.Context, userID uint, limit int) ([]dto.MaintenanceRecordDTO, error) {
	// 获取记录
	records, err := s.maintenanceRepo.GetByUserID(ctx, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get maintenance records: %w", err)
	}

	// 转换为 DTO 返回
	return s.convertToMaintenanceRecordDTOs(records), nil
}

// UpdateMaintenanceRecord 更新保养记录
func (s *maintenanceService) UpdateMaintenanceRecord(ctx context.Context, userID, recordID uint, req *dto.UpdateMaintenanceRecordDTO) (*dto.MaintenanceRecordDTO, error) {
	// 获取现有记录
	record, err := s.maintenanceRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("failed to get maintenance record: %w", err)
	}

	// 验证记录属于该用户
	if err := s.validateRecordOwnership(ctx, userID, record.ClothingItemID); err != nil {
		return nil, err
	}

	// 更新字段
	if req.MaintenanceType != nil {
		if !api.IsValidMaintenanceType(*req.MaintenanceType) {
			return nil, fmt.Errorf("invalid maintenance type: %s", *req.MaintenanceType)
		}
		record.MaintenanceType = api.MaintenanceType(*req.MaintenanceType)
	}
	if req.Cost != nil {
		record.Cost = *req.Cost
	}
	if req.MaintenanceDate != nil {
		record.MaintenanceDate = *req.MaintenanceDate
	}
	if req.ServiceProvider != nil {
		record.ServiceProvider = *req.ServiceProvider
	}
	if req.Notes != nil {
		record.Notes = *req.Notes
	}
	if req.NextMaintenanceDate != nil {
		record.NextMaintenanceDate = req.NextMaintenanceDate
	}

	// 更新记录
	err = s.maintenanceRepo.Update(ctx, record)
	if err != nil {
		return nil, fmt.Errorf("failed to update maintenance record: %w", err)
	}

	// 转换为 DTO 返回
	return s.convertToMaintenanceRecordDTO(record), nil
}

// DeleteMaintenanceRecord 删除保养记录
func (s *maintenanceService) DeleteMaintenanceRecord(ctx context.Context, userID, recordID uint) error {
	// 获取记录
	record, err := s.maintenanceRepo.GetByID(ctx, recordID)
	if err != nil {
		return fmt.Errorf("failed to get maintenance record: %w", err)
	}

	// 验证记录属于该用户
	if err := s.validateRecordOwnership(ctx, userID, record.ClothingItemID); err != nil {
		return err
	}

	// 删除记录
	err = s.maintenanceRepo.Delete(ctx, recordID)
	if err != nil {
		return fmt.Errorf("failed to delete maintenance record: %w", err)
	}

	return nil
}

// GetUpcomingMaintenance 获取即将到期的保养提醒
func (s *maintenanceService) GetUpcomingMaintenance(ctx context.Context, userID uint, days int) ([]dto.MaintenanceReminderDTO, error) {
	// 获取即将到期的保养记录
	records, err := s.maintenanceRepo.GetUpcoming(ctx, userID, days)
	if err != nil {
		return nil, fmt.Errorf("failed to get upcoming maintenance: %w", err)
	}

	// 转换为提醒 DTO
	return s.convertToMaintenanceReminderDTOs(records), nil
}

// GetOverdueMaintenance 获取过期的保养记录
func (s *maintenanceService) GetOverdueMaintenance(ctx context.Context, userID uint) ([]dto.MaintenanceReminderDTO, error) {
	// 获取过期的保养记录
	records, err := s.maintenanceRepo.GetOverdue(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get overdue maintenance: %w", err)
	}

	// 转换为提醒 DTO
	return s.convertToMaintenanceReminderDTOs(records), nil
}

// MarkReminderSent 标记提醒已发送
func (s *maintenanceService) MarkReminderSent(ctx context.Context, recordID uint) error {
	err := s.maintenanceRepo.MarkReminderSent(ctx, recordID)
	if err != nil {
		return fmt.Errorf("failed to mark reminder sent: %w", err)
	}
	return nil
}

// GetMaintenanceCostByType 获取按类型分组的保养费用
func (s *maintenanceService) GetMaintenanceCostByType(ctx context.Context, userID uint) (map[string]float64, error) {
	costMap, err := s.maintenanceRepo.GetMaintenanceCostByType(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get maintenance cost by type: %w", err)
	}
	return costMap, nil
}

// validateRecordOwnership 验证记录所有权
func (s *maintenanceService) validateRecordOwnership(ctx context.Context, userID, itemID uint) error {
	// 获取衣物项目
	item, err := s.clothingRepo.GetByID(ctx, itemID)
	if err != nil {
		return fmt.Errorf("failed to get clothing item: %w", err)
	}

	// 验证衣物项目属于该用户
	if item.UserID != userID {
		return fmt.Errorf("unauthorized: record does not belong to user")
	}

	return nil
}

// convertToMaintenanceRecordDTO 将模型转换为 DTO
func (s *maintenanceService) convertToMaintenanceRecordDTO(record *models.MaintenanceRecord) *dto.MaintenanceRecordDTO {
	return &dto.MaintenanceRecordDTO{
		ID:                  record.ID,
		ClothingItemID:      record.ClothingItemID,
		MaintenanceType:     string(record.MaintenanceType),
		Cost:                record.Cost,
		MaintenanceDate:     record.MaintenanceDate,
		ServiceProvider:     record.ServiceProvider,
		Description:         record.Notes, // 注意：这里使用 Notes 字段
		NextMaintenanceDate: record.NextMaintenanceDate,
		Notes:               record.Notes,
		CreatedAt:           record.CreatedAt,
		UpdatedAt:           record.UpdatedAt,
	}
}

// convertToMaintenanceRecordDTOs 将模型切片转换为 DTO 切片
func (s *maintenanceService) convertToMaintenanceRecordDTOs(records []models.MaintenanceRecord) []dto.MaintenanceRecordDTO {
	dtos := make([]dto.MaintenanceRecordDTO, len(records))
	for i, record := range records {
		dtos[i] = *s.convertToMaintenanceRecordDTO(&record)
	}
	return dtos
}

// convertToMaintenanceReminderDTO 将模型转换为提醒 DTO
func (s *maintenanceService) convertToMaintenanceReminderDTO(record *models.MaintenanceRecord) *dto.MaintenanceReminderDTO {
	now := time.Now()
	daysOverdue := 0
	priority := "low"
	itemName := "Unknown Item"

	// 尝试获取衣物项目名称
	if item, err := s.clothingRepo.GetByID(context.Background(), record.ClothingItemID); err == nil {
		itemName = item.Name
	}

	// 处理下次保养日期
	var nextMaintenanceDate time.Time
	if record.NextMaintenanceDate != nil {
		nextMaintenanceDate = *record.NextMaintenanceDate
		
		if record.NextMaintenanceDate.Before(now) {
			daysOverdue = int(now.Sub(*record.NextMaintenanceDate).Hours() / 24)
			if daysOverdue > 30 {
				priority = "urgent"
			} else if daysOverdue > 14 {
				priority = "high"
			} else if daysOverdue > 7 {
				priority = "medium"
			}
		} else {
			daysUntil := int(record.NextMaintenanceDate.Sub(now).Hours() / 24)
			if daysUntil <= 3 {
				priority = "high"
			} else if daysUntil <= 7 {
				priority = "medium"
			}
		}
	} else {
		// 如果没有下次保养日期，设置为当前时间
		nextMaintenanceDate = now
		priority = "low"
	}

	return &dto.MaintenanceReminderDTO{
		ID:                  record.ID,
		ClothingItemID:      record.ClothingItemID,
		ClothingItemName:    itemName,
		MaintenanceType:     string(record.MaintenanceType),
		NextMaintenanceDate: nextMaintenanceDate,
		DaysOverdue:         daysOverdue,
		Priority:            priority,
	}
}

// convertToMaintenanceReminderDTOs 将模型切片转换为提醒 DTO 切片
func (s *maintenanceService) convertToMaintenanceReminderDTOs(records []models.MaintenanceRecord) []dto.MaintenanceReminderDTO {
	dtos := make([]dto.MaintenanceReminderDTO, len(records))
	for i, record := range records {
		dtos[i] = *s.convertToMaintenanceReminderDTO(&record)
	}
	return dtos
}
