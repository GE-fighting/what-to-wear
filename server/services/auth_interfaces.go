package services

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

// AuthService 认证服务接口
type AuthService interface {
	// 用户注册
	Register(req *dto.RegisterRequest) (*models.User, error)

	// 用户登录
	Login(username, password string) (string, error)

	// 验证用户
	ValidateUser(userID uint) (*models.User, error)

	// 刷新Token
	RefreshToken(userID uint) (string, error)
}
