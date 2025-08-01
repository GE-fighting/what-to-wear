package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes 设置用户相关路由
func setupUserRoutes(protected *gin.RouterGroup) {
	user := protected.Group("/user")
	{
		user.GET("/profile", getUserProfile)
		// 可以添加更多用户相关的路由
		// user.PUT("/profile", updateUserProfile)
		// user.POST("/avatar", uploadAvatar)
		// user.GET("/preferences", getUserPreferences)
		// user.PUT("/preferences", updateUserPreferences)
	}
}

// getUserProfile 获取用户资料
func getUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"username": username,
		"message":  "User profile retrieved successfully",
	})
}
