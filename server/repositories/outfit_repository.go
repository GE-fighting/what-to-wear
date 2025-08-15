package repositories

import (
	"time"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// outfitRepository 穿搭仓库实现
type outfitRepository struct {
	db *gorm.DB
}

// NewOutfitRepository 创建穿搭仓库实例
func NewOutfitRepository(db *gorm.DB) OutfitRepository {
	return &outfitRepository{db: db}
}

// Create 创建穿搭记录
func (r *outfitRepository) Create(outfit *models.Outfit) error {
	return r.db.Create(outfit).Error
}

// GetByUserID 根据用户ID获取穿搭历史
func (r *outfitRepository) GetByUserID(userID uint, limit, offset int) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	query := r.db.Where("user_id = ?", userID).
		Order("date DESC, created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&outfits).Error
	return outfits, err
}

// GetByID 根据ID获取穿搭记录
func (r *outfitRepository) GetByID(id uint) (*models.Outfit, error) {
	var outfit models.Outfit
	err := r.db.First(&outfit, id).Error
	if err != nil {
		return nil, err
	}
	return &outfit, nil
}

// Update 更新穿搭记录
func (r *outfitRepository) Update(outfit *models.Outfit) error {
	return r.db.Save(outfit).Error
}

// Delete 删除穿搭记录
func (r *outfitRepository) Delete(id uint) error {
	return r.db.Delete(&models.Outfit{}, id).Error
}

// GetByDateRange 根据日期范围获取穿搭记录
func (r *outfitRepository) GetByDateRange(userID uint, startDate, endDate time.Time) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	err := r.db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("date DESC").
		Find(&outfits).Error
	return outfits, err
}

// GetByWeather 根据天气条件获取穿搭记录
func (r *outfitRepository) GetByWeather(userID uint, weather string) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	err := r.db.Where("user_id = ? AND weather = ?", userID, weather).
		Order("date DESC").
		Find(&outfits).Error
	return outfits, err
}

// GetByTemperatureRange 根据温度范围获取穿搭记录
func (r *outfitRepository) GetByTemperatureRange(userID uint, minTemp, maxTemp float64) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	err := r.db.Where("user_id = ? AND temperature BETWEEN ? AND ?", userID, minTemp, maxTemp).
		Order("date DESC").
		Find(&outfits).Error
	return outfits, err
}

// GetHighRatedOutfits 获取高评分穿搭
func (r *outfitRepository) GetHighRatedOutfits(userID uint, minRating int, limit int) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	query := r.db.Where("user_id = ? AND rating >= ?", userID, minRating).
		Order("rating DESC, date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&outfits).Error
	return outfits, err
}

// GetRecentOutfits 获取最近的穿搭记录
func (r *outfitRepository) GetRecentOutfits(userID uint, limit int) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	query := r.db.Where("user_id = ?", userID).
		Order("date DESC, created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&outfits).Error
	return outfits, err
}

// GetOutfitStats 获取穿搭统计信息
func (r *outfitRepository) GetOutfitStats(userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总穿搭数量
	var totalCount int64
	err := r.db.Model(&models.Outfit{}).Where("user_id = ?", userID).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	stats["total_count"] = totalCount

	// 平均评分
	var avgRating float64
	err = r.db.Model(&models.Outfit{}).
		Where("user_id = ? AND rating > 0", userID).
		Select("COALESCE(AVG(rating), 0)").
		Scan(&avgRating).Error
	if err != nil {
		return nil, err
	}
	stats["average_rating"] = avgRating

	// 最高评分
	var maxRating int
	err = r.db.Model(&models.Outfit{}).
		Where("user_id = ?", userID).
		Select("COALESCE(MAX(rating), 0)").
		Scan(&maxRating).Error
	if err != nil {
		return nil, err
	}
	stats["max_rating"] = maxRating

	// 本月穿搭数量
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	var monthlyCount int64
	err = r.db.Model(&models.Outfit{}).
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
		Count(&monthlyCount).Error
	if err != nil {
		return nil, err
	}
	stats["monthly_count"] = monthlyCount

	return stats, nil
}

// GetWeatherStats 获取天气统计
func (r *outfitRepository) GetWeatherStats(userID uint) (map[string]int64, error) {
	var results []struct {
		Weather string `json:"weather"`
		Count   int64  `json:"count"`
	}

	err := r.db.Model(&models.Outfit{}).
		Select("weather, COUNT(*) as count").
		Where("user_id = ? AND weather != ''", userID).
		Group("weather").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	weatherStats := make(map[string]int64)
	for _, result := range results {
		weatherStats[result.Weather] = result.Count
	}

	return weatherStats, nil
}

// GetColorCombinationStats 获取颜色搭配统计
func (r *outfitRepository) GetColorCombinationStats(userID uint) (map[string]int64, error) {
	var results []struct {
		ColorCombo string `json:"color_combo"`
		Count      int64  `json:"count"`
	}

	err := r.db.Model(&models.Outfit{}).
		Select("CONCAT(top_color, '-', bottom_color) as color_combo, COUNT(*) as count").
		Where("user_id = ? AND top_color != '' AND bottom_color != ''", userID).
		Group("top_color, bottom_color").
		Order("count DESC").
		Limit(10).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	colorStats := make(map[string]int64)
	for _, result := range results {
		colorStats[result.ColorCombo] = result.Count
	}

	return colorStats, nil
}

// GetOutfitsByMonth 获取按月份分组的穿搭统计
func (r *outfitRepository) GetOutfitsByMonth(userID uint, year int) (map[string]int64, error) {
	var results []struct {
		Month string `json:"month"`
		Count int64  `json:"count"`
	}

	err := r.db.Model(&models.Outfit{}).
		Select("DATE_FORMAT(date, '%Y-%m') as month, COUNT(*) as count").
		Where("user_id = ? AND YEAR(date) = ?", userID, year).
		Group("DATE_FORMAT(date, '%Y-%m')").
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

// SearchOutfits 搜索穿搭记录
func (r *outfitRepository) SearchOutfits(userID uint, query string, limit int) ([]*models.Outfit, error) {
	var outfits []*models.Outfit
	searchTerm := "%" + query + "%"

	dbQuery := r.db.Where("user_id = ?", userID).
		Where("notes LIKE ? OR weather LIKE ? OR top_type LIKE ? OR bottom_type LIKE ? OR shoes_type LIKE ? OR accessories LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm).
		Order("date DESC")

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}

	err := dbQuery.Find(&outfits).Error
	return outfits, err
}

// GetTotalCount 获取用户穿搭总数
func (r *outfitRepository) GetTotalCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Outfit{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}