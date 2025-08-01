package routes

import (
	"what-to-wear/server/controllers"
	"what-to-wear/server/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置所有路由
func SetupRoutes(r *gin.Engine) {
	// 创建控制器实例
	authController := &controllers.AuthController{}

	// API路由组
	api := r.Group("/api")
	{
		// 设置公开路由（不需要认证）
		setupPublicRoutes(api, authController)

		// 设置需要认证的路由
		setupProtectedRoutes(api)
	}
}

// setupPublicRoutes 设置公开路由
func setupPublicRoutes(api *gin.RouterGroup, authController *controllers.AuthController) {
	// 认证相关路由
	setupAuthRoutes(api, authController)

	// 其他公开路由
	setupPublicAPIRoutes(api)
}

// setupProtectedRoutes 设置需要认证的路由
func setupProtectedRoutes(api *gin.RouterGroup) {
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户相关路由
		setupUserRoutes(protected)

		// 天气相关路由
		setupWeatherRoutes(protected)

		// 其他需要认证的路由可以在这里添加
		// setupClothingRoutes(protected)
	}
}
