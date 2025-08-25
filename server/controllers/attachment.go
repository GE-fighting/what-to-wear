package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/services"
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

	var req dto.UploadAttachmentDTO
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(api.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	req.UserID = userID

	// 验证实体类型
	if !req.EntityType.IsValid() {
		c.JSON(api.StatusBadRequest, api.BadRequest("无效的实体类型"))
		return
	}

	response, err := ac.attachmentService.UploadAttachment(c.Request.Context(), &req)
	if err != nil {
		c.JSON(api.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, api.Success(response, "附件上传成功"))
}

// GetAttachmentsByEntity 获取指定实体的附件列表
func (ac *AttachmentController) GetAttachmentsByEntity(c *gin.Context) {
	entityTypeStr := c.Param("entity_type")
	entityType := api.EntityType(entityTypeStr)
	if !entityType.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的实体类型"))
		return
	}

	entityID, ok := parseUintParamRequired(c, "entity_id")
	if !ok {
		return
	}

	attachments, err := ac.attachmentService.GetAttachmentsByEntity(c.Request.Context(), entityType, entityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(attachments, "获取附件列表成功"))
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

	err := ac.attachmentService.DeleteAttachment(c.Request.Context(), attachmentID, userID)
	if err != nil {
		c.JSON(api.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, api.Success(nil, "附件删除成功"))
}

// GetAttachmentInfo 获取附件详细信息
func (ac *AttachmentController) GetAttachmentInfo(c *gin.Context) {
	attachmentID, ok := parseUintParamRequired(c, "id")
	if !ok {
		return
	}

	// 获取基本附件信息
	attachment, err := ac.attachmentService.GetAttachment(c.Request.Context(), attachmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(attachment, "获取附件信息成功"))
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

	var req dto.UpdateAttachmentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	attachment, err := ac.attachmentService.UpdateAttachment(c.Request.Context(), attachmentID, userID, &req)
	if err != nil {
		c.JSON(api.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(attachment, "附件信息更新成功"))
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
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 逐个删除附件
	successCount := 0
	var deleteErrors []string

	for _, id := range req.AttachmentIDs {
		if err := ac.attachmentService.DeleteAttachment(c.Request.Context(), id, userID); err != nil {
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

	c.JSON(http.StatusOK, api.Success(result, "批量删除完成"))
}
