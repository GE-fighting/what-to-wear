package controllers

import (
	"strconv"
	"what-to-wear/server/common"
	"what-to-wear/server/dto"
	"what-to-wear/server/errors"
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
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	user, err := uc.userService.GetProfile(userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, user, "获取用户资料成功")
}

// UpdateProfile 更新用户资料
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的更新资料请求", err.Error()))
		return
	}

	user, err := uc.userService.UpdateProfile(userID, &req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, user, "用户资料更新成功")
}

// ChangePassword 修改密码
func (uc *UserController) ChangePassword(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的修改密码请求", err.Error()))
		return
	}

	err := uc.userService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "密码修改成功")
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	targetUserID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的用户ID", err.Error()))
		return
	}

	// 检查是否是用户本人或管理员
	currentUserID := getUserID(c)
	if currentUserID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	if currentUserID != uint(targetUserID) {
		common.Error(c, errors.ErrForbidden("无权限删除该用户"))
		return
	}

	err = uc.userService.DeleteUser(uint(targetUserID))
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, nil, "用户删除成功")
}

// GetUserStats 获取用户统计信息
func (uc *UserController) GetUserStats(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	// TODO: 实现获取用户统计信息的逻辑
	// stats, err := uc.userService.GetUserStats(userID)
	// if err != nil {
	//     common.Error(c, err)
	//     return
	// }

	// 暂时返回模拟数据
	stats := dto.UserStatsResponse{
		TotalClothingItems: 0,
		TotalOutfits:       0,
		TotalSpent:         0.0,
		AccountAge:         0,
	}

	common.Success(c, stats, "获取用户统计成功")
}

// UpdatePreferences 更新用户偏好设置
func (uc *UserController) UpdatePreferences(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	var req dto.UserPreferencesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的偏好设置请求", err.Error()))
		return
	}

	// TODO: 实现更新用户偏好设置的逻辑
	// err := uc.userService.UpdatePreferences(userID, &req)
	// if err != nil {
	//     common.Error(c, err)
	//     return
	// }

	common.Success(c, nil, "偏好设置更新成功")
}

// GetPreferences 获取用户偏好设置
func (uc *UserController) GetPreferences(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	// TODO: 实现获取用户偏好设置的逻辑
	// preferences, err := uc.userService.GetPreferences(userID)
	// if err != nil {
	//     common.Error(c, err)
	//     return
	// }

	// 暂时返回默认偏好设置
	preferences := dto.UserPreferencesRequest{
		Language: "zh-CN",
		Timezone: "Asia/Shanghai",
		Currency: "CNY",
		Notifications: dto.NotificationPrefs{
			EmailNotifications:   true,
			PushNotifications:    true,
			MaintenanceReminders: true,
			OutfitSuggestions:    true,
			WeatherAlerts:        true,
		},
		PrivacySettings: dto.PrivacySettings{
			ProfileVisibility: "private",
			ShowRealName:      false,
			ShowLocation:      false,
			AllowDataExport:   true,
		},
		DisplaySettings: dto.DisplaySettings{
			Theme:         "auto",
			GridSize:      "medium",
			ShowPrices:    true,
			ShowWearCount: true,
			DefaultSortBy: "created_at",
			ItemsPerPage:  20,
		},
	}

	common.Success(c, preferences, "获取偏好设置成功")
}

// ExportData 导出用户数据
func (uc *UserController) ExportData(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	// TODO: 实现数据导出逻辑
	// exportData, err := uc.userService.ExportUserData(userID)
	// if err != nil {
	//     common.Error(c, err)
	//     return
	// }

	common.Success(c, gin.H{"message": "数据导出请求已提交，将通过邮件发送下载链接"}, "数据导出请求成功")
}
