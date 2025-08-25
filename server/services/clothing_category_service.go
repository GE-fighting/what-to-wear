package services

import (
	"context"
	"fmt"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
)

// CategoryService分类服务接口
type ClothingCategoryService interface {
	GetCategoryTree(ctx context.Context) ([]dto.CategoryTreeNode, error)
	GetCategoryPath(ctx context.Context, categoryID uint) (string, error)
	GetAllCategories(ctx context.Context) ([]dto.CategoryDTO, error)
	GetRootCategories(ctx context.Context) ([]dto.CategoryDTO, error)
	GetCategory(ctx context.Context, categoryID uint) (*dto.CategoryDTO, error)
	CreateCategory(ctx context.Context, createCategoryDTO *dto.CreateCategoryDTO) (*dto.CategoryDTO, error)
	UpdateCategory(ctx context.Context, categoryID uint, updateCategoryDTO *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error)
	DeleteCategory(ctx context.Context, categoryID uint) error
	GetCategoryStats(ctx context.Context) ([]dto.CategoryStatsItem, error)
}

// CategoryServiceImpl 分类服务实现
type CategoryServiceImpl struct {
	categoryRepo repositories.ClothingCategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService(categoryRepo repositories.ClothingCategoryRepository) ClothingCategoryService {
	return &CategoryServiceImpl{
		categoryRepo: categoryRepo,
	}
}

// GetCategoryTree 获取分类树结构
func (s *CategoryServiceImpl) GetCategoryTree(ctx context.Context) ([]dto.CategoryTreeNode, error) {
	// 获取所有分类
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 构建分类映射
	categoryMap := make(map[uint]*dto.CategoryTreeNode)
	var rootNodes []dto.CategoryTreeNode

	// 先创建所有节点
	for _, category := range categories {
		node := &dto.CategoryTreeNode{
			CategoryDTO: dto.CategoryDTO{
				ID:          category.ID,
				Name:        category.Name,
				Description: category.Description,
				ParentID:    category.ParentID,
				Icon:        category.Icon,
				SortOrder:   category.SortOrder,
				IsActive:    category.IsActive,
			},
			Children: make([]dto.CategoryTreeNode, 0),
		}
		categoryMap[category.ID] = node
	}

	// 构建树结构
	for _, category := range categories {
		node := categoryMap[category.ID]

		if category.IsRootCategory() {
			// 根节点
			rootNodes = append(rootNodes, *node)
		} else {
			// 子节点，添加到父节点下
			if parent, exists := categoryMap[*category.ParentID]; exists {
				parent.Children = append(parent.Children, *node)
				// 设置父分类名称
				node.ParentName = parent.Name
			}
		}
	}

	return rootNodes, nil
}

// GetCategoryPath 获取分类的完整路径
func (s *CategoryServiceImpl) GetCategoryPath(ctx context.Context, categoryID uint) (string, error) {
	category, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return "", err
	}

	path := category.Name

	// 递归获取父分类路径
	for category.HasParent() {
		parentCategory, err := s.categoryRepo.GetByID(ctx, *category.ParentID)
		if err != nil {
			break
		}
		path = parentCategory.Name + " > " + path
		category = parentCategory
	}

	return path, nil
}

// GetAllCategories 获取所有分类
func (s *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]dto.CategoryDTO, error) {
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.CategoryDTO
	for _, category := range categories {
		response := dto.CategoryDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    category.ParentID,
			Icon:        category.Icon,
			SortOrder:   category.SortOrder,
			IsActive:    category.IsActive,
		}

		// 如果有父分类，获取父分类名称
		if category.HasParent() {
			if parentPath, err := s.GetCategoryPath(ctx, *category.ParentID); err == nil {
				response.ParentName = parentPath
			}
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// GetRootCategories 获取根分类
func (s *CategoryServiceImpl) GetRootCategories(ctx context.Context) ([]dto.CategoryDTO, error) {
	categories, err := s.categoryRepo.GetRootCategories(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.CategoryDTO
	for _, category := range categories {
		responses = append(responses, dto.CategoryDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    category.ParentID,
			Icon:        category.Icon,
			SortOrder:   category.SortOrder,
			IsActive:    category.IsActive,
		})
	}

	return responses, nil
}

// GetCategory 获取单个分类
func (s *CategoryServiceImpl) GetCategory(ctx context.Context, categoryID uint) (*dto.CategoryDTO, error) {
	category, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryDTO := &dto.CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
		Icon:        category.Icon,
		SortOrder:   category.SortOrder,
		IsActive:    category.IsActive,
	}

	// 获取父分类名称
	if category.HasParent() {
		if parentPath, err := s.GetCategoryPath(ctx, *category.ParentID); err == nil {
			categoryDTO.ParentName = parentPath
		}
	}

	// 获取关联的衣物数量
	if count, err := s.categoryRepo.GetCategoryItemCount(ctx, categoryID); err == nil {
		categoryDTO.ItemCount = count
	}

	return categoryDTO, nil
}

// CreateCategory 创建分类
func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, req *dto.CreateCategoryDTO) (*dto.CategoryDTO, error) {
	// 验证父分类是否存在
	if req.ParentID != nil {
		if _, err := s.categoryRepo.GetByID(ctx, *req.ParentID); err != nil {
			return nil, fmt.Errorf("父分类不存在")
		}
	}

	category := &models.ClothingCategory{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		Icon:        req.Icon,
		SortOrder:   req.SortOrder,
		IsActive:    true,
	}

	if err := s.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return s.GetCategory(ctx, category.ID)
}

// UpdateCategory 更新分类
func (s *CategoryServiceImpl) UpdateCategory(ctx context.Context, categoryID uint, req *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error) {
	category, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// 验证父分类
	if req.ParentID != nil && *req.ParentID != 0 {
		if *req.ParentID == categoryID {
			return nil, fmt.Errorf("不能将分类设为自己的父分类")
		}
		if _, err := s.categoryRepo.GetByID(ctx, *req.ParentID); err != nil {
			return nil, fmt.Errorf("父分类不存在")
		}
	}

	// 更新字段
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Description != nil {
		category.Description = *req.Description
	}
	if req.ParentID != nil {
		if *req.ParentID == 0 {
			category.ParentID = nil
		} else {
			category.ParentID = req.ParentID
		}
	}
	if req.Icon != nil {
		category.Icon = *req.Icon
	}
	if req.SortOrder != nil {
		category.SortOrder = *req.SortOrder
	}
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if err := s.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return s.GetCategory(ctx, categoryID)
}

// DeleteCategory 删除分类
func (s *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryID uint) error {
	// 检查是否有子分类
	children, err := s.categoryRepo.GetChildCategories(ctx, categoryID)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return fmt.Errorf("不能删除有子分类的分类")
	}

	// 检查是否有关联的衣物
	count, err := s.categoryRepo.GetCategoryItemCount(ctx, categoryID)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("不能删除有关联衣物的分类")
	}

	return s.categoryRepo.Delete(ctx, categoryID)
}

// GetCategoryStats 获取分类统计
func (s *CategoryServiceImpl) GetCategoryStats(ctx context.Context) ([]dto.CategoryStatsItem, error) {
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var stats []dto.CategoryStatsItem
	for _, category := range categories {
		count, _ := s.categoryRepo.GetCategoryItemCount(ctx, category.ID)

		stats = append(stats, dto.CategoryStatsItem{
			CategoryName: category.Name,
			Count:        count,
		})
	}

	return stats, nil
}
