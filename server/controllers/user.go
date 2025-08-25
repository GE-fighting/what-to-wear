package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/api/errors"
	"what-to-wear/server/services"
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
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	user, err := uc.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, api.Success(user, "获取用户资料成功"))
}

// UpdateProfile 更新用户资料
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	var req dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证性别枚举
	if req.Gender != nil && !req.Gender.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的性别类型"))
		return
	}

	user, err := uc.userService.UpdateProfile(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(user, "用户资料更新成功"))
}

// ChangePassword 修改密码
func (uc *UserController) ChangePassword(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	var req dto.ChangePasswordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	err := uc.userService.ChangePassword(c.Request.Context(), userID, req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(nil, "密码修改成功"))
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	targetUserID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的用户ID"))
		return
	}

	// 检查是否是用户本人或管理员
	currentUserID := getUserID(c)
	if currentUserID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	if currentUserID != uint(targetUserID) {
		c.JSON(http.StatusForbidden, api.Forbidden("无权限删除该用户"))
		return
	}

	err = uc.userService.DeleteUser(c.Request.Context(), uint(targetUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(nil, "用户删除成功"))
}

// GetUserStats 获取用户统计信息
func (uc *UserController) GetUserStats(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(api.StatusUnauthorized, errors.ErrUnauthorized("未授权访问"))
		return
	}
	// 暂时返回模拟数据
	stats := dto.UserStatsDTO{
		TotalClothingItems: 0,
		TotalOutfits:       0,
		TotalSpent:         0.0,
		AccountAge:         0,
	}

	c.JSON(http.StatusOK, api.Success(stats, "获取用户统计成功"))
}

// UpdatePreferences 更新用户偏好设置
func (uc *UserController) UpdatePreferences(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	var req dto.UserPreferencesDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证分页设置
	if req.DisplaySettings.ItemsPerPage > api.MaxPageSize {
		req.DisplaySettings.ItemsPerPage = api.MaxPageSize
	}
	if req.DisplaySettings.ItemsPerPage <= 0 {
		req.DisplaySettings.ItemsPerPage = api.DefaultPageSize
	}

	// TODO: 实现更新用户偏好设置的逻辑
	// err := uc.userService.UpdatePreferences(userID, &req)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
	//     return
	// }

	c.JSON(http.StatusOK, api.Success(nil, "偏好设置更新成功"))
}

// GetPreferences 获取用户偏好设置
func (uc *UserController) GetPreferences(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	// TODO: 实现获取用户偏好设置的逻辑
	// preferences, err := uc.userService.GetPreferences(userID)
	// if err != nil {
	//     common.Error(c, err)
	//     return
	// }

	// 暂时返回默认偏好设置
	preferences := dto.UserPreferencesDTO{
		Language: "zh-CN",
		Timezone: "Asia/Shanghai",
		Currency: "CNY",
		Notifications: dto.NotificationPrefsDTO{
			EmailNotifications:   true,
			PushNotifications:    true,
			MaintenanceReminders: true,
			OutfitSuggestions:    true,
			WeatherAlerts:        true,
		},
		PrivacySettings: dto.PrivacySettingsDTO{
			ProfileVisibility: "private",
			ShowRealName:      false,
			ShowLocation:      false,
			AllowDataExport:   true,
		},
		DisplaySettings: dto.DisplaySettingsDTO{
			Theme:         "auto",
			GridSize:      "medium",
			ShowPrices:    true,
			ShowWearCount: true,
			DefaultSortBy: "created_at",
			ItemsPerPage:  20,
		},
	}

	c.JSON(http.StatusOK, api.Success(preferences, "获取偏好设置成功"))
}

// ExportData 导出用户数据
func (uc *UserController) ExportData(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}
	c.JSON(http.StatusOK, api.Success(gin.H{"message": "数据导出请求已提交，将通过邮件发送下载链接"}, "数据导出请求成功"))
}
