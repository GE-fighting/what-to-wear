package routes

import (
	"what-to-wear/server/controllers"

	"github.com/gin-gonic/gin"
)

// setupAuthRoutes 设置认证相关路由
func setupAuthRoutes(api *gin.RouterGroup, authController *controllers.AuthController) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		// 可以添加更多认证相关的路由
		// auth.POST("/logout", authController.Logout)
		// auth.POST("/refresh", authController.RefreshToken)
		// auth.POST("/forgot-password", authController.ForgotPassword)
	}
}
