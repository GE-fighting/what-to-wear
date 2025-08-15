package dto

// CreateOutfitRequest 创建穿搭记录请求
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

// OutfitListResponse 穿搭列表响应
type OutfitListResponse struct {
	Outfits    []OutfitSummary `json:"outfits"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}

// OutfitSummary 穿搭摘要信息
type OutfitSummary struct {
	ID          uint    `json:"id"`
	Date        string  `json:"date"`
	Temperature float64 `json:"temperature"`
	Weather     string  `json:"weather"`
	PhotoURL    string  `json:"photo_url"`
	Notes       string  `json:"notes"`
}

// RateOutfitRequest 评价穿搭请求
type RateOutfitRequest struct {
	Rating int    `json:"rating" binding:"required,min=1,max=5"`
	Notes  string `json:"notes"`
}

// OutfitResponse 穿搭响应
type OutfitResponse struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	Date        string  `json:"date"`
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
	Rating      *int    `json:"rating,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// OutfitRecommendationResponse 穿搭推荐响应
type OutfitRecommendationResponse struct {
	ID               uint                    `json:"id"`
	RecommendedItems []RecommendedClothingItem `json:"recommended_items"`
	Weather          WeatherInfo             `json:"weather"`
	Occasion         string                  `json:"occasion"`
	Confidence       float64                 `json:"confidence"`
	Reason           string                  `json:"reason"`
	CreatedAt        string                  `json:"created_at"`
}

// RecommendedClothingItem 推荐的衣物单品
type RecommendedClothingItem struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Color        string  `json:"color"`
	CategoryName string  `json:"category_name"`
	ImageURL     string  `json:"image_url"`
	Role         string  `json:"role"` // top, bottom, shoes, accessory
	Layer        int     `json:"layer"`
	Confidence   float64 `json:"confidence"`
}

// WeatherInfo 天气信息
type WeatherInfo struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Description string  `json:"description"`
}
