package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/services"
)

// OSSController OSS控制器
type OSSController struct {
	ossService services.OSSService
}

// NewOSSController 创建OSS控制器实例
func NewOSSController(ossService services.OSSService) *OSSController {
	return &OSSController{
		ossService: ossService,
	}
}

// GeneratePresignedURL 生成预签名URL
func (oc *OSSController) GeneratePresignedURL(c *gin.Context) {
	var req dto.GeneratePresignedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 不再从前端获取 bucketName 和 expires，由服务端配置决定
	url, err := oc.ossService.GeneratePresignedURL(req.FileName, req.FileType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError("Failed to generate presigned URL: "+err.Error()))
		return
	}

	response := dto.GeneratePresignedURLResponse{
		URL: url,
	}

	c.JSON(http.StatusOK, api.Success(response, "预签名URL生成成功"))
}

// GenerateDownloadURL 生成下载URL
func (oc *OSSController) GenerateDownloadURL(c *gin.Context) {
	fileName := c.Query("file_name")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, api.BadRequest("file_name is required"))
		return
	}

	url, err := oc.ossService.GenerateDownloadURL(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError("Failed to generate download URL: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(gin.H{"url": url}, "下载URL生成成功"))
}