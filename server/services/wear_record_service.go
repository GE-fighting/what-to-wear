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

// WearRecordService 穿着记录服务接口
type WearRecordService interface {
	// 基础CRUD操作
	CreateWearRecord(userID, itemID uint, req *dto.CreateWearRecordDTO) (*dto.WearRecordDTO, error)
	GetWearRecord(userID, recordID uint) (*dto.WearRecordDTO, error)
	GetWearRecords(userID uint, limit int) ([]dto.WearRecordDTO, error)
	UpdateWearRecord(userID, recordID uint, req *dto.UpdateWearRecordDTO) (*dto.WearRecordDTO, error)
	DeleteWearRecord(userID, recordID uint) error

	// 查询
	GetWearRecordsByItem(userID, itemID uint, limit int) ([]dto.WearRecordDTO, error)
	GetWearRecordsByDateRange(userID uint, startDate, endDate string) ([]dto.WearRecordDTO, error)
	GetWearRecordsByOccasion(userID uint, occasion string) ([]dto.WearRecordDTO, error)

	// 统计
	GetWearStats(userID uint) (*dto.WearStatsDTO, error)
	GetWearFrequency(userID uint) (map[string]int64, error)
	GetComfortAnalysis(userID uint) (*dto.ComfortAnalysisDTO, error)
}

// wearRecordService 穿着记录服务实现
type wearRecordService struct {
	wearRecordRepo repositories.WearRecordRepository
	clothingRepo   repositories.ClothingItemRepository
}

// NewWearRecordService 创建穿着记录服务实例
func NewWearRecordService(wearRecordRepo repositories.WearRecordRepository, clothingRepo repositories.ClothingItemRepository) WearRecordService {
	return &wearRecordService{
		wearRecordRepo: wearRecordRepo,
		clothingRepo:   clothingRepo,
	}
}

// CreateWearRecord 创建穿着记录
func (s *wearRecordService) CreateWearRecord(userID, itemID uint, req *dto.CreateWearRecordDTO) (*dto.WearRecordDTO, error) {
	ctx := context.Background()

	// 验证衣物存在且属于用户
	item, err := s.clothingRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}
	if item.UserID != userID {
		return nil, fmt.Errorf("无权操作该衣物")
	}

	// 创建穿着记录
	wearRecord := &models.WearRecord{
		ClothingItemID: itemID,
		WearDate:       req.WearDate,
		Notes:          req.Notes,
	}

	err = s.wearRecordRepo.Create(ctx, wearRecord)
	if err != nil {
		return nil, fmt.Errorf("创建穿着记录失败: %w", err)
	}

	return s.convertToDTO(wearRecord), nil
}

// GetWearRecord 获取穿着记录
func (s *wearRecordService) GetWearRecord(userID, recordID uint) (*dto.WearRecordDTO, error) {
	ctx := context.Background()

	wearRecord, err := s.wearRecordRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("穿着记录不存在: %w", err)
	}

	// 验证权限
	item, err := s.clothingRepo.GetByID(ctx, wearRecord.ClothingItemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}
	if item.UserID != userID {
		return nil, fmt.Errorf("无权访问该记录")
	}

	return s.convertToDTO(wearRecord), nil
}

// GetWearRecords 获取用户的穿着记录列表
func (s *wearRecordService) GetWearRecords(userID uint, limit int) ([]dto.WearRecordDTO, error) {
	ctx := context.Background()

	wearRecords, err := s.wearRecordRepo.GetByUserID(ctx, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("获取穿着记录失败: %w", err)
	}

	return s.convertToDTOList(wearRecords), nil
}

// UpdateWearRecord 更新穿着记录
func (s *wearRecordService) UpdateWearRecord(userID, recordID uint, req *dto.UpdateWearRecordDTO) (*dto.WearRecordDTO, error) {
	ctx := context.Background()

	wearRecord, err := s.wearRecordRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("穿着记录不存在: %w", err)
	}

	// 验证权限
	item, err := s.clothingRepo.GetByID(ctx, wearRecord.ClothingItemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}
	if item.UserID != userID {
		return nil, fmt.Errorf("无权修改该记录")
	}

	// 更新字段
	if req.WearDate != nil {
		wearRecord.WearDate = *req.WearDate
	}
	if req.Notes != nil {
		wearRecord.Notes = *req.Notes
	}

	err = s.wearRecordRepo.Update(ctx, wearRecord)
	if err != nil {
		return nil, fmt.Errorf("更新穿着记录失败: %w", err)
	}

	return s.convertToDTO(wearRecord), nil
}

// DeleteWearRecord 删除穿着记录
func (s *wearRecordService) DeleteWearRecord(userID, recordID uint) error {
	ctx := context.Background()

	wearRecord, err := s.wearRecordRepo.GetByID(ctx, recordID)
	if err != nil {
		return fmt.Errorf("穿着记录不存在: %w", err)
	}

	// 验证权限
	item, err := s.clothingRepo.GetByID(ctx, wearRecord.ClothingItemID)
	if err != nil {
		return fmt.Errorf("衣物不存在: %w", err)
	}
	if item.UserID != userID {
		return fmt.Errorf("无权删除该记录")
	}

	err = s.wearRecordRepo.Delete(ctx, recordID)
	if err != nil {
		return fmt.Errorf("删除穿着记录失败: %w", err)
	}

	return nil
}

// GetWearRecordsByItem 根据衣物ID获取穿着记录
func (s *wearRecordService) GetWearRecordsByItem(userID, itemID uint, limit int) ([]dto.WearRecordDTO, error) {
	ctx := context.Background()

	// 验证衣物权限
	item, err := s.clothingRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}
	if item.UserID != userID {
		return nil, fmt.Errorf("无权访问该衣物")
	}

	wearRecords, err := s.wearRecordRepo.GetByClothingItemID(ctx, itemID, limit)
	if err != nil {
		return nil, fmt.Errorf("获取穿着记录失败: %w", err)
	}

	return s.convertToDTOList(wearRecords), nil
}

// GetWearRecordsByDateRange 根据日期范围获取穿着记录
func (s *wearRecordService) GetWearRecordsByDateRange(userID uint, startDate, endDate string) ([]dto.WearRecordDTO, error) {
	ctx := context.Background()

	wearRecords, err := s.wearRecordRepo.GetByDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("获取穿着记录失败: %w", err)
	}

	return s.convertToDTOList(wearRecords), nil
}

// GetWearRecordsByOccasion 根据场合获取穿着记录
func (s *wearRecordService) GetWearRecordsByOccasion(userID uint, occasion string) ([]dto.WearRecordDTO, error) {
	ctx := context.Background()

	wearRecords, err := s.wearRecordRepo.GetByOccasion(ctx, userID, occasion)
	if err != nil {
		return nil, fmt.Errorf("获取穿着记录失败: %w", err)
	}

	return s.convertToDTOList(wearRecords), nil
}

// GetWearStats 获取穿着统计
func (s *wearRecordService) GetWearStats(userID uint) (*dto.WearStatsDTO, error) {
	ctx := context.Background()

	// 获取穿着频率统计
	frequency, err := s.wearRecordRepo.GetWearFrequency(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取穿着频率失败: %w", err)
	}

	// 获取最近穿着记录
	recentRecords, err := s.wearRecordRepo.GetByUserID(ctx, userID, 10)
	if err != nil {
		return nil, fmt.Errorf("获取最近穿着记录失败: %w", err)
	}

	// 计算总穿着次数
	totalWears := int64(0)
	for _, count := range frequency {
		totalWears += count
	}

	// 计算平均每件穿着次数
	averagePerItem := float64(0)
	if len(recentRecords) > 0 {
		averagePerItem = float64(totalWears) / float64(len(recentRecords))
	}

	// 获取最后穿着日期
	var lastWearDate *time.Time
	if len(recentRecords) > 0 {
		lastWearDate = &recentRecords[0].WearDate
	}

	stats := &dto.WearStatsDTO{
		TotalWears:      totalWears,
		AveragePerItem:  averagePerItem,
		WearsByCategory: make(map[string]int64), // TODO: 实现分类统计
		WearsByOccasion: frequency,
		WearsByWeather:  make(map[api.WeatherType]int64), // TODO: 实现天气统计
		MostWornItems:   []dto.ClothingItemSummary{},     // TODO: 实现最常穿衣物
		RecentWears:     s.convertToDTOList(recentRecords),
		LastWearDate:    lastWearDate,
	}

	return stats, nil
}

// GetWearFrequency 获取穿着频率
func (s *wearRecordService) GetWearFrequency(userID uint) (map[string]int64, error) {
	ctx := context.Background()
	return s.wearRecordRepo.GetWearFrequency(ctx, userID)
}

// GetComfortAnalysis 获取舒适度分析
func (s *wearRecordService) GetComfortAnalysis(userID uint) (*dto.ComfortAnalysisDTO, error) {
	// 由于简化的模型中没有舒适度字段，返回空的分析
	analysis := &dto.ComfortAnalysisDTO{
		AverageComfort:         0,
		AverageStyle:           0,
		AverageAppropriateness: 0,
		ComfortByCategory:      make(map[string]float64),
		ComfortByWeather:       make(map[api.WeatherType]float64),
	}

	return analysis, nil
}

// convertToDTO 将模型转换为DTO
func (s *wearRecordService) convertToDTO(record *models.WearRecord) *dto.WearRecordDTO {
	return &dto.WearRecordDTO{
		ID:             record.ID,
		ClothingItemID: record.ClothingItemID,
		WearDate:       record.WearDate,
		Notes:          record.Notes,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
}

// convertToDTOList 将模型列表转换为DTO列表
func (s *wearRecordService) convertToDTOList(records []models.WearRecord) []dto.WearRecordDTO {
	dtos := make([]dto.WearRecordDTO, len(records))
	for i, record := range records {
		dtos[i] = *s.convertToDTO(&record)
	}
	return dtos
}
