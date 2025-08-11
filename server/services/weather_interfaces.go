package services

import (
	"what-to-wear/server/models"
)

// WeatherService 天气服务接口
type WeatherService interface {
	// 获取当前天气
	GetCurrentWeather(location string) (*models.Weather, error)

	// 获取天气预报
	GetWeatherForecast(location string, days int) ([]*models.WeatherForecast, error)

	// 更新天气数据
	UpdateWeatherData(location string) error
}
