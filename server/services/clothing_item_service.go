package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

// ClothingItemService 衣物服务接口
type ClothingItemService interface {
	// 基础CRUD操作
	CreateClothingItem(ctx context.Context, userID uint, req *dto.CreateClothingItemDTO) (*dto.ClothingItemDTO, error)
	GetClothingItem(ctx context.Context, userID, itemID uint) (*dto.ClothingItemDTO, error)
	GetClothingItems(ctx context.Context, userID uint, req *dto.ClothingItemListDTO) ([]models.ClothingItem, int, error)
	UpdateClothingItem(ctx context.Context, userID, itemID uint, req *dto.UpdateClothingItemDTO) (*dto.ClothingItemDTO, error)
	DeleteClothingItem(ctx context.Context, userID, itemID uint) error

	// 批量操作
	BatchDeleteClothingItems(ctx context.Context, userID uint, itemIDs []uint) (map[string]interface{}, error)

	// 高级功能
	GetClothingStats(ctx context.Context, userID uint) (*dto.ClothingStatsDTO, error)
	SearchClothingItems(ctx context.Context, userID uint, query string, limit int) ([]dto.ClothingItemSummary, error)
	GetRecommendations(ctx context.Context, userID uint, occasion string, weather string) ([]dto.ClothingItemSummary, error)
}

// clothingItemService 衣物服务实现
type clothingItemService struct {
	clothingItemRepo     repositories.ClothingItemRepository
	clothingCategoryRepo repositories.ClothingCategoryRepository
	attachmentRepo       repositories.AttachmentRepository
	purchaseRecordRepo   repositories.PurchaseRecordRepository
	wearRecordRepo       repositories.WearRecordRepository
}

// NewClothingItemService 创建衣物服务实例
func NewClothingItemService(
	clothingItemRepo repositories.ClothingItemRepository,
	clothingCategoryRepo repositories.ClothingCategoryRepository,
	attachmentRepo repositories.AttachmentRepository,
	purchaseRecordRepo repositories.PurchaseRecordRepository,
	wearRecordRepo repositories.WearRecordRepository,
) ClothingItemService {
	return &clothingItemService{
		clothingItemRepo:     clothingItemRepo,
		clothingCategoryRepo: clothingCategoryRepo,
		attachmentRepo:       attachmentRepo,
		purchaseRecordRepo:   purchaseRecordRepo,
		wearRecordRepo:       wearRecordRepo,
	}
}

// CreateClothingItem 创建衣物
func (s *clothingItemService) CreateClothingItem(ctx context.Context, userID uint, req *dto.CreateClothingItemDTO) (*dto.ClothingItemDTO, error) {
	// 验证分类是否存在
	category, err := s.clothingCategoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("分类不存在: %w", err)
	}

	// 创建衣物模型
	clothingItem := &models.ClothingItem{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Brand:      req.Brand,
		Color:      req.Color,
		Material:   req.Material,
		Condition:  req.Status,
		IsActive:   true,
		IsFavorite: false,
	}

	// 设置尺码
	if req.Size != "" {
		clothingItem.Size = models.ClothingSize{Size: req.Size, System: "CN"}
	}

	// 设置价格（如果有购买信息）
	if req.PurchaseInfo != nil {
		clothingItem.Price = req.PurchaseInfo.PurchasePrice
		clothingItem.PurchaseDate = &req.PurchaseInfo.PurchaseDate
	}

	// 创建衣物
	err = s.clothingItemRepo.Create(ctx, clothingItem)
	if err != nil {
		return nil, fmt.Errorf("创建衣物失败: %w", err)
	}

	// 添加标签
	if len(req.Tags) > 0 {
		err = s.clothingItemRepo.AddTags(ctx, clothingItem.ID, req.Tags)
		if err != nil {
			return nil, fmt.Errorf("添加标签失败: %w", err)
		}
	}

	return s.convertToDTO(clothingItem, category), nil
}

// GetClothingItem 获取衣物详情
func (s *clothingItemService) GetClothingItem(ctx context.Context, userID, itemID uint) (*dto.ClothingItemDTO, error) {
	// 获取衣物
	item, err := s.clothingItemRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}

	// 验证权限
	if item.UserID != userID {
		return nil, errors.New("无权访问该衣物")
	}

	// 获取分类信息
	category, err := s.clothingCategoryRepo.GetByID(ctx, item.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("获取分类信息失败: %w", err)
	}

	return s.convertToDTO(item, category), nil
}

// GetClothingItems 获取衣物列表
func (s *clothingItemService) GetClothingItems(ctx context.Context, userID uint, req *dto.ClothingItemListDTO) ([]models.ClothingItem, int, error) {
	// 获取衣物列表
	items, total, err := s.clothingItemRepo.GetByUserID(ctx, userID, req)
	if err != nil {
		return nil, 0, fmt.Errorf("获取衣物列表失败: %w", err)
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	// 构建响应（这里简化返回请求DTO，实际应该构建响应DTO）
	return items, totalPages, nil
}

// UpdateClothingItem 更新衣物
func (s *clothingItemService) UpdateClothingItem(ctx context.Context, userID, itemID uint, req *dto.UpdateClothingItemDTO) (*dto.ClothingItemDTO, error) {
	// 获取衣物
	item, err := s.clothingItemRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("衣物不存在: %w", err)
	}

	// 验证权限
	if item.UserID != userID {
		return nil, errors.New("无权修改该衣物")
	}

	// 更新字段
	if req.CategoryID != nil {
		item.CategoryID = *req.CategoryID
	}
	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Brand != nil {
		item.Brand = *req.Brand
	}
	if req.Color != nil {
		item.Color = *req.Color
	}
	if req.Material != nil {
		item.Material = *req.Material
	}
	if req.Status != nil {
		item.Condition = *req.Status
	}

	// 更新衣物
	err = s.clothingItemRepo.Update(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("更新衣物失败: %w", err)
	}

	// 获取分类信息
	category, _ := s.clothingCategoryRepo.GetByID(ctx, item.CategoryID)

	return s.convertToDTO(item, category), nil
}

// DeleteClothingItem 删除衣物
func (s *clothingItemService) DeleteClothingItem(ctx context.Context, userID, itemID uint) error {
	// 获取衣物
	item, err := s.clothingItemRepo.GetByID(ctx, itemID)
	if err != nil {
		return fmt.Errorf("衣物不存在: %w", err)
	}

	// 验证权限
	if item.UserID != userID {
		return errors.New("无权删除该衣物")
	}

	// 软删除衣物
	return s.clothingItemRepo.Delete(ctx, itemID)
}

// BatchDeleteClothingItems 批量删除衣物
func (s *clothingItemService) BatchDeleteClothingItems(ctx context.Context, userID uint, itemIDs []uint) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	successCount := 0
	failedCount := 0

	for _, itemID := range itemIDs {
		err := s.DeleteClothingItem(ctx, userID, itemID)
		if err != nil {
			failedCount++
		} else {
			successCount++
		}
	}

	result["success_count"] = successCount
	result["failed_count"] = failedCount
	result["total_count"] = len(itemIDs)

	return result, nil
}

// GetClothingStats 获取衣物统计
func (s *clothingItemService) GetClothingStats(ctx context.Context, userID uint) (*dto.ClothingStatsDTO, error) {
	// 获取基础统计
	req := &dto.ClothingItemListDTO{
		SearchRequest: dto.SearchRequest{
			PaginationRequest: dto.PaginationRequest{Page: 1, PageSize: 1},
		},
	}
	_, total, err := s.clothingItemRepo.GetByUserID(ctx, userID, req)
	if err != nil {
		return nil, fmt.Errorf("获取衣物总数失败: %w", err)
	}

	stats := &dto.ClothingStatsDTO{
		TotalItems:    total,
		ByCategory:    make(map[string]int64),
		ByStatus:      make(map[api.ClothingStatus]int64),
		BySeason:      make(map[string]int64),
		ByOccasion:    make(map[string]int64),
		ByBrand:       make(map[string]int64),
		ByColor:       make(map[string]int64),
		TotalValue:    0,
		AveragePrice:  0,
		MostWornItems: []dto.ClothingItemSummary{},
		RecentlyAdded: []dto.ClothingItemSummary{},
		LastUpdated:   time.Now(),
	}

	return stats, nil
}

// SearchClothingItems 搜索衣物
func (s *clothingItemService) SearchClothingItems(ctx context.Context, userID uint, query string, limit int) ([]dto.ClothingItemSummary, error) {
	// 搜索衣物
	items, err := s.clothingItemRepo.Search(ctx, userID, query, limit)
	if err != nil {
		return nil, fmt.Errorf("搜索衣物失败: %w", err)
	}

	return s.convertToSummaryList(items), nil
}

// GetRecommendations 获取推荐衣物
func (s *clothingItemService) GetRecommendations(ctx context.Context, userID uint, occasion string, weather string) ([]dto.ClothingItemSummary, error) {
	// 简化的推荐逻辑
	req := &dto.ClothingItemListDTO{
		SearchRequest: dto.SearchRequest{
			PaginationRequest: dto.PaginationRequest{
				Page:     1,
				PageSize: 20,
			},
		},
	}
	items, _, err := s.clothingItemRepo.GetByUserID(ctx, userID, req)
	if err != nil {
		return nil, fmt.Errorf("获取衣物列表失败: %w", err)
	}

	// 返回前10个作为推荐
	if len(items) > 10 {
		items = items[:10]
	}

	return s.convertToSummaryList(items), nil
}

// convertToDTO 将模型转换为DTO
func (s *clothingItemService) convertToDTO(item *models.ClothingItem, category *models.ClothingCategory) *dto.ClothingItemDTO {
	categoryName := ""
	if category != nil {
		categoryName = category.Name
	}

	return &dto.ClothingItemDTO{
		ID:                 item.ID,
		UserID:             item.UserID,
		CategoryID:         item.CategoryID,
		CategoryName:       categoryName,
		Name:               item.Name,
		Brand:              item.Brand,
		Color:              item.Color,
		Size:               item.Size.Size,
		Material:           item.Material,
		Season:             []string{},
		Occasion:           []string{},
		Style:              "",
		Description:        "",
		Status:             item.Condition,
		Tags:               []dto.TagDTO{},
		Attachments:        []dto.AttachmentDTO{},
		PurchaseRecord:     nil,
		MaintenanceRecords: []dto.MaintenanceRecordDTO{},
		WearRecords:        []dto.WearRecordDTO{},
		WearCount:          item.WearCount,
		LastWornDate:       item.LastWornDate,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}
}

// convertToSummaryList 将模型列表转换为摘要DTO列表
func (s *clothingItemService) convertToSummaryList(items []models.ClothingItem) []dto.ClothingItemSummary {
	summaries := make([]dto.ClothingItemSummary, len(items))

	for i, item := range items {
		summaries[i] = dto.ClothingItemSummary{
			ID:            item.ID,
			Name:          item.Name,
			Brand:         item.Brand,
			Color:         item.Color,
			CategoryName:  "", // TODO: 获取分类名称
			ImageURL:      "", // TODO: 获取图片URL
			Status:        item.Condition,
			WearCount:     item.WearCount,
			LastWornDate:  item.LastWornDate,
			PurchasePrice: &item.Price,
		}
	}

	return summaries
}
