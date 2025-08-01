package models

import (
	"time"

	"gorm.io/gorm"
)

// Outfit 穿搭记录模型
type Outfit struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	Date        time.Time      `json:"date" gorm:"not null"`
	Temperature float64        `json:"temperature"`
	Weather     string         `json:"weather"`
	TopType     string         `json:"top_type"`
	TopColor    string         `json:"top_color"`
	BottomType  string         `json:"bottom_type"`
	BottomColor string         `json:"bottom_color"`
	ShoesType   string         `json:"shoes_type"`
	ShoesColor  string         `json:"shoes_color"`
	Accessories string         `json:"accessories"`
	Rating      int            `json:"rating" gorm:"default:0"`
	Notes       string         `json:"notes"`
	PhotoURL    string         `json:"photo_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// OutfitRecommendation 穿搭推荐模型
type OutfitRecommendation struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	Temperature float64        `json:"temperature" gorm:"not null"`
	Weather     string         `json:"weather" gorm:"not null"`
	TopType     string         `json:"top_type"`
	TopColor    string         `json:"top_color"`
	BottomType  string         `json:"bottom_type"`
	BottomColor string         `json:"bottom_color"`
	ShoesType   string         `json:"shoes_type"`
	ShoesColor  string         `json:"shoes_color"`
	Accessories string         `json:"accessories"`
	Reason      string         `json:"reason"`
	Confidence  float64        `json:"confidence" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
