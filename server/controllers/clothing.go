package controllers

import (
	"fmt"
	"what-to-wear/server/common"
	"what-to-wear/server/dto"
	"what-to-wear/server/errors"
	"what-to-wear/server/models"
	"what-to-wear/server/services"

	"github.com/gin-gonic/gin"
)

// ClothingController 衣物控制器
type ClothingController struct {
	clothingService services.ClothingItemService
	categoryService services.ClothingCategoryService
	tagService      services.ClothingTagService
}

// NewClothingController 创建衣物控制器
func NewClothingController(
	clothingService services.ClothingItemService,
	categoryService services.ClothingCategoryService,
	tagService services.ClothingTagService,
) *ClothingController {
	return &ClothingController{
		clothingService: clothingService,
		categoryService: categoryService,
		tagService:      tagService,
	}
}

// CreateClothingItem 创建衣物
func (cc *ClothingController) CreateClothingItem(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.CreateClothingItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的创建衣物请求", err.Error()))
		return
	}

	item, err := cc.clothingService.CreateClothingItem(userID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, item, "衣物创建成功")
}

// GetClothingItem 获取单个衣物
func (cc *ClothingController) GetClothingItem(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	itemID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	item, err := cc.clothingService.GetClothingItem(userID, itemID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, item, "获取衣物成功")
}

// GetClothingItems 获取衣物列表
func (cc *ClothingController) GetClothingItems(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.ClothingItemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的查询参数", err.Error()))
		return
	}

	// 验证并设置默认值
	validatePagination(&req.Page, &req.PageSize)

	// 验证排序参数
	if req.SortBy != "" && !isValidClothingSortBy(req.SortBy) {
		common.Error(c, errors.ErrInvalidRequest("无效的排序字段"))
		return
	}
	if req.SortOrder != "" && !isValidSortOrder(req.SortOrder) {
		common.Error(c, errors.ErrInvalidRequest("无效的排序方向"))
		return
	}

	items, err := cc.clothingService.GetClothingItems(userID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, items, "获取衣物列表成功")
}

// UpdateClothingItem 更新衣物
func (cc *ClothingController) UpdateClothingItem(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	itemID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateClothingItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的更新衣物请求", err.Error()))
		return
	}

	item, err := cc.clothingService.UpdateClothingItem(userID, itemID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, item, "衣物更新成功")
}

// DeleteClothingItem 删除衣物
func (cc *ClothingController) DeleteClothingItem(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	itemID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	err := cc.clothingService.DeleteClothingItem(userID, itemID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "衣物删除成功")
}

// GetClothingStats 获取衣物统计
func (cc *ClothingController) GetClothingStats(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	stats, err := cc.clothingService.GetClothingStats(userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, stats, "获取衣物统计成功")
}

// BatchDeleteClothingItems 批量删除衣物
func (cc *ClothingController) BatchDeleteClothingItems(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	type BatchDeleteRequest struct {
		ItemIDs []uint `json:"item_ids" binding:"required,min=1"`
	}

	var req BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的批量删除请求", err.Error()))
		return
	}

	// 逐个删除衣物
	successCount := 0
	var deleteErrors []string

	for _, itemID := range req.ItemIDs {
		if err := cc.clothingService.DeleteClothingItem(userID, itemID); err != nil {
			deleteErrors = append(deleteErrors, fmt.Sprintf("删除衣物 %d 失败: %v", itemID, err))
		} else {
			successCount++
		}
	}

	result := map[string]interface{}{
		"success_count": successCount,
		"failure_count": len(req.ItemIDs) - successCount,
		"errors":        deleteErrors,
	}

	common.Success(c, result, "批量删除完成")
}

// GetCategories 获取分类列表
func (cc *ClothingController) GetCategories(c *gin.Context) {
	categories, err := cc.categoryService.GetAllCategories()
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, categories, "获取分类列表成功")
}

// GetCategoryTree 获取分类树
func (cc *ClothingController) GetCategoryTree(c *gin.Context) {
	tree, err := cc.categoryService.GetCategoryTree()
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, tree, "获取分类树成功")
}

// CreateCategory 创建分类
func (cc *ClothingController) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的创建分类请求", err.Error()))
		return
	}

	category, err := cc.categoryService.CreateCategory(&req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, category, "分类创建成功")
}

// UpdateCategory 更新分类
func (cc *ClothingController) UpdateCategory(c *gin.Context) {
	categoryID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的更新分类请求", err.Error()))
		return
	}

	category, err := cc.categoryService.UpdateCategory(categoryID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, category, "分类更新成功")
}

// GetTags 获取标签列表
func (cc *ClothingController) GetTags(c *gin.Context) {
	userID := getUserID(c)

	tags, err := cc.tagService.GetAllTags(userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, tags, "获取标签列表成功")
}

// GetTagsByType 根据类型获取标签
func (cc *ClothingController) GetTagsByType(c *gin.Context) {
	userID := getUserID(c)
	tagType := c.Param("type")

	// 验证标签类型
	if !isValidTagType(tagType) {
		common.Error(c, errors.ErrInvalidRequest("无效的标签类型"))
		return
	}

	tags, err := cc.tagService.GetTagsByType(models.TagType(tagType), &userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, tags, "获取标签成功")
}

// CreateTag 创建标签
func (cc *ClothingController) CreateTag(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的创建标签请求", err.Error()))
		return
	}

	// 验证标签类型
	if !isValidTagType(req.Type) {
		common.Error(c, errors.ErrInvalidRequest("无效的标签类型"))
		return
	}

	tag, err := cc.tagService.CreateTag(userID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, tag, "标签创建成功")
}

// UpdateTag 更新标签
func (cc *ClothingController) UpdateTag(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	tagID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的更新标签请求", err.Error()))
		return
	}

	tag, err := cc.tagService.UpdateTag(userID, tagID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, tag, "标签更新成功")
}

// RecordWear 记录穿着
func (cc *ClothingController) RecordWear(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	itemID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.CreateWearRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的穿着记录请求", err.Error()))
		return
	}

	err := cc.clothingService.RecordWear(userID, itemID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "穿着记录添加成功")
}

// ToggleFavorite 切换收藏状态
func (cc *ClothingController) ToggleFavorite(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	itemID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	err := cc.clothingService.ToggleFavorite(userID, itemID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "收藏状态更新成功")
}

// GetWearRecords 获取穿着记录
func (cc *ClothingController) GetWearRecords(c *gin.Context) {
	//userID, ok := getUserIDRequired(c)
	//if !ok {
	//	return
	//}
	//
	//itemID, ok := parseUintParamRequired(c, "id")
	//if !ok {
	//	return
	//}
	//
	//page := parseIntQuery(c, "page", 1)
	//pageSize := parseIntQuery(c, "page_size", 20)
	//validatePagination(&page, &pageSize)

	// 暂时返回空记录
	records := []map[string]interface{}{}

	common.Success(c, records, "获取穿着记录成功")
}

// GetMaintenanceRecords 获取保养记录
func (cc *ClothingController) GetMaintenanceRecords(c *gin.Context) {
	//userID, ok := getUserIDRequired(c)
	//if !ok {
	//	return
	//}
	//
	//itemID, ok := parseUintParamRequired(c, "id")
	//if !ok {
	//	return
	//}

	// 暂时返回空记录
	records := []map[string]interface{}{}

	common.Success(c, records, "获取保养记录成功")
}

// CreateMaintenanceRecord 创建保养记录
func (cc *ClothingController) CreateMaintenanceRecord(c *gin.Context) {
	//userID, ok := getUserIDRequired(c)
	//if !ok {
	//	return
	//}
	//
	//itemID, ok := parseUintParamRequired(c, "id")
	//if !ok {
	//	return
	//}

	var req dto.CreateMaintenanceRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的保养记录请求", err.Error()))
		return
	}

	// 暂时返回成功状态
	common.Created(c, nil, "保养记录创建成功")
}
