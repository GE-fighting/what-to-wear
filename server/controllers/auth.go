package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"what-to-wear/server/api"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/services"
)

// AuthController 认证控制器
type AuthController struct {
	authService services.AuthService
}

// NewAuthController 创建认证控制器实例
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register 用户注册
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// 验证性别枚举
	if req.Gender != "" && !req.Gender.IsValid() {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的性别类型"))
		return
	}

	user, err := ac.authService.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, api.Success(user, "用户注册成功"))
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	loginResp, err := ac.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))

		return
	}

	c.JSON(http.StatusOK, api.Success(loginResp, "登录成功"))
}

// RefreshToken 刷新访问令牌
func (ac *AuthController) RefreshToken(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return
	}

	token, err := ac.authService.RefreshToken(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(gin.H{"token": token}, "令牌刷新成功"))
}

// Logout 用户登出
func (ac *AuthController) Logout(c *gin.Context) {
	// 在实际应用中，这里可能需要将token加入黑名单
	// 目前只是返回成功响应
	c.JSON(http.StatusOK, api.Success(nil, "登出成功"))
}

// ValidateToken 验证令牌
func (ac *AuthController) ValidateToken(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("无效的令牌"))
		return
	}

	user, err := ac.authService.ValidateUser(c.Request.Context(), userID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, api.InternalError(err.Error()))

		return
	}

	c.JSON(http.StatusOK, api.Success(user, "令牌验证成功"))
}

// ForgotPassword 忘记密码
func (ac *AuthController) ForgotPassword(c *gin.Context) {
	type ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("无效的邮箱地址: "+err.Error()))
		return
	}

	// TODO: 实现发送重置密码邮件的逻辑
	// 这里暂时返回成功响应
	c.JSON(http.StatusOK, api.Success(nil, "密码重置邮件已发送"))
}

// ResetPassword 重置密码
func (ac *AuthController) ResetPassword(c *gin.Context) {
	type ResetPasswordRequest struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest("请求参数错误: "+err.Error()))
		return
	}

	// TODO: 实现验证重置令牌并更新密码的逻辑
	// 这里暂时返回成功响应
	c.JSON(http.StatusOK, api.Success(nil, "密码重置成功"))
}
