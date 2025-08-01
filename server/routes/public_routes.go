package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupPublicAPIRoutes 设置公开API路由
func setupPublicAPIRoutes(api *gin.RouterGroup) {
	// 健康检查
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from go server!",
		})
	})

	// 可以添加更多公开的API
	// api.GET("/version", getVersion)
	// api.GET("/health", healthCheck)
}
