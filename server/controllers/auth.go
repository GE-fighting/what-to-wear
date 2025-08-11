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

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	common.Success(c, user, "User created successfully")
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, errors.ErrInvalidRequest("无效的登录请求", err.Error()))
		return
	}

	token, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		common.Error(c, err)
		return
	}

	common.Success(c, gin.H{"token": token}, "Login successful")
}
