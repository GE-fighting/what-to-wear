package dto

// UpdateProfileRequest 更新用户资料请求
type UpdateProfileRequest struct {
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Height    *int   `json:"height"`
	Weight    *int   `json:"weight"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// UserStatsResponse 用户统计响应
type UserStatsResponse struct {
	TotalClothingItems int64   `json:"total_clothing_items"`
	TotalOutfits       int64   `json:"total_outfits"`
	TotalSpent         float64 `json:"total_spent"`
	AccountAge         int     `json:"account_age_days"`
	LastLoginAt        string  `json:"last_login_at,omitempty"`
}

// UserPreferencesRequest 用户偏好设置请求
type UserPreferencesRequest struct {
	Language         string            `json:"language"`
	Timezone         string            `json:"timezone"`
	Currency         string            `json:"currency"`
	Notifications    NotificationPrefs `json:"notifications"`
	PrivacySettings  PrivacySettings   `json:"privacy_settings"`
	DisplaySettings  DisplaySettings   `json:"display_settings"`
}

// NotificationPrefs 通知偏好设置
type NotificationPrefs struct {
	EmailNotifications bool `json:"email_notifications"`
	PushNotifications  bool `json:"push_notifications"`
	MaintenanceReminders bool `json:"maintenance_reminders"`
	OutfitSuggestions   bool `json:"outfit_suggestions"`
	WeatherAlerts       bool `json:"weather_alerts"`
}

// PrivacySettings 隐私设置
type PrivacySettings struct {
	ProfileVisibility string `json:"profile_visibility"` // public, friends, private
	ShowRealName      bool   `json:"show_real_name"`
	ShowLocation      bool   `json:"show_location"`
	AllowDataExport   bool   `json:"allow_data_export"`
}

// DisplaySettings 显示设置
type DisplaySettings struct {
	Theme           string `json:"theme"`           // light, dark, auto
	GridSize        string `json:"grid_size"`       // small, medium, large
	ShowPrices      bool   `json:"show_prices"`
	ShowWearCount   bool   `json:"show_wear_count"`
	DefaultSortBy   string `json:"default_sort_by"`
	ItemsPerPage    int    `json:"items_per_page"`
}
