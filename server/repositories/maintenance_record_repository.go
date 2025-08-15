package repositories

import (
	"time"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// maintenanceRecordRepository 保养记录仓库实现
type maintenanceRecordRepository struct {
	db *gorm.DB
}

// NewMaintenanceRecordRepository 创建保养记录仓库实例
func NewMaintenanceRecordRepository(db *gorm.DB) MaintenanceRecordRepository {
	return &maintenanceRecordRepository{db: db}
}

// Create 创建保养记录
func (r *maintenanceRecordRepository) Create(record *models.MaintenanceRecord) error {
	return r.db.Create(record).Error
}

// GetByID 根据ID获取保养记录
func (r *maintenanceRecordRepository) GetByID(id uint) (*models.MaintenanceRecord, error) {
	var record models.MaintenanceRecord
	err := r.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetByClothingItemID 根据衣物ID获取保养记录
func (r *maintenanceRecordRepository) GetByClothingItemID(clothingItemID uint) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	err := r.db.Where("clothing_item_id = ?", clothingItemID).
		Order("maintenance_date DESC").
		Find(&records).Error
	return records, err
}

// Update 更新保养记录
func (r *maintenanceRecordRepository) Update(record *models.MaintenanceRecord) error {
	return r.db.Save(record).Error
}

// Delete 删除保养记录
func (r *maintenanceRecordRepository) Delete(id uint) error {
	return r.db.Delete(&models.MaintenanceRecord{}, id).Error
}

// GetByUserID 根据用户ID获取保养记录
func (r *maintenanceRecordRepository) GetByUserID(userID uint, limit int) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	query := r.db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("maintenance_records.maintenance_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetByType 根据保养类型获取记录
func (r *maintenanceRecordRepository) GetByType(userID uint, maintenanceType models.MaintenanceType) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	err := r.db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.maintenance_type = ?", userID, maintenanceType).
		Order("maintenance_records.maintenance_date DESC").
		Find(&records).Error
	return records, err
}

// GetUpcoming 获取即将到期的保养提醒
func (r *maintenanceRecordRepository) GetUpcoming(userID uint, days int) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	futureDate := time.Now().AddDate(0, 0, days)

	err := r.db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.next_maintenance_date <= ? AND maintenance_records.reminder_sent = ?",
			userID, futureDate, false).
		Order("maintenance_records.next_maintenance_date ASC").
		Find(&records).Error

	return records, err
}

// GetOverdue 获取过期的保养记录
func (r *maintenanceRecordRepository) GetOverdue(userID uint) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	now := time.Now()

	err := r.db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.next_maintenance_date < ?", userID, now).
		Order("maintenance_records.next_maintenance_date ASC").
		Find(&records).Error

	return records, err
}

// GetMaintenanceCost 获取用户保养总费用
func (r *maintenanceRecordRepository) GetMaintenanceCost(userID uint) (float64, error) {
	var totalCost float64
	err := r.db.Model(&models.MaintenanceRecord{}).
		Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(SUM(maintenance_records.cost), 0)").
		Scan(&totalCost).Error
	return totalCost, err
}

// GetMaintenanceFrequency 获取保养频率统计
func (r *maintenanceRecordRepository) GetMaintenanceFrequency(userID uint) (map[string]int64, error) {
	var results []struct {
		MaintenanceType string `json:"maintenance_type"`
		Count           int64  `json:"count"`
	}

	err := r.db.Model(&models.MaintenanceRecord{}).
		Select("maintenance_type, COUNT(*) as count").
		Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Group("maintenance_type").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	frequencyMap := make(map[string]int64)
	for _, result := range results {
		frequencyMap[result.MaintenanceType] = result.Count
	}

	return frequencyMap, nil
}

// GetMaintenanceCostByType 获取按类型分组的保养费用
func (r *maintenanceRecordRepository) GetMaintenanceCostByType(userID uint) (map[string]float64, error) {
	var results []struct {
		MaintenanceType string  `json:"maintenance_type"`
		TotalCost       float64 `json:"total_cost"`
	}

	err := r.db.Model(&models.MaintenanceRecord{}).
		Select("maintenance_type, COALESCE(SUM(cost), 0) as total_cost").
		Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Group("maintenance_type").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	costMap := make(map[string]float64)
	for _, result := range results {
		costMap[result.MaintenanceType] = result.TotalCost
	}

	return costMap, nil
}

// GetAverageEffectiveness 获取平均保养效果评分
func (r *maintenanceRecordRepository) GetAverageEffectiveness(userID uint) (float64, error) {
	var avgEffectiveness float64
	err := r.db.Model(&models.MaintenanceRecord{}).
		Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.effectiveness_score > 0", userID).
		Select("COALESCE(AVG(maintenance_records.effectiveness_score), 0)").
		Scan(&avgEffectiveness).Error
	return avgEffectiveness, err
}

// GetMaintenanceByDateRange 根据日期范围获取保养记录
func (r *maintenanceRecordRepository) GetMaintenanceByDateRange(userID uint, startDate, endDate time.Time) ([]models.MaintenanceRecord, error) {
	var records []models.MaintenanceRecord
	err := r.db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.maintenance_date BETWEEN ? AND ?",
			userID, startDate, endDate).
		Order("maintenance_records.maintenance_date DESC").
		Find(&records).Error
	return records, err
}

// MarkReminderSent 标记提醒已发送
func (r *maintenanceRecordRepository) MarkReminderSent(recordID uint) error {
	return r.db.Model(&models.MaintenanceRecord{}).
		Where("id = ?", recordID).
		Update("reminder_sent", true).Error
}

// GetMaintenanceStats 获取保养统计信息
func (r *maintenanceRecordRepository) GetMaintenanceStats(userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总费用
	totalCost, err := r.GetMaintenanceCost(userID)
	if err != nil {
		return nil, err
	}
	stats["total_cost"] = totalCost

	// 总次数
	var totalCount int64
	err = r.db.Model(&models.MaintenanceRecord{}).
		Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	stats["total_count"] = totalCount

	// 平均效果
	avgEffectiveness, err := r.GetAverageEffectiveness(userID)
	if err != nil {
		return nil, err
	}
	stats["average_effectiveness"] = avgEffectiveness

	// 即将到期数量
	upcomingRecords, err := r.GetUpcoming(userID, 30)
	if err != nil {
		return nil, err
	}
	stats["upcoming_count"] = len(upcomingRecords)

	// 过期数量
	overdueRecords, err := r.GetOverdue(userID)
	if err != nil {
		return nil, err
	}
	stats["overdue_count"] = len(overdueRecords)

	return stats, nil
}