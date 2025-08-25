package services

import (
	"context"
	"time"
	"what-to-wear/server/api/dto"
	"what-to-wear/server/api/errors"
	"what-to-wear/server/logger"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
	"what-to-wear/server/utils"
)

type AuthService interface {
	// 用户注册
	Register(ctx context.Context, req *dto.RegisterDTO) (*models.User, error)

	// 用户登录
	Login(ctx context.Context, username, password string) (string, error)

	// 验证用户
	ValidateUser(ctx context.Context, userID uint) (*models.User, error)

	// 刷新Token
	RefreshToken(ctx context.Context, userID uint) (string, error)
}

// authService 认证服务实现
type authService struct {
	userRepo repositories.UserRepository
}

// NewAuthService 创建认证服务实例
func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Register 用户注册
func (s *authService) Register(ctx context.Context, req *dto.RegisterDTO) (*models.User, error) {
	// 检查用户名是否已存在
	exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.NewInternalError("failed to check username existence", err.Error())
	}
	if exists {
		return nil, errors.NewInternalError("username already exists")
	}
	// 检查邮箱是否已存在
	exists, err = s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NewInternalError("failed to check email existence", err.Error())
	}
	if exists {
		return nil, errors.NewInternalError("email already exists")
	}
	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.NewInternalError("failed to hash password", err.Error())
	}

	// 处理生日字段
	var birthDate *time.Time
	if req.BirthDate != nil {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate.String()); err == nil {
			birthDate = &parsed
		}
	}

	// 设置昵称默认值
	nickname := req.Nickname
	if nickname == "" {
		nickname = req.Username
	}

	// 创建用户
	user := &models.User{
		Username:  req.Username,
		Password:  hashedPassword,
		Email:     req.Email,
		Nickname:  nickname,
		Gender:    req.Gender,
		BirthDate: birthDate,
		Height:    req.Height,
		Weight:    req.Weight,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, errors.NewInternalError("failed to create user")
	}
	// 清除密码字段，不返回给客户端
	user.Password = ""
	return user, nil
}

// Login 用户登录
func (s *authService) Login(ctx context.Context, username, password string) (string, error) {
	log := logger.GetLogger()
	log.Info("User login attempt", logger.Fields{
		"username": username,
	})

	// 查找用户
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		log.Warn("Login failed: user not found", logger.Fields{
			"username": username,
		})
		return "", errors.NewInternalError("invalid username or password")
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		log.Warn("Login failed: invalid password", logger.Fields{
			"username": username,
			"user_id":  user.ID,
		})
		return "", errors.NewInternalError("invalid username or password")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		log.ErrorWithErr(err, "Failed to generate token", logger.Fields{
			"username": username,
			"user_id":  user.ID,
		})
		return "", errors.NewInternalError("failed to generate token")
	}

	log.Info("User login successful", logger.Fields{
		"username": username,
		"user_id":  user.ID,
	})

	return token, nil
}

// ValidateUser 验证用户
func (s *authService) ValidateUser(ctx context.Context, userID uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.NewInternalError("user not found")
	}

	// 清除密码字段
	user.Password = ""
	return user, nil
}

// RefreshToken 刷新Token
func (s *authService) RefreshToken(ctx context.Context, userID uint) (string, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return "", errors.NewInternalError("user not found")
	}

	// 生成新的JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.NewInternalError("failed to generate token")
	}

	return token, nil
}
