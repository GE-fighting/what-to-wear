package models

import (
	"time"

	"gorm.io/gorm"
)

// MaintenanceType 保养类型枚举
type MaintenanceType string

const (
	MaintenanceWashing     MaintenanceType = "washing"      // 清洗
	MaintenanceDryCleaning MaintenanceType = "dry_cleaning" // 干洗
	MaintenanceRepair      MaintenanceType = "repair"       // 修补
	MaintenancePolishing   MaintenanceType = "polishing"    // 抛光
	MaintenanceWaterproof  MaintenanceType = "waterproof"   // 防水处理
	MaintenanceStorage     MaintenanceType = "storage"      // 收纳保存
	MaintenanceOther       MaintenanceType = "other"        // 其他
)

// MaintenanceRecord 保养记录模型
type MaintenanceRecord struct {
	gorm.Model
	ClothingItemID      uint              `json:"clothing_item_id" gorm:"not null;index"`
	MaintenanceType     MaintenanceType   `json:"maintenance_type" gorm:"not null"`
	Cost                float64           `json:"cost" gorm:"type:decimal(10,2);default:0"`
	MaintenanceDate     time.Time         `json:"maintenance_date" gorm:"not null"`
	ServiceProvider     string            `json:"service_provider"`                     // 服务提供商（如干洗店名称）
	ServiceLocation     string            `json:"service_location"`                     // 服务地点
	BeforeCondition     ClothingCondition `json:"before_condition"`                     // 保养前状态
	AfterCondition      ClothingCondition `json:"after_condition"`                      // 保养后状态
	EffectivenessScore  int               `json:"effectiveness_score" gorm:"default:0"` // 保养效果评分 (1-10)
	Notes               string            `json:"notes"`
	Images              []string          `json:"images" gorm:"type:json"`            // 保养前后对比图
	NextMaintenanceDate *time.Time        `json:"next_maintenance_date"`              // 下次保养建议时间
	IsScheduled         bool              `json:"is_scheduled" gorm:"default:false"`  // 是否为计划保养
	ReminderSent        bool              `json:"reminder_sent" gorm:"default:false"` // 是否已发送提醒
}

// TableName 指定表名
func (MaintenanceRecord) TableName() string {
	return "maintenance_records"
}

// CalculateNextMaintenanceDate 计算下次保养建议时间
func (m *MaintenanceRecord) CalculateNextMaintenanceDate() {
	var interval time.Duration

	switch m.MaintenanceType {
	case MaintenanceWashing:
		interval = 30 * 24 * time.Hour // 30天后
	case MaintenanceDryCleaning:
		interval = 90 * 24 * time.Hour // 90天后
	case MaintenanceRepair:
		interval = 180 * 24 * time.Hour // 180天后
	case MaintenancePolishing:
		interval = 60 * 24 * time.Hour // 60天后
	case MaintenanceWaterproof:
		interval = 365 * 24 * time.Hour // 1年后
	case MaintenanceStorage:
		interval = 180 * 24 * time.Hour // 180天后
	default:
		interval = 90 * 24 * time.Hour // 默认90天后
	}

	nextDate := m.MaintenanceDate.Add(interval)
	m.NextMaintenanceDate = &nextDate
}

// GetMaintenanceEffect 获取保养效果对耐久度的影响
func (m *MaintenanceRecord) GetMaintenanceEffect() float64 {
	baseEffect := map[MaintenanceType]float64{
		MaintenanceWashing:     2.0,
		MaintenanceDryCleaning: 3.0,
		MaintenanceRepair:      5.0,
		MaintenancePolishing:   2.5,
		MaintenanceWaterproof:  1.5,
		MaintenanceStorage:     1.0,
		MaintenanceOther:       1.5,
	}

	effect := baseEffect[m.MaintenanceType]

	// 根据效果评分调整
	if m.EffectivenessScore > 0 {
		effect = effect * (float64(m.EffectivenessScore) / 10.0)
	}

	return effect
}

// IsValidMaintenanceType 检查保养类型是否有效
func IsValidMaintenanceType(maintenanceType string) bool {
	validTypes := []MaintenanceType{
		MaintenanceWashing, MaintenanceDryCleaning, MaintenanceRepair,
		MaintenancePolishing, MaintenanceWaterproof, MaintenanceStorage, MaintenanceOther,
	}

	for _, validType := range validTypes {
		if MaintenanceType(maintenanceType) == validType {
			return true
		}
	}
	return false
}

// BeforeCreate GORM钩子：创建前自动计算下次保养时间
func (m *MaintenanceRecord) BeforeCreate(tx *gorm.DB) error {
	m.CalculateNextMaintenanceDate()
	return nil
}

// BeforeUpdate GORM钩子：更新前重新计算下次保养时间
func (m *MaintenanceRecord) BeforeUpdate(tx *gorm.DB) error {
	m.CalculateNextMaintenanceDate()
	return nil
}

// GetMaintenanceRecordsByItem 获取指定衣物的保养记录
func GetMaintenanceRecordsByItem(db *gorm.DB, clothingItemID uint) ([]MaintenanceRecord, error) {
	var records []MaintenanceRecord
	err := db.Where("clothing_item_id = ?", clothingItemID).
		Order("maintenance_date DESC").
		Find(&records).Error
	return records, err
}

// GetUpcomingMaintenance 获取即将到期的保养提醒
func GetUpcomingMaintenance(db *gorm.DB, userID uint, days int) ([]MaintenanceRecord, error) {
	var records []MaintenanceRecord
	futureDate := time.Now().AddDate(0, 0, days)

	err := db.Joins("JOIN clothing_items ON maintenance_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND maintenance_records.next_maintenance_date <= ? AND maintenance_records.reminder_sent = ?",
			userID, futureDate, false).
		Find(&records).Error

	return records, err
}
