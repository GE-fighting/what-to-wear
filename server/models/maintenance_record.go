package models

import (
	"time"

	"gorm.io/gorm"
	"what-to-wear/server/api"
)

// MaintenanceRecord 保养记录模型
type MaintenanceRecord struct {
	gorm.Model
	ClothingItemID      uint                `json:"clothing_item_id" gorm:"not null;index"`
	MaintenanceType     api.MaintenanceType `json:"maintenance_type" gorm:"not null"`
	Cost                float64             `json:"cost" gorm:"type:decimal(10,2);default:0"`
	MaintenanceDate     time.Time           `json:"maintenance_date" gorm:"not null"`
	ServiceProvider     string              `json:"service_provider"`                     // 服务提供商（如干洗店名称）
	ServiceLocation     string              `json:"service_location"`                     // 服务地点
	BeforeCondition     *api.ClothingStatus `json:"before_condition"`                     // 保养前状态
	AfterCondition      *api.ClothingStatus `json:"after_condition"`                      // 保养后状态
	EffectivenessScore  int                 `json:"effectiveness_score" gorm:"default:0"` // 保养效果评分 (1-10)
	Notes               string              `json:"notes"`
	Images              []string            `json:"images" gorm:"type:json"`            // 保养前后对比图
	NextMaintenanceDate *time.Time          `json:"next_maintenance_date"`              // 下次保养建议时间
	IsScheduled         bool                `json:"is_scheduled" gorm:"default:false"`  // 是否为计划保养
	ReminderSent        bool                `json:"reminder_sent" gorm:"default:false"` // 是否已发送提醒
}

// TableName 指定表名
func (MaintenanceRecord) TableName() string {
	return "maintenance_records"
}

// CalculateNextMaintenanceDate 计算下次保养建议时间
func (m *MaintenanceRecord) CalculateNextMaintenanceDate() {
	var interval time.Duration

	switch m.MaintenanceType {
	case api.MaintenanceWashing:
		interval = 30 * 24 * time.Hour // 30天后
	case api.MaintenanceDryCleaning:
		interval = 90 * 24 * time.Hour // 90天后
	case api.MaintenanceRepair:
		interval = 180 * 24 * time.Hour // 180天后
	case api.MaintenancePolishing:
		interval = 60 * 24 * time.Hour // 60天后
	case api.MaintenanceWaterproof:
		interval = 365 * 24 * time.Hour // 1年后
	case api.MaintenanceStorage:
		interval = 180 * 24 * time.Hour // 180天后
	default:
		interval = 90 * 24 * time.Hour // 默认90天后
	}

	nextDate := m.MaintenanceDate.Add(interval)
	m.NextMaintenanceDate = &nextDate
}

// GetMaintenanceEffect 获取保养效果对耐久度的影响
func (m *MaintenanceRecord) GetMaintenanceEffect() float64 {
	baseEffect := map[api.MaintenanceType]float64{
		api.MaintenanceWashing:     2.0,
		api.MaintenanceDryCleaning: 3.0,
		api.MaintenanceRepair:      5.0,
		api.MaintenancePolishing:   2.5,
		api.MaintenanceWaterproof:  1.5,
		api.MaintenanceStorage:     1.0,
		api.MaintenanceOther:       1.5,
	}

	effect := baseEffect[m.MaintenanceType]

	// 根据效果评分调整
	if m.EffectivenessScore > 0 {
		effect = effect * (float64(m.EffectivenessScore) / 10.0)
	}

	return effect
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
