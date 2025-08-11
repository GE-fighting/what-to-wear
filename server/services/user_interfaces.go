package services

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

// UserService 用户服务接口
type UserService interface {
	// 获取用户资料
	GetProfile(userID uint) (*models.User, error)

	// 更新用户资料
	UpdateProfile(userID uint, req *dto.UpdateProfileRequest) (*models.User, error)

	// 更改密码
	ChangePassword(userID uint, oldPassword, newPassword string) error

	// 删除用户
	DeleteUser(userID uint) error
}
