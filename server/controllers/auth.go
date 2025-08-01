package controllers

import (
	"what-to-wear/server/common"
	"what-to-wear/server/logger"
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
	requestID := c.GetString("request_id")

	var req services.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid registration request", logger.Fields{
			"request_id": requestID,
			"error":      err.Error(),
		})
		common.HandleError(c, common.ErrInvalidRequest)
		return
	}

	logger.Info("Processing user registration", logger.Fields{
		"request_id": requestID,
		"username":   req.Username,
		"email":      req.Email,
	})

	user, err := ac.authService.Register(&req)
	if err != nil {
		logger.Error("User registration failed", logger.Fields{
			"request_id": requestID,
			"username":   req.Username,
			"error":      err.Error(),
		})
		common.HandleError(c, err)
		return
	}

	logger.Info("User registration successful", logger.Fields{
		"request_id": requestID,
		"user_id":    user.ID,
		"username":   user.Username,
	})

	common.CreatedResponse(c, user, "User created successfully")
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleError(c, common.ErrInvalidRequest)
		return
	}

	token, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		common.HandleError(c, err)
		return
	}

	common.SuccessResponse(c, gin.H{"token": token}, "Login successful")
}
