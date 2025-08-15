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

// AttachmentController 附件控制器
type AttachmentController struct {
	attachmentService services.AttachmentServiceInterface
}

// NewAttachmentController 创建附件控制器实例
func NewAttachmentController(attachmentService services.AttachmentServiceInterface) *AttachmentController {
	return &AttachmentController{
		attachmentService: attachmentService,
	}
}

// UploadAttachment 上传附件
func (ac *AttachmentController) UploadAttachment(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.UploadAttachmentRequest
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("请求参数错误", err.Error()))
		return
	}

	req.UserID = userID

	// 验证实体类型
	if !models.IsValidEntityType(string(req.EntityType)) {
		common.Error(c, errors.ErrInvalidRequest("无效的实体类型"))
		return
	}

	response, err := ac.attachmentService.UploadAttachment(&req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, response, "附件上传成功")
}

// BatchUploadAttachments 批量上传附件
func (ac *AttachmentController) BatchUploadAttachments(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.BatchUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("请求参数错误", err.Error()))
		return
	}

	req.UserID = userID

	response, err := ac.attachmentService.BatchUploadAttachments(&req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, response, "批量上传完成")
}

// GetAttachmentsByEntity 获取指定实体的附件列表
func (ac *AttachmentController) GetAttachmentsByEntity(c *gin.Context) {
	entityTypeStr := c.Param("entity_type")
	if !models.IsValidEntityType(entityTypeStr) {
		common.Error(c, errors.ErrInvalidRequest("无效的实体类型"))
		return
	}

	entityID, ok := parseUintParamRequired(c, "entity_id")
	if !ok {
		return
	}

	attachments, err := ac.attachmentService.GetAttachmentsByEntity(
		models.EntityType(entityTypeStr), entityID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, attachments, "获取附件列表成功")
}

// GetUserAttachments 获取用户的附件列表
func (ac *AttachmentController) GetUserAttachments(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	limit := parseIntQuery(c, "limit", 20)
	offset := parseIntQuery(c, "offset", 0)

	// 验证分页参数
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	response, err := ac.attachmentService.GetAttachmentsByUser(userID, limit, offset)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, response, "获取用户附件成功")
}

// DeleteAttachment 删除附件
func (ac *AttachmentController) DeleteAttachment(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	attachmentID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	err := ac.attachmentService.DeleteAttachment(attachmentID, userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "附件删除成功")
}

// UpdateAttachmentOrder 更新附件排序
func (ac *AttachmentController) UpdateAttachmentOrder(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	var req dto.UpdateAttachmentOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("请求参数错误", err.Error()))
		return
	}

	err := ac.attachmentService.UpdateAttachmentOrder(req.AttachmentID, req.SortOrder, userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "排序更新成功")
}

// GetAttachmentInfo 获取附件详细信息
func (ac *AttachmentController) GetAttachmentInfo(c *gin.Context) {
	attachmentID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	// 获取基本附件信息
	attachment, err := ac.attachmentService.GetAttachment(attachmentID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, attachment, "获取附件信息成功")
}

// UpdateAttachmentInfo 更新附件信息
func (ac *AttachmentController) UpdateAttachmentInfo(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	attachmentID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("请求参数错误", err.Error()))
		return
	}

	attachment, err := ac.attachmentService.UpdateAttachment(attachmentID, userID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, attachment, "附件信息更新成功")
}

// BatchDeleteAttachments 批量删除附件
func (ac *AttachmentController) BatchDeleteAttachments(c *gin.Context) {
	userID, ok := getUserIDRequired(c)
	if !ok {
		return
	}

	type BatchDeleteRequest struct {
		AttachmentIDs []uint `json:"attachment_ids" binding:"required,min=1"`
	}

	var req BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("请求参数错误", err.Error()))
		return
	}

	// 逐个删除附件
	successCount := 0
	var deleteErrors []string

	for _, id := range req.AttachmentIDs {
		if err := ac.attachmentService.DeleteAttachment(id, userID); err != nil {
			deleteErrors = append(deleteErrors, fmt.Sprintf("删除附件 %d 失败: %v", id, err))
		} else {
			successCount++
		}
	}

	result := map[string]interface{}{
		"success_count": successCount,
		"failure_count": len(req.AttachmentIDs) - successCount,
		"errors":        deleteErrors,
	}

	common.Success(c, result, "批量删除完成")
}

// GetAttachmentStats 获取附件统计信息
func (ac *AttachmentController) GetAttachmentStats(c *gin.Context) {
	// 暂时返回基础统计信息
	stats := map[string]interface{}{
		"total_attachments": 0,
		"total_size":        0,
		"by_type":           map[string]int{},
		"by_entity":         map[string]int{},
	}

	common.Success(c, stats, "获取附件统计成功")
}
