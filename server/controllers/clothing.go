package controllers

import (
	"fmt"
	"net/http"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/services"

	"github.com/gin-gonic/gin"
)

// ClothingController 衣物控制器
type ClothingController struct {
	clothingService   services.ClothingItemService
	categoryService   services.ClothingCategoryService
	tagService        services.ClothingTagService
	wearRecordService services.WearRecordService
}

// NewClothingController 创建衣物控制器
func NewClothingController(
	clothingService services.ClothingItemService,
	categoryService services.ClothingCategoryService,
	tagService services.ClothingTagService,
	wearRecordService services.WearRecordService,
) *ClothingController {
	return &ClothingController{
		clothingService:   clothingService,
		categoryService:   categoryService,
		tagService:        tagService,
		wearRecordService: wearRecordService,
	}
}

// CreateClothingItem 创建衣物
func (cc *ClothingController) CreateClothingItem(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.CreateClothingItemDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证衣物状态枚举
	if req.Status != "" && !req.Status.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的衣物状态"))
		return
	}

	item, err := cc.clothingService.CreateClothingItem(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, api.Success(item, "衣物创建成功"))
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

	item, err := cc.clothingService.GetClothingItem(c.Request.Context(), userID, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(item, "获取衣物成功"))
}

// GetClothingItems 获取衣物列表
func (cc *ClothingController) GetClothingItems(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.ClothingItemListDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("查询参数错误: "+err.Error()))
		return
	}
	// 验证并设置默认值
	validatePagination(&req.Page, &req.PageSize)

	// 验证排序参数
	if req.SortBy != "" && !isValidClothingSortBy(req.SortBy) {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的排序字段"))
		return
	}
	if req.SortOrder != "" && !isValidSortOrder(req.SortOrder) {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的排序方向"))
		return
	}
	// 验证衣物状态枚举
	if req.Status != nil && !req.Status.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的衣物状态"))
		return
	}

	items, total, err := cc.clothingService.GetClothingItems(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.SuccessWithPage(items, int64(total), req.Page, req.PageSize, "获取衣物列表成功"))
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

	var req dto.UpdateClothingItemDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证衣物状态枚举
	if req.Status != nil && !req.Status.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的衣物状态"))
		return
	}

	item, err := cc.clothingService.UpdateClothingItem(c.Request.Context(), userID, itemID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(item, "衣物更新成功"))
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

	err := cc.clothingService.DeleteClothingItem(c.Request.Context(), userID, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(nil, "衣物删除成功"))
}

// GetClothingStats 获取衣物统计
func (cc *ClothingController) GetClothingStats(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	stats, err := cc.clothingService.GetClothingStats(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(stats, "获取衣物统计成功"))
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
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 逐个删除衣物
	successCount := 0
	var deleteErrors []string

	for _, itemID := range req.ItemIDs {
		if err := cc.clothingService.DeleteClothingItem(c.Request.Context(), userID, itemID); err != nil {
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

	c.JSON(http.StatusOK, api.Success(result, "批量删除完成"))
}

// GetCategories 获取分类列表
func (cc *ClothingController) GetCategories(c *gin.Context) {
	categories, err := cc.categoryService.GetAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(categories, "获取分类列表成功"))
}

// GetCategoryTree 获取分类树
func (cc *ClothingController) GetCategoryTree(c *gin.Context) {
	tree, err := cc.categoryService.GetCategoryTree(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(tree, "获取分类树成功"))
}

// CreateCategory 创建分类
func (cc *ClothingController) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	category, err := cc.categoryService.CreateCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, api.Success(category, "分类创建成功"))
}

// UpdateCategory 更新分类
func (cc *ClothingController) UpdateCategory(c *gin.Context) {
	categoryID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateCategoryDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	category, err := cc.categoryService.UpdateCategory(c.Request.Context(), categoryID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(category, "分类更新成功"))
}

// GetTags 获取标签列表
func (cc *ClothingController) GetTags(c *gin.Context) {
	userID := getUserID(c)

	tags, err := cc.tagService.GetAllTags(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(tags, "获取标签列表成功"))
}

// GetTagsByType 根据类型获取标签
func (cc *ClothingController) GetTagsByType(c *gin.Context) {
	userID := getUserID(c)
	tagType := c.Param("type")

	// 验证标签类型
	if !isValidTagType(tagType) {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的标签类型"))
		return
	}

	tags, err := cc.tagService.GetTagsByType(c.Request.Context(), api.TagType(tagType), &userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(tags, "获取标签成功"))
}

// CreateTag 创建标签
func (cc *ClothingController) CreateTag(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.CreateTagDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证标签类型
	if !isValidTagType(req.Type) {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的标签类型"))
		return
	}

	tag, err := cc.tagService.CreateTag(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, api.Success(tag, "标签创建成功"))
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

	var req dto.UpdateTagDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	tag, err := cc.tagService.UpdateTag(c.Request.Context(), userID, tagID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(tag, "标签更新成功"))
}

// RecordWear 记录穿着
func (cc *ClothingController) RecordWear(c *gin.Context) {

	var req dto.CreateWearRecordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	c.JSON(http.StatusCreated, api.Success(nil, "穿着记录添加成功"))
}
