package controllers

import (
	"strconv"
	"what-to-wear/server/common"
	"what-to-wear/server/services"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetProfile 获取用户资料
func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	user, err := uc.userService.GetProfile(userID)
	if err != nil {
		common.HandleError(c, err)
		return
	}

	common.SuccessResponse(c, user, "User profile retrieved successfully")
}

// UpdateProfile 更新用户资料
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req services.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleError(c, common.ErrInvalidRequest)
		return
	}

	user, err := uc.userService.UpdateProfile(userID, &req)
	if err != nil {
		common.HandleError(c, err)
		return
	}

	common.SuccessResponse(c, user, "User profile updated successfully")
}

// ChangePasswordRequest 修改密码请求结构体
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// ChangePassword 修改密码
func (uc *UserController) ChangePassword(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleError(c, common.ErrInvalidRequest)
		return
	}

	err := uc.userService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		common.HandleError(c, err)
		return
	}

	common.SuccessResponse(c, nil, "Password changed successfully")
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		common.HandleError(c, common.ErrInvalidRequest)
		return
	}

	// 检查是否是用户本人或管理员
	currentUserID := c.GetUint("user_id")
	if currentUserID != uint(userID) {
		common.HandleError(c, common.ErrForbidden)
		return
	}

	err = uc.userService.DeleteUser(uint(userID))
	if err != nil {
		common.HandleError(c, err)
		return
	}

	common.SuccessResponse(c, nil, "User deleted successfully")
}
