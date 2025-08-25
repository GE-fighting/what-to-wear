package models

import (
	"time"

	"what-to-wear/server/api"

	"gorm.io/gorm"
)

// Outfit 穿搭记录模型
type Outfit struct {
	gorm.Model
	UserID      uint              `json:"user_id" gorm:"not null;index"`
	Name        string            `json:"name" gorm:"not null"`
	Date        time.Time         `json:"date" gorm:"not null"`
	Temperature *float64          `json:"temperature"`
	Weather     *api.WeatherType  `json:"weather"`
	Occasion    string            `json:"occasion"`
	Location    string            `json:"location"`
	Rating      *api.OutfitRating `json:"rating"`
	RatingNotes string            `json:"rating_notes"`
	Notes       string            `json:"notes"`
	Tags        []string          `json:"tags" gorm:"type:json"`
	IsPublic    bool              `json:"is_public" gorm:"default:false"`
}

// OutfitRecommendation 穿搭推荐模型
type OutfitRecommendation struct {
	gorm.Model
	UserID           uint             `json:"user_id" gorm:"not null;index"`
	RecommendedItems []uint           `json:"recommended_items" gorm:"type:json"` // 推荐的衣物ID列表
	Weather          *api.WeatherType `json:"weather"`
	Temperature      *float64         `json:"temperature"`
	Occasion         string           `json:"occasion"`
	Confidence       float64          `json:"confidence"`                         // 推荐置信度
	Reason           string           `json:"reason"`                             // 推荐理由
	AlternativeItems [][]uint         `json:"alternative_items" gorm:"type:json"` // 替代选择的衣物ID列表
}
