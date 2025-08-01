package services

import (
	"what-to-wear/server/models"
)

// AuthService 认证服务接口
type AuthService interface {
	// 用户注册
	Register(req *RegisterRequest) (*models.User, error)
	
	// 用户登录
	Login(username, password string) (string, error)
	
	// 验证用户
	ValidateUser(userID uint) (*models.User, error)
	
	// 刷新Token
	RefreshToken(userID uint) (string, error)
}

// UserService 用户服务接口
type UserService interface {
	// 获取用户资料
	GetProfile(userID uint) (*models.User, error)
	
	// 更新用户资料
	UpdateProfile(userID uint, req *UpdateProfileRequest) (*models.User, error)
	
	// 更改密码
	ChangePassword(userID uint, oldPassword, newPassword string) error
	
	// 删除用户
	DeleteUser(userID uint) error
}

// WeatherService 天气服务接口
type WeatherService interface {
	// 获取当前天气
	GetCurrentWeather(location string) (*models.Weather, error)
	
	// 获取天气预报
	GetWeatherForecast(location string, days int) ([]*models.WeatherForecast, error)
	
	// 更新天气数据
	UpdateWeatherData(location string) error
}

// OutfitService 穿搭服务接口
type OutfitService interface {
	// 创建穿搭记录
	CreateOutfit(userID uint, req *CreateOutfitRequest) (*models.Outfit, error)
	
	// 获取用户穿搭历史
	GetUserOutfits(userID uint, page, pageSize int) ([]*models.Outfit, int64, error)
	
	// 获取穿搭推荐
	GetOutfitRecommendation(userID uint, weather *models.Weather) (*models.OutfitRecommendation, error)
	
	// 评价穿搭
	RateOutfit(userID, outfitID uint, rating int, notes string) error
}

// 请求结构体定义
type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=20"`
	Password  string `json:"password" binding:"required,min=6"`
	Email     string `json:"email" binding:"required,email"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Height    *int   `json:"height"`
	Weight    *int   `json:"weight"`
}

type UpdateProfileRequest struct {
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Height    *int   `json:"height"`
	Weight    *int   `json:"weight"`
}

type CreateOutfitRequest struct {
	Date        string  `json:"date" binding:"required"`
	Temperature float64 `json:"temperature"`
	Weather     string  `json:"weather"`
	TopType     string  `json:"top_type"`
	TopColor    string  `json:"top_color"`
	BottomType  string  `json:"bottom_type"`
	BottomColor string  `json:"bottom_color"`
	ShoesType   string  `json:"shoes_type"`
	ShoesColor  string  `json:"shoes_color"`
	Accessories string  `json:"accessories"`
	Notes       string  `json:"notes"`
	PhotoURL    string  `json:"photo_url"`
}
