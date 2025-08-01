package repositories

import (
	"what-to-wear/server/models"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	// 创建用户
	Create(user *models.User) error
	
	// 根据用户名查找用户
	GetByUsername(username string) (*models.User, error)
	
	// 根据邮箱查找用户
	GetByEmail(email string) (*models.User, error)
	
	// 根据ID查找用户
	GetByID(id uint) (*models.User, error)
	
	// 更新用户信息
	Update(user *models.User) error
	
	// 删除用户
	Delete(id uint) error
	
	// 检查用户名是否存在
	ExistsByUsername(username string) (bool, error)
	
	// 检查邮箱是否存在
	ExistsByEmail(email string) (bool, error)
}

// WeatherRepository 天气数据访问接口
type WeatherRepository interface {
	// 获取当前天气
	GetCurrentWeather(location string) (*models.Weather, error)
	
	// 获取天气预报
	GetWeatherForecast(location string, days int) ([]*models.Weather, error)
	
	// 保存天气数据
	SaveWeatherData(weather *models.Weather) error
}

// OutfitRepository 穿搭数据访问接口
type OutfitRepository interface {
	// 创建穿搭记录
	Create(outfit *models.Outfit) error
	
	// 根据用户ID获取穿搭历史
	GetByUserID(userID uint, limit, offset int) ([]*models.Outfit, error)
	
	// 根据ID获取穿搭记录
	GetByID(id uint) (*models.Outfit, error)
	
	// 更新穿搭记录
	Update(outfit *models.Outfit) error
	
	// 删除穿搭记录
	Delete(id uint) error
}
