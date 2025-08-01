package services

import (
	"errors"
	"time"
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
	// 检查用户名是否已存在
	exists, err := s.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		return nil, errors.New("failed to check username existence")
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	exists, err = s.userRepo.ExistsByEmail(req.Email)
	if err != nil {
		return nil, errors.New("failed to check email existence")
	}
	if exists {
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
		return nil, errors.New("failed to create user")
	}

	// 清除密码字段，不返回给客户端
	user.Password = ""
	return user, nil
}

// Login 用户登录
func (s *authService) Login(username, password string) (string, error) {
	// 查找用户
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

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
