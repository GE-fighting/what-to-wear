package dto

// RegisterRequest 用户注册请求
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

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  *UserProfile `json:"user"`
}

// UserProfile 用户资料信息
type UserProfile struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Height    *int   `json:"height"`
	Weight    *int   `json:"weight"`
}
