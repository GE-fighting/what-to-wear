package services

import (
	"context"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/api/errors"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

// ClothingTagService 衣物标签服务接口
type ClothingTagService interface {
	// 基础CRUD操作
	CreateTag(ctx context.Context, userID uint, req *dto.CreateTagDTO) (*dto.TagDTO, error)
	GetTag(ctx context.Context, tagID uint) (*dto.TagDTO, error)
	GetAllTags(ctx context.Context, userID uint) ([]dto.TagDTO, error)
	UpdateTag(ctx context.Context, userID, tagID uint, req *dto.UpdateTagDTO) (*dto.TagDTO, error)
	DeleteTag(ctx context.Context, userID, tagID uint) error

	// 按类型查询
	GetTagsByType(ctx context.Context, tagType api.TagType, userID *uint) ([]dto.TagDTO, error)
	GetSystemTags(ctx context.Context) ([]dto.TagDTO, error)
	GetUserTags(ctx context.Context, userID uint) ([]dto.TagDTO, error)

	// 系统标签枚举查询（从内存获取）
	GetSystemTagEnumsByType(tagType api.TagType) ([]dto.TagDTO, error)
	GetAllSystemTagEnums() (map[string][]dto.TagDTO, error)

	// 统计
	GetPopularTags(ctx context.Context, userID uint, limit int) ([]dto.TagDTO, error)
	GetTagStats(ctx context.Context, userID uint) ([]dto.TagStatsItem, error)
}

// clothingTagService 衣物标签服务实现
type clothingTagService struct {
	tagRepo repositories.ClothingTagRepository
}

// NewClothingTagService 创建衣物标签服务实例
func NewClothingTagService(tagRepo repositories.ClothingTagRepository) ClothingTagService {
	return &clothingTagService{
		tagRepo: tagRepo,
	}
}

// CreateTag 创建标签
func (s *clothingTagService) CreateTag(ctx context.Context, userID uint, req *dto.CreateTagDTO) (*dto.TagDTO, error) {
	// 验证标签类型
	if !api.IsValidTagType(req.Type) {
		return nil, errors.ErrInvalidRequest("invalid tag type")
	}

	// 创建模型
	tag := &models.ClothingTag{
		Name:        req.Name,
		Type:        api.TagType(req.Type),
		Description: req.Description,
		IsSystem:    false, // 用户创建的标签都不是系统标签
		IsActive:    true,
		UserID:      &userID,
	}

	// 调用仓库层创建标签
	if err := s.tagRepo.Create(ctx, tag); err != nil {
		return nil, err
	}

	// 转换为DTO返回
	return s.convertToDTO(tag), nil
}

// GetTag 根据ID获取标签
func (s *clothingTagService) GetTag(ctx context.Context, tagID uint) (*dto.TagDTO, error) {
	tag, err := s.tagRepo.GetByID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	return s.convertToDTO(tag), nil
}

// GetAllTags 获取所有标签（系统标签和用户标签）
func (s *clothingTagService) GetAllTags(ctx context.Context, userID uint) ([]dto.TagDTO, error) {
	tags, err := s.tagRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return s.convertToDTOList(tags), nil
}

// UpdateTag 更新标签
func (s *clothingTagService) UpdateTag(ctx context.Context, userID, tagID uint, req *dto.UpdateTagDTO) (*dto.TagDTO, error) {
	// 先获取标签
	tag, err := s.tagRepo.GetByID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	// 检查权限：只有标签创建者或系统管理员可以更新标签
	if tag.UserID != nil && *tag.UserID != userID && !tag.IsSystem {
		return nil, errors.ErrForbidden("permission denied")
	}

	// 系统标签不允许更新
	if tag.IsSystem {
		return nil, errors.ErrForbidden("system tags cannot be updated")
	}

	// 更新字段
	if req.Name != nil {
		tag.Name = *req.Name
	}
	if req.Description != nil {
		tag.Description = *req.Description
	}

	// 调用仓库层更新
	if err := s.tagRepo.Update(ctx, tag); err != nil {
		return nil, err
	}

	return s.convertToDTO(tag), nil
}

// DeleteTag 删除标签
func (s *clothingTagService) DeleteTag(ctx context.Context, userID, tagID uint) error {
	// 先获取标签
	tag, err := s.tagRepo.GetByID(ctx, tagID)
	if err != nil {
		return err
	}

	// 检查权限：只有标签创建者或系统管理员可以删除标签
	if tag.UserID != nil && *tag.UserID != userID && !tag.IsSystem {
		return errors.ErrForbidden("permission denied")
	}

	// 系统标签不允许删除
	if tag.IsSystem {
		return errors.ErrForbidden("system tags cannot be deleted")
	}

	// 调用仓库层删除（软删除）
	return s.tagRepo.Delete(ctx, tagID)
}

// GetTagsByType 根据类型获取标签
func (s *clothingTagService) GetTagsByType(ctx context.Context, tagType api.TagType, userID *uint) ([]dto.TagDTO, error) {
	tags, err := s.tagRepo.GetByType(ctx, tagType, userID)
	if err != nil {
		return nil, err
	}

	return s.convertToDTOList(tags), nil
}

// GetSystemTags 获取系统标签
func (s *clothingTagService) GetSystemTags(ctx context.Context) ([]dto.TagDTO, error) {
	tags, err := s.tagRepo.GetSystemTags(ctx)
	if err != nil {
		return nil, err
	}

	return s.convertToDTOList(tags), nil
}

// GetUserTags 获取用户自定义标签
func (s *clothingTagService) GetUserTags(ctx context.Context, userID uint) ([]dto.TagDTO, error) {
	tags, err := s.tagRepo.GetUserTags(ctx, userID)
	if err != nil {
		return nil, err
	}

	return s.convertToDTOList(tags), nil
}

// GetPopularTags 获取热门标签
func (s *clothingTagService) GetPopularTags(ctx context.Context, userID uint, limit int) ([]dto.TagDTO, error) {
	tags, err := s.tagRepo.GetPopularTags(ctx, userID, limit)
	if err != nil {
		return nil, err
	}

	return s.convertToDTOList(tags), nil
}

// GetTagStats 获取标签统计
func (s *clothingTagService) GetTagStats(ctx context.Context, userID uint) ([]dto.TagStatsItem, error) {
	// 获取用户的所有标签
	tags, err := s.tagRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 获取标签使用统计
	statsMap, err := s.tagRepo.GetTagUsageStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 计算总数
	var total int64
	for _, count := range statsMap {
		total += count
	}

	// 构建统计结果
	var result []dto.TagStatsItem
	for _, tag := range tags {
		count := statsMap[tag.ID]
		percentage := 0.0
		if total > 0 {
			percentage = float64(count) / float64(total) * 100
		}

		result = append(result, dto.TagStatsItem{
			TagName:    tag.Name,
			Count:      count,
			Percentage: percentage,
		})
	}

	return result, nil
}

// GetSystemTagEnumsByType 根据标签类型获取系统标签枚举（从内存）
func (s *clothingTagService) GetSystemTagEnumsByType(tagType api.TagType) ([]dto.TagDTO, error) {
	// 从内存中获取系统标签枚举
	systemTags := api.GetSystemTagsByType(tagType)
	if len(systemTags) == 0 {
		return []dto.TagDTO{}, nil
	}

	// 转换为DTO
	result := make([]dto.TagDTO, len(systemTags))
	for i, tag := range systemTags {
		result[i] = dto.TagDTO{
			ID:          tag.ID,
			Name:        tag.Name,
			Type:        tag.Type,
			Description: tag.Description,
		}
	}

	return result, nil
}

// GetAllSystemTagEnums 获取所有系统标签枚举（从内存）
func (s *clothingTagService) GetAllSystemTagEnums() (map[string][]dto.TagDTO, error) {
	// 从内存中获取所有系统标签
	allSystemTags := api.GetAllSystemTags()

	result := make(map[string][]dto.TagDTO)
	for tagType, systemTags := range allSystemTags {
		tagTypeStr := string(tagType)
		tagDTOs := make([]dto.TagDTO, len(systemTags))

		for i, tag := range systemTags {
			tagDTOs[i] = dto.TagDTO{
				ID:          tag.ID,
				Name:        tag.Name,
				Type:        tag.Type,
				Description: tag.Description,
			}
		}

		result[tagTypeStr] = tagDTOs
	}

	return result, nil
}

// convertToDTO 将模型转换为DTO
func (s *clothingTagService) convertToDTO(tag *models.ClothingTag) *dto.TagDTO {
	return &dto.TagDTO{
		ID:          tag.ID,
		Name:        tag.Name,
		Type:        string(tag.Type),
		Description: tag.Description,
	}
}

// convertToDTOList 将模型列表转换为DTO列表
func (s *clothingTagService) convertToDTOList(tags []models.ClothingTag) []dto.TagDTO {
	result := make([]dto.TagDTO, len(tags))
	for i, tag := range tags {
		result[i] = *s.convertToDTO(&tag)
	}
	return result
}
