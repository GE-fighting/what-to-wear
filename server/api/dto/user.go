package dto

import (
	"time"
	"what-to-wear/server/api"
)

// RegisterDTO 用户注册DTO
type RegisterDTO struct {
	Username  string     `json:"username" binding:"required,min=3,max=20"`
	Password  string     `json:"password" binding:"required,min=6"`
	Email     string     `json:"email" binding:"required,email"`
	Nickname  string     `json:"nickname"`
	Gender    api.Gender `json:"gender"`
	BirthDate *time.Time `json:"birth_date"`
	Height    *int       `json:"height"`
	Weight    *int       `json:"weight"`
}

// LoginDTO 用户登录DTO
type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应DTO
type LoginResponseDTO struct {
	Token string          `json:"token"`
	User  *UserProfileDTO `json:"user"`
}

// UserProfileDTO 用户资料信息DTO
type UserProfileDTO struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Nickname  string     `json:"nickname"`
	Gender    api.Gender `json:"gender"`
	BirthDate *time.Time `json:"birth_date"`
	Height    *int       `json:"height"`
	Weight    *int       `json:"weight"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// UpdateProfileDTO 更新用户资料DTO
type UpdateProfileDTO struct {
	Nickname  *string     `json:"nickname"`
	Email     *string     `json:"email"`
	Gender    *api.Gender `json:"gender"`
	BirthDate *time.Time  `json:"birth_date"`
	Height    *int        `json:"height"`
	Weight    *int        `json:"weight"`
}

// ChangePasswordDTO 修改密码DTO
type ChangePasswordDTO struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// UserStatsDTO 用户统计DTO
type UserStatsDTO struct {
	TotalClothingItems int64     `json:"total_clothing_items"`
	TotalOutfits       int64     `json:"total_outfits"`
	TotalSpent         float64   `json:"total_spent"`
	AccountAge         int       `json:"account_age_days"`
	LastLoginAt        time.Time `json:"last_login_at,omitempty"`
	FavoriteColors     []string  `json:"favorite_colors"`
	FavoriteBrands     []string  `json:"favorite_brands"`
	MostWornItems      []uint    `json:"most_worn_items"`
}

// UserPreferencesDTO 用户偏好设置DTO
type UserPreferencesDTO struct {
	Language        string               `json:"language"`
	Timezone        string               `json:"timezone"`
	Currency        string               `json:"currency"`
	Notifications   NotificationPrefsDTO `json:"notifications"`
	PrivacySettings PrivacySettingsDTO   `json:"privacy_settings"`
	DisplaySettings DisplaySettingsDTO   `json:"display_settings"`
}

// NotificationPrefs 通知偏好设置DTO
type NotificationPrefsDTO struct {
	EmailNotifications   bool `json:"email_notifications"`
	PushNotifications    bool `json:"push_notifications"`
	MaintenanceReminders bool `json:"maintenance_reminders"`
	OutfitSuggestions    bool `json:"outfit_suggestions"`
	WeatherAlerts        bool `json:"weather_alerts"`
}

// PrivacySettingsDTO 隐私设置DTO
type PrivacySettingsDTO struct {
	ProfileVisibility string `json:"profile_visibility"` // public, friends, private
	ShowRealName      bool   `json:"show_real_name"`
	ShowLocation      bool   `json:"show_location"`
	AllowDataExport   bool   `json:"allow_data_export"`
}

// DisplaySettingsDTO 显示设置DTO
type DisplaySettingsDTO struct {
	Theme         string `json:"theme"`     // light, dark, auto
	GridSize      string `json:"grid_size"` // small, medium, large
	ShowPrices    bool   `json:"show_prices"`
	ShowWearCount bool   `json:"show_wear_count"`
	DefaultSortBy string `json:"default_sort_by"`
	ItemsPerPage  int    `json:"items_per_page"`
}
