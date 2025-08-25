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

// OutfitService 穿搭服务接口
type OutfitService interface {
	// 创建穿搭记录
	CreateOutfit(userID uint, req *dto.CreateOutfitDTO) (*dto.Outfit, error)

	// 获取用户穿搭历史
	GetUserOutfits(userID uint, page, pageSize int) ([]*dto.Outfit, int64, error)

	// 获取穿搭推荐
	GetOutfitRecommendation(userID uint, weather string) (*dto.OutfitRecommendation, error)

	// 评价穿搭
	RateOutfit(userID, outfitID uint, rating int, notes string) error
}

// outfitService 穿搭服务实现
type outfitService struct {
	outfitRepo           repositories.OutfitRepository
	outfitItemRepo       repositories.OutfitItemRepository
	clothingItemRepo     repositories.ClothingItemRepository
	clothingCategoryRepo repositories.ClothingCategoryRepository
	attachmentRepo       repositories.AttachmentRepository
}

// NewOutfitService 创建穿搭服务实例
func NewOutfitService(
	outfitRepo repositories.OutfitRepository,
	outfitItemRepo repositories.OutfitItemRepository,
	clothingItemRepo repositories.ClothingItemRepository,
	clothingCategoryRepo repositories.ClothingCategoryRepository,
	attachmentRepo repositories.AttachmentRepository,
) OutfitService {
	return &outfitService{
		outfitRepo:           outfitRepo,
		outfitItemRepo:       outfitItemRepo,
		clothingItemRepo:     clothingItemRepo,
		clothingCategoryRepo: clothingCategoryRepo,
		attachmentRepo:       attachmentRepo,
	}
}

// CreateOutfit 创建穿搭记录
func (s *outfitService) CreateOutfit(userID uint, req *dto.CreateOutfitDTO) (*dto.Outfit, error) {
	ctx := context.Background()

	// 验证衣物ID是否属于该用户
	for _, clothingID := range req.ClothingIDs {
		item, err := s.clothingItemRepo.GetByID(ctx, clothingID)
		if err != nil {
			return nil, fmt.Errorf("衣物ID %d 不存在", clothingID)
		}
		if item.UserID != userID {
			return nil, fmt.Errorf("衣物ID %d 不属于当前用户", clothingID)
		}
	}

	// 创建穿搭记录
	outfit := &models.Outfit{
		UserID:      userID,
		Name:        req.Name,
		Date:        req.Date,
		Temperature: req.Temperature,
		Weather:     req.Weather,
		Occasion:    req.Occasion,
		Location:    req.Location,
		Notes:       req.Notes,
		Tags:        req.Tags,
		IsPublic:    req.IsPublic,
	}

	if err := s.outfitRepo.Create(ctx, outfit); err != nil {
		return nil, fmt.Errorf("创建穿搭记录失败: %w", err)
	}

	// 创建穿搭单品关联记录
	outfitItems := make([]models.OutfitItem, 0, len(req.ClothingIDs))
	for i, clothingID := range req.ClothingIDs {
		outfitItem := models.OutfitItem{
			OutfitID:       outfit.ID,
			ClothingItemID: clothingID,
			LayerOrder:     i + 1,                    // 从1开始
			ItemRole:       string(api.ItemRoleMain), // 默认为主要单品
		}
		outfitItems = append(outfitItems, outfitItem)
	}

	if err := s.outfitItemRepo.CreateBatch(ctx, outfitItems); err != nil {
		return nil, fmt.Errorf("创建穿搭单品关联失败: %w", err)
	}

	// 转换为DTO并返回
	return s.convertToOutfitDTO(ctx, outfit)
}

// GetUserOutfits 获取用户穿搭历史
func (s *outfitService) GetUserOutfits(userID uint, page, pageSize int) ([]*dto.Outfit, int64, error) {
	ctx := context.Background()

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 获取穿搭记录
	outfits, err := s.outfitRepo.GetByUserID(ctx, userID, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户穿搭历史失败: %w", err)
	}

	// 转换为DTO
	outfitDTOs := make([]*dto.Outfit, 0, len(outfits))
	for _, outfit := range outfits {
		outfitDTO, err := s.convertToOutfitDTO(ctx, outfit)
		if err != nil {
			return nil, 0, fmt.Errorf("转换穿搭数据失败: %w", err)
		}
		outfitDTOs = append(outfitDTOs, outfitDTO)
	}

	// 获取总数 (这里简化实现，实际项目中应该在repository中实现Count方法)
	allOutfits, err := s.outfitRepo.GetByUserID(ctx, userID, 0, 0)
	if err != nil {
		return nil, 0, fmt.Errorf("获取总数失败: %w", err)
	}
	total := int64(len(allOutfits))

	return outfitDTOs, total, nil
}

// GetOutfitRecommendation 获取穿搭推荐
func (s *outfitService) GetOutfitRecommendation(userID uint, weather string) (*dto.OutfitRecommendation, error) {
	ctx := context.Background()

	// 解析天气类型
	weatherType := api.WeatherType(weather)
	if !weatherType.IsValid() {
		return nil, errors.New("无效的天气类型")
	}

	// 获取用户所有衣物
	clothingListReq := &dto.ClothingItemListDTO{
		SearchRequest: dto.SearchRequest{
			PaginationRequest: dto.PaginationRequest{
				Page:     1,
				PageSize: 100, // 暂时限制为100，实际项目中可以分批处理
			},
		},
	}

	clothingItems, _, err := s.clothingItemRepo.GetByUserID(ctx, userID, clothingListReq)
	if err != nil {
		return nil, fmt.Errorf("获取用户衣物失败: %w", err)
	}

	if len(clothingItems) == 0 {
		return nil, errors.New("用户暂无衣物，无法生成推荐")
	}

	// 基于天气和季节进行简单推荐逻辑
	recommendedItems := s.generateRecommendations(ctx, clothingItems, weatherType)

	// 构建推荐结果
	recommendation := &dto.OutfitRecommendation{
		ID:               0, // 推荐不保存到数据库，所以ID为0
		RecommendedItems: recommendedItems,
		Weather: dto.WeatherInfo{
			Temperature: s.getTemperatureByWeather(weatherType),
			Condition:   weatherType,
			Description: s.getWeatherDescription(weatherType),
		},
		Occasion:   "日常", // 默认场合
		Confidence: s.calculateConfidence(recommendedItems),
		Reason:     s.generateRecommendationReason(weatherType, recommendedItems),
		CreatedAt:  time.Now(),
	}

	return recommendation, nil
}

// RateOutfit 评价穿搭
func (s *outfitService) RateOutfit(userID, outfitID uint, rating int, notes string) error {
	ctx := context.Background()

	// 验证评分范围
	outfitRating := api.OutfitRating(rating)
	if !outfitRating.IsValid() {
		return errors.New("评分必须在1-5之间")
	}

	// 获取穿搭记录
	outfit, err := s.outfitRepo.GetByID(ctx, outfitID)
	if err != nil {
		return fmt.Errorf("获取穿搭记录失败: %w", err)
	}

	// 验证穿搭属于该用户
	if outfit.UserID != userID {
		return errors.New("无权限修改此穿搭记录")
	}

	// 更新评分
	outfit.Rating = &outfitRating
	outfit.RatingNotes = notes

	if err := s.outfitRepo.Update(ctx, outfit); err != nil {
		return fmt.Errorf("更新穿搭评分失败: %w", err)
	}

	return nil
}

// convertToOutfitDTO 将模型转换为DTO
func (s *outfitService) convertToOutfitDTO(ctx context.Context, outfit *models.Outfit) (*dto.Outfit, error) {
	// 获取穿搭单品
	outfitItems, err := s.outfitItemRepo.GetByOutfitID(ctx, outfit.ID)
	if err != nil {
		return nil, fmt.Errorf("获取穿搭单品失败: %w", err)
	}

	// 获取衣物详情并转换
	clothingItems := make([]dto.OutfitClothingItem, 0, len(outfitItems))
	for _, item := range outfitItems {
		clothingItem, err := s.clothingItemRepo.GetByID(ctx, item.ClothingItemID)
		if err != nil {
			continue // 跳过不存在的衣物
		}

		// 获取分类名称
		var categoryName string
		if category, err := s.clothingCategoryRepo.GetByID(ctx, clothingItem.CategoryID); err == nil {
			categoryName = category.Name
		}

		// 获取主图片
		var imageURL string
		// 这里简化实现，实际中应该获取第一张图片

		clothingDTO := dto.OutfitClothingItem{
			ID:           clothingItem.ID,
			Name:         clothingItem.Name,
			Brand:        clothingItem.Brand,
			Color:        clothingItem.Color,
			CategoryName: categoryName,
			ImageURL:     imageURL,
			Layer:        item.LayerOrder,
			Position:     item.ItemRole,
		}
		clothingItems = append(clothingItems, clothingDTO)
	}

	// 获取附件信息
	attachments, err := s.attachmentRepo.GetByEntityID(ctx, api.EntityTypeOutfit, outfit.ID)
	if err != nil {
		// 如果获取附件失败，记录日志但不中断流程
		attachments = []models.Attachment{}
	}

	// 转换附件为DTO
	attachmentDTOs := make([]dto.AttachmentDTO, 0, len(attachments))
	for _, attachment := range attachments {
		attachmentDTO := dto.AttachmentDTO{
			ID:             attachment.ID,
			OriginalName:   attachment.OriginalName,
			FileName:       attachment.FileName,
			FileSize:       attachment.FileSize,
			MimeType:       attachment.MimeType,
			AttachmentType: attachment.AttachmentType,
			EntityType:     attachment.EntityType,
			EntityID:       attachment.EntityID,
			PublicURL:      attachment.PublicURL,
			Description:    attachment.Description,
			Tags:           attachment.Tags,
			SortOrder:      attachment.SortOrder,
			CreatedAt:      attachment.CreatedAt,
			UpdatedAt:      attachment.UpdatedAt,
		}
		if attachment.Width != nil {
			attachmentDTO.Width = attachment.Width
		}
		if attachment.Height != nil {
			attachmentDTO.Height = attachment.Height
		}
		if attachment.Duration != nil {
			attachmentDTO.Duration = attachment.Duration
		}
		if attachment.Thumbnail != nil {
			attachmentDTO.Thumbnail = *attachment.Thumbnail
		}
		attachmentDTOs = append(attachmentDTOs, attachmentDTO)
	}

	outfitDTO := &dto.Outfit{
		ID:          outfit.ID,
		UserID:      outfit.UserID,
		Name:        outfit.Name,
		Date:        outfit.Date,
		Temperature: outfit.Temperature,
		Weather:     outfit.Weather,
		Occasion:    outfit.Occasion,
		Location:    outfit.Location,
		Notes:       outfit.Notes,
		IsPublic:    outfit.IsPublic,
		ClothingIDs: s.extractClothingIDs(outfitItems),
		Tags:        outfit.Tags,
		CreatedAt:   outfit.CreatedAt,
		UpdatedAt:   outfit.UpdatedAt,
	}

	return outfitDTO, nil
}

// extractClothingIDs 提取衣物ID列表
func (s *outfitService) extractClothingIDs(outfitItems []models.OutfitItem) []uint {
	ids := make([]uint, 0, len(outfitItems))
	for _, item := range outfitItems {
		ids = append(ids, item.ClothingItemID)
	}
	return ids
}

// generateRecommendations 生成推荐算法（简化实现）
func (s *outfitService) generateRecommendations(ctx context.Context, clothingItems []models.ClothingItem, weather api.WeatherType) []dto.RecommendedClothingItem {
	recommendations := make([]dto.RecommendedClothingItem, 0)

	// 按分类分组衣物
	itemsByCategory := make(map[string][]models.ClothingItem)
	for _, item := range clothingItems {
		// 获取分类名称
		category, err := s.clothingCategoryRepo.GetByID(ctx, item.CategoryID)
		if err != nil {
			continue
		}
		itemsByCategory[category.Name] = append(itemsByCategory[category.Name], item)
	}

	// 根据天气推荐合适的单品
	recommendationStrategy := s.getRecommendationStrategy(weather)

	for categoryName, strategy := range recommendationStrategy {
		if items, exists := itemsByCategory[categoryName]; exists && len(items) > 0 {
			// 选择第一个合适的单品（简化逻辑）
			selectedItem := items[0]

			recommendation := dto.RecommendedClothingItem{
				ID:           selectedItem.ID,
				Name:         selectedItem.Name,
				Brand:        selectedItem.Brand,
				Color:        selectedItem.Color,
				CategoryName: categoryName,
				Position:     strategy.Position,
				Layer:        strategy.Layer,
				Confidence:   strategy.Confidence,
				Reason:       strategy.Reason,
			}
			recommendations = append(recommendations, recommendation)
		}
	}

	return recommendations
}

// RecommendationStrategy 推荐策略
type RecommendationStrategy struct {
	Position   string
	Layer      int
	Confidence float64
	Reason     string
}

// getRecommendationStrategy 根据天气获取推荐策略
func (s *outfitService) getRecommendationStrategy(weather api.WeatherType) map[string]RecommendationStrategy {
	strategies := make(map[string]RecommendationStrategy)

	switch weather {
	case api.WeatherTypeSunny:
		strategies["上衣"] = RecommendationStrategy{
			Position:   "top",
			Layer:      1,
			Confidence: 0.9,
			Reason:     "晴天适合轻薄透气的上衣",
		}
		strategies["裤子"] = RecommendationStrategy{
			Position:   "bottom",
			Layer:      1,
			Confidence: 0.9,
			Reason:     "晴天适合轻便的下装",
		}
		strategies["鞋子"] = RecommendationStrategy{
			Position:   "shoes",
			Layer:      1,
			Confidence: 0.8,
			Reason:     "晴天可选择透气的鞋子",
		}
	case api.WeatherTypeRainy:
		strategies["外套"] = RecommendationStrategy{
			Position:   "outer",
			Layer:      2,
			Confidence: 0.95,
			Reason:     "雨天需要防水外套",
		}
		strategies["裤子"] = RecommendationStrategy{
			Position:   "bottom",
			Layer:      1,
			Confidence: 0.8,
			Reason:     "雨天适合不易湿的下装",
		}
		strategies["鞋子"] = RecommendationStrategy{
			Position:   "shoes",
			Layer:      1,
			Confidence: 0.9,
			Reason:     "雨天需要防水鞋",
		}
	case api.WeatherTypeSnowy:
		strategies["外套"] = RecommendationStrategy{
			Position:   "outer",
			Layer:      3,
			Confidence: 0.95,
			Reason:     "雪天需要保暖外套",
		}
		strategies["内衣"] = RecommendationStrategy{
			Position:   "inner",
			Layer:      1,
			Confidence: 0.9,
			Reason:     "雪天需要保暖内衣",
		}
		strategies["靴子"] = RecommendationStrategy{
			Position:   "shoes",
			Layer:      1,
			Confidence: 0.9,
			Reason:     "雪天需要保暖防滑靴",
		}
	default:
		// 默认推荐
		strategies["上衣"] = RecommendationStrategy{
			Position:   "top",
			Layer:      1,
			Confidence: 0.7,
			Reason:     "基础搭配",
		}
		strategies["裤子"] = RecommendationStrategy{
			Position:   "bottom",
			Layer:      1,
			Confidence: 0.7,
			Reason:     "基础搭配",
		}
	}

	return strategies
}

// getTemperatureByWeather 根据天气获取参考温度
func (s *outfitService) getTemperatureByWeather(weather api.WeatherType) float64 {
	switch weather {
	case api.WeatherTypeSunny:
		return 25.0
	case api.WeatherTypeRainy:
		return 18.0
	case api.WeatherTypeCloudy:
		return 20.0
	case api.WeatherTypeSnowy:
		return 0.0
	case api.WeatherTypeFoggy:
		return 15.0
	case api.WeatherTypeWindy:
		return 15.0
	default:
		return 20.0
	}
}

// getWeatherDescription 获取天气描述
func (s *outfitService) getWeatherDescription(weather api.WeatherType) string {
	switch weather {
	case api.WeatherTypeSunny:
		return "晴朗温暖，适合轻薄衣物"
	case api.WeatherTypeRainy:
		return "阴雨天气，注意防水保暖"
	case api.WeatherTypeCloudy:
		return "多云天气，温度适中"
	case api.WeatherTypeSnowy:
		return "雪天寒冷，注意保暖防滑"
	case api.WeatherTypeFoggy:
		return "雾天湿润，注意保暖"
	case api.WeatherTypeWindy:
		return "大风天气，注意防风保暖"
	default:
		return "天气状况良好"
	}
}

// calculateConfidence 计算推荐置信度
func (s *outfitService) calculateConfidence(items []dto.RecommendedClothingItem) float64 {
	if len(items) == 0 {
		return 0.0
	}

	totalConfidence := 0.0
	for _, item := range items {
		totalConfidence += item.Confidence
	}

	avgConfidence := totalConfidence / float64(len(items))
	return math.Round(avgConfidence*100) / 100 // 保留2位小数
}

// generateRecommendationReason 生成推荐理由
func (s *outfitService) generateRecommendationReason(weather api.WeatherType, items []dto.RecommendedClothingItem) string {
	if len(items) == 0 {
		return "暂无合适的推荐"
	}

	weatherDesc := s.getWeatherDescription(weather)
	itemCount := len(items)

	return fmt.Sprintf("基于%s的天气条件，为您推荐了%d件单品的搭配方案", weatherDesc, itemCount)
}
