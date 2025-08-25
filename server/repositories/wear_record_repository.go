package repositories

import (
	"context"
	"time"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// WearRecordRepository 穿着记录仓库接口
type WearRecordRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, record *models.WearRecord) error
	GetByID(ctx context.Context, id uint) (*models.WearRecord, error)
	GetByClothingItemID(ctx context.Context, clothingItemID uint, limit int) ([]models.WearRecord, error)
	Update(ctx context.Context, record *models.WearRecord) error
	Delete(ctx context.Context, id uint) error

	// 查询
	GetByUserID(ctx context.Context, userID uint, limit int) ([]models.WearRecord, error)
	GetByDateRange(ctx context.Context, userID uint, startDate, endDate string) ([]models.WearRecord, error)
	GetByOccasion(ctx context.Context, userID uint, occasion string) ([]models.WearRecord, error)
	GetByWeather(ctx context.Context, userID uint, weather string) ([]models.WearRecord, error)

	// 统计
	GetWearStats(ctx context.Context, clothingItemID uint) (map[string]interface{}, error)
	GetWearFrequency(ctx context.Context, userID uint) (map[string]int64, error)
	GetComfortRatings(ctx context.Context, userID uint) (map[uint]float64, error)
}

// wearRecordRepository 穿着记录仓库实现
type wearRecordRepository struct {
	db *gorm.DB
}

// NewWearRecordRepository 创建穿着记录仓库实例
func NewWearRecordRepository(db *gorm.DB) WearRecordRepository {
	return &wearRecordRepository{db: db}
}

// Create 创建穿着记录
func (r *wearRecordRepository) Create(ctx context.Context, record *models.WearRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

// GetByID 根据ID获取穿着记录
func (r *wearRecordRepository) GetByID(ctx context.Context, id uint) (*models.WearRecord, error) {
	var record models.WearRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetByClothingItemID 根据衣物ID获取穿着记录
func (r *wearRecordRepository) GetByClothingItemID(ctx context.Context, clothingItemID uint, limit int) ([]models.WearRecord, error) {
	var records []models.WearRecord
	query := r.db.WithContext(ctx).Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// Update 更新穿着记录
func (r *wearRecordRepository) Update(ctx context.Context, record *models.WearRecord) error {
	return r.db.WithContext(ctx).Save(record).Error
}

// Delete 删除穿着记录
func (r *wearRecordRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.WearRecord{}, id).Error
}

// GetByUserID 根据用户ID获取穿着记录
func (r *wearRecordRepository) GetByUserID(ctx context.Context, userID uint, limit int) ([]models.WearRecord, error) {
	var records []models.WearRecord
	query := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("wear_records.wear_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetByDateRange 根据日期范围获取穿着记录
func (r *wearRecordRepository) GetByDateRange(ctx context.Context, userID uint, startDate, endDate string) ([]models.WearRecord, error) {
	var records []models.WearRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.wear_date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("wear_records.wear_date DESC").
		Find(&records).Error
	return records, err
}

// GetByOccasion 根据场合获取穿着记录
func (r *wearRecordRepository) GetByOccasion(ctx context.Context, userID uint, occasion string) ([]models.WearRecord, error) {
	var records []models.WearRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.occasion = ?", userID, occasion).
		Order("wear_records.wear_date DESC").
		Find(&records).Error
	return records, err
}

// GetByWeather 根据天气获取穿着记录
func (r *wearRecordRepository) GetByWeather(ctx context.Context, userID uint, weather string) ([]models.WearRecord, error) {
	var records []models.WearRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.weather_condition = ?", userID, weather).
		Order("wear_records.wear_date DESC").
		Find(&records).Error
	return records, err
}

// GetWearStats 获取衣物穿着统计
func (r *wearRecordRepository) GetWearStats(ctx context.Context, clothingItemID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 基本统计
	var basicStats struct {
		TotalWears    int64      `json:"total_wears"`
		TotalHours    int64      `json:"total_hours"`
		AvgComfort    float64    `json:"avg_comfort"`
		AvgStyle      float64    `json:"avg_style"`
		LastWornDate  *time.Time `json:"last_worn_date"`
		FirstWornDate *time.Time `json:"first_worn_date"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Where("clothing_item_id = ?", clothingItemID).
		Select("COUNT(*) as total_wears, COALESCE(SUM(duration_hours), 0) as total_hours, COALESCE(AVG(comfort_rating), 0) as avg_comfort, COALESCE(AVG(style_rating), 0) as avg_style").
		Scan(&basicStats).Error

	if err != nil {
		return nil, err
	}

	// 获取最后穿着日期
	var lastRecord models.WearRecord
	err = r.db.WithContext(ctx).Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date DESC").
		First(&lastRecord).Error
	if err == nil {
		basicStats.LastWornDate = &lastRecord.WearDate
	}

	// 获取首次穿着日期
	var firstRecord models.WearRecord
	err = r.db.WithContext(ctx).Where("clothing_item_id = ?", clothingItemID).
		Order("wear_date ASC").
		First(&firstRecord).Error
	if err == nil {
		basicStats.FirstWornDate = &firstRecord.WearDate
	}

	stats["total_wears"] = basicStats.TotalWears
	stats["total_hours"] = basicStats.TotalHours
	stats["avg_comfort"] = basicStats.AvgComfort
	stats["avg_style"] = basicStats.AvgStyle
	stats["last_worn_date"] = basicStats.LastWornDate
	stats["first_worn_date"] = basicStats.FirstWornDate

	return stats, nil
}

// GetWearFrequency 获取穿着频率统计
func (r *wearRecordRepository) GetWearFrequency(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		Occasion string `json:"occasion"`
		Count    int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("occasion, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.occasion != ''", userID).
		Group("occasion").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	frequencyMap := make(map[string]int64)
	for _, result := range results {
		frequencyMap[result.Occasion] = result.Count
	}

	return frequencyMap, nil
}

// GetComfortRatings 获取舒适度评分统计
func (r *wearRecordRepository) GetComfortRatings(ctx context.Context, userID uint) (map[uint]float64, error) {
	var results []struct {
		ClothingItemID uint    `json:"clothing_item_id"`
		AvgComfort     float64 `json:"avg_comfort"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("clothing_item_id, COALESCE(AVG(comfort_rating), 0) as avg_comfort").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.comfort_rating > 0", userID).
		Group("clothing_item_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	comfortMap := make(map[uint]float64)
	for _, result := range results {
		comfortMap[result.ClothingItemID] = result.AvgComfort
	}

	return comfortMap, nil
}

// GetWeatherStats 获取天气统计
func (r *wearRecordRepository) GetWeatherStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		WeatherCondition string `json:"weather_condition"`
		Count            int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("weather_condition, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.weather_condition != ''", userID).
		Group("weather_condition").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	weatherMap := make(map[string]int64)
	for _, result := range results {
		weatherMap[result.WeatherCondition] = result.Count
	}

	return weatherMap, nil
}

// GetActivityStats 获取活动统计
func (r *wearRecordRepository) GetActivityStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		Activity string `json:"activity"`
		Count    int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("activity, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.activity != ''", userID).
		Group("activity").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	activityMap := make(map[string]int64)
	for _, result := range results {
		activityMap[result.Activity] = result.Count
	}

	return activityMap, nil
}

// GetMoodStats 获取心情统计
func (r *wearRecordRepository) GetMoodStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		Mood  string `json:"mood"`
		Count int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("mood, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.mood != ''", userID).
		Group("mood").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	moodMap := make(map[string]int64)
	for _, result := range results {
		moodMap[result.Mood] = result.Count
	}

	return moodMap, nil
}

// GetWearIntensityStats 获取穿着强度统计
func (r *wearRecordRepository) GetWearIntensityStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		WearIntensity string `json:"wear_intensity"`
		Count         int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("wear_intensity, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.wear_intensity != ''", userID).
		Group("wear_intensity").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	intensityMap := make(map[string]int64)
	for _, result := range results {
		intensityMap[result.WearIntensity] = result.Count
	}

	return intensityMap, nil
}

// GetRecentWearRecords 获取最近的穿着记录
func (r *wearRecordRepository) GetRecentWearRecords(ctx context.Context, userID uint, limit int) ([]models.WearRecord, error) {
	var records []models.WearRecord
	query := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("wear_records.wear_date DESC, wear_records.created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetHighRatedWears 获取高评分穿着记录
func (r *wearRecordRepository) GetHighRatedWears(ctx context.Context, userID uint, minRating int, limit int) ([]models.WearRecord, error) {
	var records []models.WearRecord
	query := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND (wear_records.comfort_rating >= ? OR wear_records.style_rating >= ?)", userID, minRating, minRating).
		Order("(wear_records.comfort_rating + wear_records.style_rating) DESC, wear_records.wear_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetWearsByTemperatureRange 根据温度范围获取穿着记录
func (r *wearRecordRepository) GetWearsByTemperatureRange(ctx context.Context, userID uint, minTemp, maxTemp float64) ([]models.WearRecord, error) {
	var records []models.WearRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND wear_records.temperature BETWEEN ? AND ?", userID, minTemp, maxTemp).
		Order("wear_records.wear_date DESC").
		Find(&records).Error
	return records, err
}

// GetTotalWearTime 获取用户总穿着时间
func (r *wearRecordRepository) GetTotalWearTime(ctx context.Context, userID uint) (int64, error) {
	var totalHours int64
	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(SUM(wear_records.duration_hours), 0)").
		Scan(&totalHours).Error
	return totalHours, err
}

// GetWearRecordsByMonth 获取按月份分组的穿着统计
func (r *wearRecordRepository) GetWearRecordsByMonth(ctx context.Context, userID uint, year int) (map[string]int64, error) {
	var results []struct {
		Month string `json:"month"`
		Count int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.WearRecord{}).
		Select("DATE_FORMAT(wear_date, '%Y-%m') as month, COUNT(*) as count").
		Joins("JOIN clothing_items ON wear_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND YEAR(wear_records.wear_date) = ?", userID, year).
		Group("DATE_FORMAT(wear_date, '%Y-%m')").
		Order("month").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	monthlyStats := make(map[string]int64)
	for _, result := range results {
		monthlyStats[result.Month] = result.Count
	}

	return monthlyStats, nil
}
