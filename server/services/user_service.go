package services

import (
	"context"
	"errors"
	"time"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
	"what-to-wear/server/utils"
)

type UserService interface {
	// 获取用户资料
	GetProfile(ctx context.Context, userId uint) (*models.User, error)

	// 更新用户资料
	UpdateProfile(ctx context.Context, userID uint, req *dto.UpdateProfileDTO) (*models.User, error)

	// 更改密码
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error

	// 删除用户
	DeleteUser(ctx context.Context, userID uint) error
}

// userService 用户服务实现
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetProfile 获取用户资料
func (s *userService) GetProfile(ctx context.Context, userID uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 清除密码字段
	user.Password = ""
	return user, nil
}

// UpdateProfile 更新用户资料
func (s *userService) UpdateProfile(ctx context.Context, userId uint, req *dto.UpdateProfileDTO) (*models.User, error) {
	// 获取现有用户信息
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 如果邮箱有变更，检查新邮箱是否已被使用
	if req.Email != nil && *req.Email != user.Email {
		exists, err := s.userRepo.ExistsByEmail(ctx, *req.Email)
		if err != nil {
			return nil, errors.New("failed to check email existence")
		}
		if exists {
			return nil, errors.New("email already exists")
		}
		user.Email = *req.Email
	}

	// 更新其他字段
	if req.Nickname != nil && *req.Nickname != user.Nickname {
		user.Nickname = *req.Nickname
	}

	if req.Gender != nil && *req.Gender != user.Gender {
		user.Gender = *req.Gender
	}

	if req.BirthDate != nil {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate.String()); err == nil {
			user.BirthDate = &parsed
		}
	}

	if req.Height != nil {
		user.Height = req.Height
	}

	if req.Weight != nil {
		user.Weight = req.Weight
	}

	// 保存更新
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, errors.New("failed to update user profile")
	}

	// 清除密码字段
	user.Password = ""
	return user, nil
}

// ChangePassword 更改密码
func (s *userService) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	// 获取用户信息
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// 验证旧密码
	if !utils.CheckPassword(oldPassword, user.Password) {
		return errors.New("invalid old password")
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	// 更新密码
	user.Password = hashedPassword
	if err := s.userRepo.Update(ctx, user); err != nil {
		return errors.New("failed to update password")
	}

	return nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, userID uint) error {
	// 检查用户是否存在
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// 删除用户
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
