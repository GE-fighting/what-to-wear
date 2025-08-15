package models

import (
	"time"

	"gorm.io/gorm"
)

// WearRecord 穿着记录模型
type WearRecord struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	ClothingItemID   uint           `json:"clothing_item_id" gorm:"not null;index"`
	WearDate         time.Time      `json:"wear_date" gorm:"not null"`
	DurationHours    int            `json:"duration_hours" gorm:"default:8"` // 穿着时长（小时）
	Occasion         string         `json:"occasion"`                        // 场合
	WeatherCondition string         `json:"weather_condition"`               // 天气状况
	Temperature      float64        `json:"temperature"`                     // 温度
	Activity         string         `json:"activity"`                        // 活动类型
	ComfortRating    int            `json:"comfort_rating" gorm:"default:5"` // 舒适度评分 (1-10)
	StyleRating      int            `json:"style_rating" gorm:"default:5"`   // 造型评分 (1-10)
	Notes            string         `json:"notes"`
	Photos           []string       `json:"photos" gorm:"type:json"`         // 穿着照片
	Location         string         `json:"location"`                        // 穿着地点
	Companions       []string       `json:"companions" gorm:"type:json"`     // 同行人员
	Mood             string         `json:"mood"`                            // 心情
	WearIntensity    string         `json:"wear_intensity"`                  // 穿着强度 (light, normal, heavy)
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (WearRecord) TableName() string {
	return "wear_records"
}

// WearIntensity 穿着强度枚举
type WearIntensity string

const (
	WearIntensityLight  WearIntensity = "light"  // 轻度穿着
	WearIntensityNormal WearIntensity = "normal" // 正常穿着
	WearIntensityHeavy  WearIntensity = "heavy"  // 重度穿着
)

// GetWearDamage 根据穿着强度计算磨损值
func (w *WearRecord) GetWearDamage() float64 {
	intensityDamage := map[WearIntensity]float64{
		WearIntensityLight:  0.5,
		WearIntensityNormal: 1.0,
		WearIntensityHeavy:  2.0,
	}
	
	damage := intensityDamage[WearIntensity(w.WearIntensity)]
	
	// 根据穿着时长调整磨损
	if w.DurationHours > 8 {
		damage *= 1.2
	} else if w.DurationHours < 4 {
		damage *= 0.8
	}
	
	return damage
}

// IsValidWearIntensity 检查穿着强度是否有效
func IsValidWearIntensity(intensity string) bool {
	validIntensities := []WearIntensity{
		WearIntensityLight, WearIntensityNormal, WearIntensityHeavy,
	}
	
	for _, validIntensity := range validIntensities {
		if WearIntensity(intensity) == validIntensity {
			return true
		}
	}
	return false
}

// GetWearRecordsByItem 获取指定衣物的穿着记录
func GetWearRecordsByItem(db *gorm.DB, clothingItemID uint, limit int) ([]WearRecord, error) {
	var records []WearRecord
	query := db.Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date DESC")
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	err := query.Find(&records).Error
	return records, err
}

// GetWearStatsByItem 获取指定衣物的穿着统计
func GetWearStatsByItem(db *gorm.DB, clothingItemID uint) (map[string]interface{}, error) {
	var stats struct {
		TotalWears    int64   `json:"total_wears"`
		TotalHours    int64   `json:"total_hours"`
		AvgComfort    float64 `json:"avg_comfort"`
		AvgStyle      float64 `json:"avg_style"`
		LastWornDate  *time.Time `json:"last_worn_date"`
		FirstWornDate *time.Time `json:"first_worn_date"`
	}
	
	// 获取基本统计
	err := db.Model(&WearRecord{}).
		Where("clothing_item_id = ?", clothingItemID).
		Select("COUNT(*) as total_wears, SUM(duration_hours) as total_hours, AVG(comfort_rating) as avg_comfort, AVG(style_rating) as avg_style").
		Scan(&stats).Error
	
	if err != nil {
		return nil, err
	}
	
	// 获取最后穿着日期
	var lastRecord WearRecord
	err = db.Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date DESC").
		First(&lastRecord).Error
	if err == nil {
		stats.LastWornDate = &lastRecord.WearDate
	}
	
	// 获取首次穿着日期
	var firstRecord WearRecord
	err = db.Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date ASC").
		First(&firstRecord).Error
	if err == nil {
		stats.FirstWornDate = &firstRecord.WearDate
	}
	
	result := map[string]interface{}{
		"total_wears":     stats.TotalWears,
		"total_hours":     stats.TotalHours,
		"avg_comfort":     stats.AvgComfort,
		"avg_style":       stats.AvgStyle,
		"last_worn_date":  stats.LastWornDate,
		"first_worn_date": stats.FirstWornDate,
	}
	
	return result, nil
}

// GetWearRecordsByDateRange 获取指定日期范围内的穿着记录
func GetWearRecordsByDateRange(db *gorm.DB, userID uint, startDate, endDate time.Time) ([]WearRecord, error) {
	var records []WearRecord
	err := db.Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.wear_date BETWEEN ? AND ?", 
			userID, startDate, endDate).
		Order("wear_records.wear_date DESC").
		Find(&records).Error
	
	return records, err
}