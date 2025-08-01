package models

import (
	"time"

	"gorm.io/gorm"
)

// Weather 天气模型
type Weather struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Location    string         `json:"location" gorm:"not null"`
	Temperature float64        `json:"temperature" gorm:"not null"`
	Condition   string         `json:"condition" gorm:"not null"`
	Humidity    int            `json:"humidity"`
	WindSpeed   float64        `json:"wind_speed"`
	Pressure    float64        `json:"pressure"`
	Visibility  float64        `json:"visibility"`
	UVIndex     int            `json:"uv_index"`
	Date        time.Time      `json:"date" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// WeatherForecast 天气预报模型
type WeatherForecast struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Location    string         `json:"location" gorm:"not null"`
	Date        time.Time      `json:"date" gorm:"not null"`
	MaxTemp     float64        `json:"max_temp" gorm:"not null"`
	MinTemp     float64        `json:"min_temp" gorm:"not null"`
	Condition   string         `json:"condition" gorm:"not null"`
	Humidity    int            `json:"humidity"`
	WindSpeed   float64        `json:"wind_speed"`
	Pressure    float64        `json:"pressure"`
	UVIndex     int            `json:"uv_index"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
