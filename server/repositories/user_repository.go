package repositories

import (
	"context"
	"errors"
	"what-to-wear/server/logger"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	// 创建用户
	Create(ctx context.Context, user *models.User) error

	// 根据用户名查找用户
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	// 根据邮箱查找用户
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// 根据ID查找用户
	GetByID(ctx context.Context, id uint) (*models.User, error)

	// 更新用户信息
	Update(ctx context.Context, user *models.User) error

	// 删除用户
	Delete(ctx context.Context, id uint) error

	// 检查用户名是否存在
	ExistsByUsername(ctx context.Context, username string) (bool, error)

	// 检查邮箱是否存在
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// userRepository 用户仓储实现
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create 创建用户
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	log := logger.GetLogger()

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.ErrorWithErr(err, "Failed to create user in database", logger.Fields{
			"username": user.Username,
			"email":    user.Email,
		})
		return err
	}

	log.Debug("User created successfully in database", logger.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	})
	return nil
}

// GetByUsername 根据用户名查找用户
func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱查找用户
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByID 根据ID查找用户
func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ExistsByUsername 检查用户名是否存在
func (r *userRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByEmail 检查邮箱是否存在
func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
