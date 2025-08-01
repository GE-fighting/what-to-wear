package services

import (
	"errors"
	"time"
	"what-to-wear/server/logger"
	"what-to-wear/server/models"
	"what-to-wear/server/repositories"
	"what-to-wear/server/utils"
)

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
func (s *authService) Register(req *RegisterRequest) (*models.User, error) {
	logger.Info("User registration started", logger.Fields{
		"username": req.Username,
		"email":    req.Email,
	})

	// 检查用户名是否已存在
	exists, err := s.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		logger.ErrorWithErr(err, "Failed to check username existence", logger.Fields{
			"username": req.Username,
		})
		return nil, errors.New("failed to check username existence")
	}
	if exists {
		logger.Warn("Registration failed: username already exists", logger.Fields{
			"username": req.Username,
		})
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	exists, err = s.userRepo.ExistsByEmail(req.Email)
	if err != nil {
		logger.ErrorWithErr(err, "Failed to check email existence", logger.Fields{
			"email": req.Email,
		})
		return nil, errors.New("failed to check email existence")
	}
	if exists {
		logger.Warn("Registration failed: email already exists", logger.Fields{
			"email": req.Email,
		})
		return nil, errors.New("email already exists")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 处理生日字段
	var birthDate *time.Time
	if req.BirthDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
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

	if err := s.userRepo.Create(user); err != nil {
		logger.ErrorWithErr(err, "Failed to create user", logger.Fields{
			"username": req.Username,
			"email":    req.Email,
		})
		return nil, errors.New("failed to create user")
	}

	logger.Info("User registration completed successfully", logger.Fields{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

	// 清除密码字段，不返回给客户端
	user.Password = ""
	return user, nil
}

// Login 用户登录
func (s *authService) Login(username, password string) (string, error) {
	logger.Info("User login attempt", logger.Fields{
		"username": username,
	})

	// 查找用户
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		logger.Warn("Login failed: user not found", logger.Fields{
			"username": username,
		})
		return "", errors.New("invalid username or password")
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		logger.Warn("Login failed: invalid password", logger.Fields{
			"username": username,
			"user_id":  user.ID,
		})
		return "", errors.New("invalid username or password")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.ErrorWithErr(err, "Failed to generate token", logger.Fields{
			"username": username,
			"user_id":  user.ID,
		})
		return "", errors.New("failed to generate token")
	}

	logger.Info("User login successful", logger.Fields{
		"username": username,
		"user_id":  user.ID,
	})

	return token, nil
}

// ValidateUser 验证用户
func (s *authService) ValidateUser(userID uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 清除密码字段
	user.Password = ""
	return user, nil
}

// RefreshToken 刷新Token
func (s *authService) RefreshToken(userID uint) (string, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// 生成新的JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
