package routes

import (
	"github.com/gin-gonic/gin"
	"what-to-wear/server/container"
)

// SetupRoutes 配置所有路由
func SetupRoutes(r *gin.Engine, container *container.Container) {
	// API路由组
	api := r.Group("/api")
	{
		// 设置公开路由（不需要认证）
		setupPublicRoutes(api, container)

		// 设置需要认证的路由
		setupProtectedRoutes(api, container)
	}
}

// setupPublicRoutes 设置公开路由
func setupPublicRoutes(api *gin.RouterGroup, container *container.Container) {
	// 认证相关路由
	setupAuthRoutes(api, container.GetAuthController())

	// 其他公开路由
	setupPublicAPIRoutes(api)
}

// setupProtectedRoutes 路由认证在不同分类路由内部做
func setupProtectedRoutes(api *gin.RouterGroup, container *container.Container) {
	{
		// 用户相关路由
		setupUserRoutes(api, container.GetUserController())

		// 衣服相关路由
		SetupClothingRoutes(api, container.GetClothingController())
	}
}
