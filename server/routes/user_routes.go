package routes

import (
	"what-to-wear/server/controllers"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes 设置用户相关路由
func setupUserRoutes(protected *gin.RouterGroup, userController *controllers.UserController) {
	user := protected.Group("/user")
	{
		user.GET("/profile", userController.GetProfile)
		user.PUT("/profile", userController.UpdateProfile)
		user.PUT("/password", userController.ChangePassword)
		user.DELETE("/:id", userController.DeleteUser)
		// 可以添加更多用户相关的路由
		// user.POST("/avatar", uploadAvatar)
		// user.GET("/preferences", getUserPreferences)
		// user.PUT("/preferences", updateUserPreferences)
	}
}
