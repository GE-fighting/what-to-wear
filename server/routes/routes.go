package routes

import (
	"what-to-wear/server/container"
	"what-to-wear/server/middleware"

	"github.com/gin-gonic/gin"
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

// setupProtectedRoutes 设置需要认证的路由
func setupProtectedRoutes(api *gin.RouterGroup, container *container.Container) {
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户相关路由
		setupUserRoutes(protected, container.GetUserController())

		// 其他需要认证的路由可以在这里添加
		// setupClothingRoutes(protected)
	}
}
