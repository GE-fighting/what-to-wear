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
