package controllers

import (
	"what-to-wear/server/common"
	"what-to-wear/server/dto"
	"what-to-wear/server/errors"
	"what-to-wear/server/services"

	"github.com/gin-gonic/gin"
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
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的注册请求", err.Error()))
		return
	}

	user, err := ac.authService.Register(&req)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Created(c, user, "用户注册成功")
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的登录请求", err.Error()))
		return
	}

	token, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, gin.H{"token": token}, "登录成功")
}

// RefreshToken 刷新访问令牌
func (ac *AuthController) RefreshToken(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("未授权访问"))
		return
	}

	token, err := ac.authService.RefreshToken(userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, gin.H{"token": token}, "令牌刷新成功")
}

// Logout 用户登出
func (ac *AuthController) Logout(c *gin.Context) {
	// 在实际应用中，这里可能需要将token加入黑名单
	// 目前只是返回成功响应
	common.Success(c, nil, "登出成功")
}

// ValidateToken 验证令牌
func (ac *AuthController) ValidateToken(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		common.Error(c, errors.ErrUnauthorized("无效的令牌"))
		return
	}

	user, err := ac.authService.ValidateUser(userID)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, user, "令牌验证成功")
}

// ForgotPassword 忘记密码
func (ac *AuthController) ForgotPassword(c *gin.Context) {
	type ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的邮箱地址", err.Error()))
		return
	}

	// TODO: 实现发送重置密码邮件的逻辑
	// 这里暂时返回成功响应
	common.Success(c, nil, "密码重置邮件已发送")
}

// ResetPassword 重置密码
func (ac *AuthController) ResetPassword(c *gin.Context) {
	type ResetPasswordRequest struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的重置密码请求", err.Error()))
		return
	}

	// TODO: 实现验证重置令牌并更新密码的逻辑
	// 这里暂时返回成功响应
	common.Success(c, nil, "密码重置成功")
}
