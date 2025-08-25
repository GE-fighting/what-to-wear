package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

// PurchaseRecordService 购买记录服务接口
type PurchaseRecordService interface {
	// 基础CRUD操作
	CreatePurchaseRecord(userID, itemID uint, req *dto.CreatePurchaseRecordDTO) (*dto.PurchaseRecordDTO, error)
	GetPurchaseRecord(userID, recordID uint) (*dto.PurchaseRecordDTO, error)
	GetPurchaseRecords(userID uint, limit int) ([]dto.PurchaseRecordDTO, error)
	UpdatePurchaseRecord(userID, recordID uint, req *dto.UpdatePurchaseRecordDTO) (*dto.PurchaseRecordDTO, error)
	DeletePurchaseRecord(userID, recordID uint) error

	// 查询
	GetPurchasesByDateRange(userID uint, startDate, endDate string) ([]dto.PurchaseRecordDTO, error)
	GetPurchasesByStore(userID uint, storeName string) ([]dto.PurchaseRecordDTO, error)

	// 统计
	GetSpendingStats(userID uint) (*dto.SpendingStatsDTO, error)
	GetSpendingByMonth(userID uint, year int) (map[string]float64, error)
	GetSpendingByCategory(userID uint) (map[string]float64, error)
}

// purchaseRecordService 购买记录服务实现
type purchaseRecordService struct {
	purchaseRepo         repositories.PurchaseRecordRepository
	clothingItemRepo     repositories.ClothingItemRepository
	clothingCategoryRepo repositories.ClothingCategoryRepository
}

// NewPurchaseRecordService 创建购买记录服务实例
func NewPurchaseRecordService(
	purchaseRepo repositories.PurchaseRecordRepository,
	clothingItemRepo repositories.ClothingItemRepository,
	clothingCategoryRepo repositories.ClothingCategoryRepository,
) PurchaseRecordService {
	return &purchaseRecordService{
		purchaseRepo:         purchaseRepo,
		clothingItemRepo:     clothingItemRepo,
		clothingCategoryRepo: clothingCategoryRepo,
	}
}

// CreatePurchaseRecord 创建购买记录
func (s *purchaseRecordService) CreatePurchaseRecord(userID, itemID uint, req *dto.CreatePurchaseRecordDTO) (*dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 验证衣物是否属于该用户
	clothingItem, err := s.clothingItemRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}
	if clothingItem.UserID != userID {
		return nil, errors.New("无权限为此衣物创建购买记录")
	}

	// 检查是否已存在购买记录（一个衣物只能有一个购买记录）
	existingRecord, err := s.purchaseRepo.GetByClothingItemID(ctx, itemID)
	if err == nil && existingRecord != nil {
		return nil, errors.New("该衣物已存在购买记录")
	}

	// 创建购买记录模型
	purchaseRecord := &models.PurchaseRecord{
		ClothingItemID: itemID,
		Price:          req.PurchasePrice,
		Store:          req.Store,
		PurchaseDate:   req.PurchaseDate,
		Notes:          req.Notes,
	}

	// 如果有在线商店信息，优先使用在线商店
	if req.OnlineStore != "" {
		purchaseRecord.Store = req.OnlineStore
	}

	if err := s.purchaseRepo.Create(ctx, purchaseRecord); err != nil {
		return nil, fmt.Errorf("创建购买记录失败: %w", err)
	}

	// 转换为DTO并返回
	return s.convertToDTO(ctx, purchaseRecord), nil
}

// GetPurchaseRecord 获取单个购买记录
func (s *purchaseRecordService) GetPurchaseRecord(userID, recordID uint) (*dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 获取购买记录
	record, err := s.purchaseRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("购买记录不存在: %w", err)
	}

	// 验证权限
	clothingItem, err := s.clothingItemRepo.GetByID(ctx, record.ClothingItemID)
	if err != nil {
		return nil, fmt.Errorf("关联的衣物不存在: %w", err)
	}
	if clothingItem.UserID != userID {
		return nil, errors.New("无权限访问此购买记录")
	}

	// 转换为DTO并返回
	return s.convertToDTO(ctx, record), nil
}

// GetPurchaseRecords 获取用户的购买记录列表
func (s *purchaseRecordService) GetPurchaseRecords(userID uint, limit int) ([]dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 获取购买记录列表
	records, err := s.purchaseRepo.GetByUserID(ctx, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("获取购买记录失败: %w", err)
	}

	// 转换为DTO列表
	result := make([]dto.PurchaseRecordDTO, 0, len(records))
	for _, record := range records {
		dto := s.convertToDTO(ctx, &record)
		result = append(result, *dto)
	}

	return result, nil
}

// UpdatePurchaseRecord 更新购买记录
func (s *purchaseRecordService) UpdatePurchaseRecord(userID, recordID uint, req *dto.UpdatePurchaseRecordDTO) (*dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 获取现有记录
	record, err := s.purchaseRepo.GetByID(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("购买记录不存在: %w", err)
	}

	// 验证权限
	clothingItem, err := s.clothingItemRepo.GetByID(ctx, record.ClothingItemID)
	if err != nil {
		return nil, fmt.Errorf("关联的衣物不存在: %w", err)
	}
	if clothingItem.UserID != userID {
		return nil, errors.New("无权限修改此购买记录")
	}

	// 更新字段
	if req.PurchasePrice != nil {
		record.Price = *req.PurchasePrice
	}
	if req.PurchaseDate != nil {
		record.PurchaseDate = *req.PurchaseDate
	}
	if req.Store != nil {
		record.Store = *req.Store
	}
	if req.OnlineStore != nil && *req.OnlineStore != "" {
		record.Store = *req.OnlineStore // 优先使用在线商店
	}
	if req.Notes != nil {
		record.Notes = *req.Notes
	}

	// 保存更新
	if err := s.purchaseRepo.Update(ctx, record); err != nil {
		return nil, fmt.Errorf("更新购买记录失败: %w", err)
	}

	// 转换为DTO并返回
	return s.convertToDTO(ctx, record), nil
}

// DeletePurchaseRecord 删除购买记录
func (s *purchaseRecordService) DeletePurchaseRecord(userID, recordID uint) error {
	ctx := context.Background()

	// 获取购买记录
	record, err := s.purchaseRepo.GetByID(ctx, recordID)
	if err != nil {
		return fmt.Errorf("购买记录不存在: %w", err)
	}

	// 验证权限
	clothingItem, err := s.clothingItemRepo.GetByID(ctx, record.ClothingItemID)
	if err != nil {
		return fmt.Errorf("关联的衣物不存在: %w", err)
	}
	if clothingItem.UserID != userID {
		return errors.New("无权限删除此购买记录")
	}

	// 删除记录
	if err := s.purchaseRepo.Delete(ctx, recordID); err != nil {
		return fmt.Errorf("删除购买记录失败: %w", err)
	}

	return nil
}

// GetPurchasesByDateRange 根据日期范围获取购买记录
func (s *purchaseRecordService) GetPurchasesByDateRange(userID uint, startDate, endDate string) ([]dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 获取日期范围内的购买记录
	records, err := s.purchaseRepo.GetByDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("获取购买记录失败: %w", err)
	}

	// 转换为DTO列表
	result := make([]dto.PurchaseRecordDTO, 0, len(records))
	for _, record := range records {
		dto := s.convertToDTO(ctx, &record)
		result = append(result, *dto)
	}

	return result, nil
}

// GetPurchasesByStore 根据商店获取购买记录
func (s *purchaseRecordService) GetPurchasesByStore(userID uint, storeName string) ([]dto.PurchaseRecordDTO, error) {
	ctx := context.Background()

	// 获取指定商店的购买记录
	records, err := s.purchaseRepo.GetByStore(ctx, userID, storeName)
	if err != nil {
		return nil, fmt.Errorf("获取购买记录失败: %w", err)
	}

	// 转换为DTO列表
	result := make([]dto.PurchaseRecordDTO, 0, len(records))
	for _, record := range records {
		dto := s.convertToDTO(ctx, &record)
		result = append(result, *dto)
	}

	return result, nil
}

// GetSpendingStats 获取支出统计
func (s *purchaseRecordService) GetSpendingStats(userID uint) (*dto.SpendingStatsDTO, error) {
	ctx := context.Background()

	// 获取总支出
	totalSpent, err := s.purchaseRepo.GetTotalSpent(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取总支出失败: %w", err)
	}

	// 获取平均价格
	averagePrice, err := s.purchaseRepo.GetAverageItemPrice(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取平均价格失败: %w", err)
	}

	// 获取当年的月度支出
	currentYear := time.Now().Year()
	monthlySpending, err := s.purchaseRepo.GetSpentByMonth(ctx, userID, currentYear)
	if err != nil {
		return nil, fmt.Errorf("获取月度支出失败: %w", err)
	}

	// 获取分类支出
	categorySpending, err := s.purchaseRepo.GetSpentByCategory(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取分类支出失败: %w", err)
	}

	// 获取商店支出（作为品牌支出的替代）
	brandSpending, err := s.purchaseRepo.GetSpentByStore(ctx, userID)
	if err != nil {
		// 如果获取失败，使用空map
		brandSpending = make(map[string]float64)
	}

	// 计算每次穿着成本（简化实现）
	costPerWear := 0.0
	if totalSpent > 0 {
		// 这里应该结合穿着记录来计算，暂时简化为平均值
		costPerWear = totalSpent / 100 // 假设平均每件衣物穿着100次
	}

	// 构建统计DTO
	stats := &dto.SpendingStatsDTO{
		TotalSpent:        totalSpent,
		MonthlySpending:   monthlySpending,
		CategorySpending:  categorySpending,
		BrandSpending:     brandSpending,
		AverageItemPrice:  averagePrice,
		CostPerWear:       costPerWear,
		MostExpensiveItem: nil,                         // TODO: 实现最贵商品查询
		BestValueItems:    []dto.ClothingItemSummary{}, // TODO: 实现最佳性价比商品查询
	}

	return stats, nil
}

// GetSpendingByMonth 获取月度支出统计
func (s *purchaseRecordService) GetSpendingByMonth(userID uint, year int) (map[string]float64, error) {
	ctx := context.Background()

	monthlySpending, err := s.purchaseRepo.GetSpentByMonth(ctx, userID, year)
	if err != nil {
		return nil, fmt.Errorf("获取月度支出失败: %w", err)
	}

	return monthlySpending, nil
}

// GetSpendingByCategory 获取分类支出统计
func (s *purchaseRecordService) GetSpendingByCategory(userID uint) (map[string]float64, error) {
	ctx := context.Background()

	categorySpending, err := s.purchaseRepo.GetSpentByCategory(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取分类支出失败: %w", err)
	}

	return categorySpending, nil
}

// convertToDTO 将模型转换为DTO
func (s *purchaseRecordService) convertToDTO(ctx context.Context, record *models.PurchaseRecord) *dto.PurchaseRecordDTO {
	dto := &dto.PurchaseRecordDTO{
		ID:             record.ID,
		ClothingItemID: record.ClothingItemID,
		PurchasePrice:  record.Price,
		OriginalPrice:  record.Price, // 简化模型中没有原价，使用实际价格
		PurchaseDate:   record.PurchaseDate,
		Store:          record.Store,
		OnlineStore:    "", // 简化模型中合并了线上线下商店
		PaymentMethod:  "", // 简化模型中没有此字段
		OrderNumber:    "", // 简化模型中没有此字段
		Notes:          record.Notes,
		ReceiptURL:     "", // 简化模型中没有此字段
		WarrantyPeriod: 0,  // 简化模型中没有此字段
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}

	return dto
}
