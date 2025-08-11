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
