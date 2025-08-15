package dto

// WeatherRequest 天气查询请求
type WeatherRequest struct {
	City      string  `json:"city" form:"city"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
	Units     string  `json:"units" form:"units"` // metric, imperial, kelvin
}

// WeatherResponse 天气响应
type WeatherResponse struct {
	Location    LocationInfo    `json:"location"`
	Current     CurrentWeather  `json:"current"`
	Forecast    []ForecastDay   `json:"forecast,omitempty"`
	Alerts      []WeatherAlert  `json:"alerts,omitempty"`
	LastUpdated string          `json:"last_updated"`
}

// LocationInfo 位置信息
type LocationInfo struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

// CurrentWeather 当前天气
type CurrentWeather struct {
	Temperature   float64 `json:"temperature"`
	FeelsLike     float64 `json:"feels_like"`
	Humidity      int     `json:"humidity"`
	Pressure      float64 `json:"pressure"`
	Visibility    float64 `json:"visibility"`
	UVIndex       float64 `json:"uv_index"`
	WindSpeed     float64 `json:"wind_speed"`
	WindDirection int     `json:"wind_direction"`
	WindGust      float64 `json:"wind_gust,omitempty"`
	Condition     string  `json:"condition"`
	Description   string  `json:"description"`
	Icon          string  `json:"icon"`
	IsDay         bool    `json:"is_day"`
}

// ForecastDay 预报天气
type ForecastDay struct {
	Date        string      `json:"date"`
	MaxTemp     float64     `json:"max_temp"`
	MinTemp     float64     `json:"min_temp"`
	Condition   string      `json:"condition"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
	Humidity    int         `json:"humidity"`
	WindSpeed   float64     `json:"wind_speed"`
	Precipitation float64   `json:"precipitation"`
	Hours       []HourlyWeather `json:"hours,omitempty"`
}

// HourlyWeather 小时天气
type HourlyWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	Icon        string  `json:"icon"`
	WindSpeed   float64 `json:"wind_speed"`
	Humidity    int     `json:"humidity"`
	Precipitation float64 `json:"precipitation"`
}

// WeatherAlert 天气预警
type WeatherAlert struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"` // minor, moderate, severe, extreme
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Areas       []string `json:"areas"`
}

// ClothingRecommendationRequest 基于天气的穿衣建议请求
type ClothingRecommendationRequest struct {
	Weather   WeatherResponse `json:"weather"`
	Occasion  string         `json:"occasion,omitempty"`
	UserID    uint           `json:"user_id"`
	Gender    string         `json:"gender,omitempty"`
	Style     string         `json:"style,omitempty"`
}

// ClothingRecommendationResponse 穿衣建议响应
type ClothingRecommendationResponse struct {
	Recommendations []ClothingRecommendation `json:"recommendations"`
	Weather         WeatherSummary          `json:"weather"`
	Tips            []string                `json:"tips"`
	Confidence      float64                 `json:"confidence"`
}

// ClothingRecommendation 单项穿衣建议
type ClothingRecommendation struct {
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	Material    string  `json:"material,omitempty"`
	Color       string  `json:"color,omitempty"`
	Layer       int     `json:"layer"`
	Priority    int     `json:"priority"`
	Reason      string  `json:"reason"`
	Confidence  float64 `json:"confidence"`
}

// WeatherSummary 天气摘要
type WeatherSummary struct {
	Temperature   float64 `json:"temperature"`
	FeelsLike     float64 `json:"feels_like"`
	Condition     string  `json:"condition"`
	Humidity      int     `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	Precipitation float64 `json:"precipitation"`
	UVIndex       float64 `json:"uv_index"`
}