package routes

import (
	"what-to-wear/server/controllers"

	"github.com/gin-gonic/gin"
)

// setupOSSRoutes 设置OSS相关路由
func setupOSSRoutes(api *gin.RouterGroup, ossController *controllers.OSSController) {
	// 只有当OSS控制器不为nil时才注册路由
	if ossController == nil {
		return
	}

	oss := api.Group("/oss")
	{
		// 生成预签名上传URL
		oss.POST("/presign-upload", ossController.GeneratePresignedURL)

		// 生成预签名下载URL
		oss.POST("/presign-download", ossController.GenerateDownloadURL)
	}
}
