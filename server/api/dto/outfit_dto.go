package dto

import (
	"time"
	"what-to-wear/server/api"
)

// CreateOutfitDTO 创建穿搭记录DTO
type CreateOutfitDTO struct {
	Name        string           `json:"name" binding:"required"`
	Date        time.Time        `json:"date" binding:"required"`
	Temperature *float64         `json:"temperature"`
	Weather     *api.WeatherType `json:"weather"`
	Occasion    string           `json:"occasion"`
	Location    string           `json:"location"`
	Notes       string           `json:"notes"`
	ClothingIDs []uint           `json:"clothing_ids" binding:"required,min=1"`
	Tags        []string         `json:"tags"`
	IsPublic    bool             `json:"is_public"`
}

// UpdateOutfitDTO 更新穿搭DTO
type UpdateOutfitDTO struct {
	Name        *string          `json:"name"`
	Date        *time.Time       `json:"date"`
	Temperature *float64         `json:"temperature"`
	Weather     *api.WeatherType `json:"weather"`
	Occasion    *string          `json:"occasion"`
	Location    *string          `json:"location"`
	Notes       *string          `json:"notes"`
	ClothingIDs []uint           `json:"clothing_ids"`
	Tags        []string         `json:"tags"`
	IsPublic    *bool            `json:"is_public"`
}

// OutfitDTO 穿搭DTO
type OutfitDTO struct {
	ID            uint                 `json:"id"`
	UserID        uint                 `json:"user_id"`
	Name          string               `json:"name"`
	Date          time.Time            `json:"date"`
	Temperature   *float64             `json:"temperature"`
	Weather       *api.WeatherType     `json:"weather"`
	Occasion      string               `json:"occasion"`
	Location      string               `json:"location"`
	Notes         string               `json:"notes"`
	Tags          []string             `json:"tags"`
	IsPublic      bool                 `json:"is_public"`
	ClothingItems []OutfitClothingItem `json:"clothing_items"`
	Attachments   []AttachmentDTO      `json:"attachments"`
	Rating        *api.OutfitRating    `json:"rating,omitempty"`
	RatingNotes   string               `json:"rating_notes,omitempty"`
	WearCount     int                  `json:"wear_count"`
	LastWornDate  *time.Time           `json:"last_worn_date"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

// OutfitClothingItem 穿搭中的衣物单品
type OutfitClothingItem struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	Color        string `json:"color"`
	CategoryName string `json:"category_name"`
	ImageURL     string `json:"image_url"`
	Layer        int    `json:"layer"`    // 穿着层次
	Position     string `json:"position"` // 位置：top, bottom, shoes, accessory
}

// OutfitListRequest 穿搭列表请求DTO
type OutfitListDTO struct {
	Weather  *api.WeatherType  `form:"weather"`
	Occasion string            `form:"occasion"`
	Season   string            `form:"season"`
	DateFrom *time.Time        `form:"date_from"`
	DateTo   *time.Time        `form:"date_to"`
	Rating   *api.OutfitRating `form:"rating"`
	IsPublic *bool             `form:"is_public"`
	SearchRequest
}

// OutfitSummary 穿搭摘要信息DTO
type OutfitSummaryDTO struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	Date         time.Time         `json:"date"`
	Temperature  *float64          `json:"temperature"`
	Weather      *api.WeatherType  `json:"weather"`
	Occasion     string            `json:"occasion"`
	ImageURL     string            `json:"image_url"`
	ItemCount    int               `json:"item_count"`
	Rating       *api.OutfitRating `json:"rating,omitempty"`
	WearCount    int               `json:"wear_count"`
	LastWornDate *time.Time        `json:"last_worn_date"`
}

// RateOutfitDTO 评价穿搭DTO
type RateOutfitDTO struct {
	Rating api.OutfitRating `json:"rating" binding:"required,min=1,max=5"`
	Notes  string           `json:"notes"`
}

// OutfitRecommendationDTO 穿搭推荐DTO
type OutfitRecommendationDTO struct {
	ID               uint                        `json:"id"`
	RecommendedItems []RecommendedClothingItem   `json:"recommended_items"`
	Weather          WeatherInfo                 `json:"weather"`
	Occasion         string                      `json:"occasion"`
	Confidence       float64                     `json:"confidence"`
	Reason           string                      `json:"reason"`
	AlternativeItems [][]RecommendedClothingItem `json:"alternative_items"` // 替代选择
	CreatedAt        time.Time                   `json:"created_at"`
}

// Outfit 穿搭记录
type Outfit struct {
	ID          uint             `json:"id"`
	UserID      uint             `json:"user_id"`
	Name        string           `json:"name"`
	Date        time.Time        `json:"date"`
	Temperature *float64         `json:"temperature,omitempty"`
	Weather     *api.WeatherType `json:"weather,omitempty"`
	Occasion    string           `json:"occasion"`
	Location    string           `json:"location"`
	Notes       string           `json:"notes"`
	IsPublic    bool             `json:"is_public"`
	ClothingIDs []uint           `json:"clothing_ids"`
	Tags        []string         `json:"tags"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

// OutfitRecommendation 穿搭推荐
type OutfitRecommendation struct {
	ID               uint                      `json:"id"`
	RecommendedItems []RecommendedClothingItem `json:"recommended_items"`
	Weather          WeatherInfo               `json:"weather"`
	Occasion         string                    `json:"occasion"`
	Confidence       float64                   `json:"confidence"`
	Reason           string                    `json:"reason"`
	CreatedAt        time.Time                 `json:"created_at"`
}

// RecommendedClothingItem 推荐的衣物单品
type RecommendedClothingItem struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Color        string  `json:"color"`
	CategoryName string  `json:"category_name"`
	ImageURL     string  `json:"image_url"`
	Position     string  `json:"position"`   // top, bottom, shoes, accessory
	Layer        int     `json:"layer"`      // 穿着层次
	Confidence   float64 `json:"confidence"` // 推荐置信度
	Reason       string  `json:"reason"`     // 推荐理由
}

// WeatherInfo 天气信息
type WeatherInfo struct {
	Temperature float64         `json:"temperature"`
	Condition   api.WeatherType `json:"condition"`
	Humidity    *float64        `json:"humidity,omitempty"`
	WindSpeed   *float64        `json:"wind_speed,omitempty"`
	Description string          `json:"description"`
}

// OutfitStatsDTO 穿搭统计DTO
type OutfitStatsDTO struct {
	TotalOutfits     int64                     `json:"total_outfits"`
	FavoriteWeather  map[api.WeatherType]int64 `json:"favorite_weather"`
	FavoriteOccasion map[string]int64          `json:"favorite_occasion"`
	AverageRating    float64                   `json:"average_rating"`
	MostWornOutfits  []OutfitSummaryDTO        `json:"most_worn_outfits"`
	RecentOutfits    []OutfitSummaryDTO        `json:"recent_outfits"`
}
