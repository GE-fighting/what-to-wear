package repositories

import (
	"context"
	"time"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// PurchaseRecordRepository 购买记录仓库接口
type PurchaseRecordRepository interface {
	// 基础CRUD操作
	Create(ctx context.Context, record *models.PurchaseRecord) error
	GetByID(ctx context.Context, id uint) (*models.PurchaseRecord, error)
	GetByClothingItemID(ctx context.Context, clothingItemID uint) (*models.PurchaseRecord, error)
	Update(ctx context.Context, record *models.PurchaseRecord) error
	Delete(ctx context.Context, id uint) error

	// 查询
	GetByUserID(ctx context.Context, userID uint, limit int) ([]models.PurchaseRecord, error)
	GetByDateRange(ctx context.Context, userID uint, startDate, endDate string) ([]models.PurchaseRecord, error)
	GetByStore(ctx context.Context, userID uint, storeName string) ([]models.PurchaseRecord, error)

	// 统计
	GetTotalSpent(ctx context.Context, userID uint) (float64, error)
	GetSpentByMonth(ctx context.Context, userID uint, year int) (map[string]float64, error)
	GetSpentByCategory(ctx context.Context, userID uint) (map[string]float64, error)
	GetAverageItemPrice(ctx context.Context, userID uint) (float64, error)
	GetSpentByStore(ctx context.Context, userID uint) (map[string]float64, error)
}

// purchaseRecordRepository 购买记录仓库实现
type purchaseRecordRepository struct {
	db *gorm.DB
}

// NewPurchaseRecordRepository 创建购买记录仓库实例
func NewPurchaseRecordRepository(db *gorm.DB) PurchaseRecordRepository {
	return &purchaseRecordRepository{db: db}
}

// Create 创建购买记录
func (r *purchaseRecordRepository) Create(ctx context.Context, record *models.PurchaseRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

// GetByID 根据ID获取购买记录
func (r *purchaseRecordRepository) GetByID(ctx context.Context, id uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetByClothingItemID 根据衣物ID获取购买记录
func (r *purchaseRecordRepository) GetByClothingItemID(ctx context.Context, clothingItemID uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.WithContext(ctx).Where("clothing_item_id = ?", clothingItemID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 更新购买记录
func (r *purchaseRecordRepository) Update(ctx context.Context, record *models.PurchaseRecord) error {
	return r.db.WithContext(ctx).Save(record).Error
}

// Delete 删除购买记录
func (r *purchaseRecordRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.PurchaseRecord{}, id).Error
}

// GetByUserID 根据用户ID获取购买记录
func (r *purchaseRecordRepository) GetByUserID(ctx context.Context, userID uint, limit int) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	query := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("purchase_records.purchase_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetByDateRange 根据日期范围获取购买记录
func (r *purchaseRecordRepository) GetByDateRange(ctx context.Context, userID uint, startDate, endDate string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.purchase_date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetByStore 根据商店获取购买记录
func (r *purchaseRecordRepository) GetByStore(ctx context.Context, userID uint, storeName string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.store LIKE ?",
			userID, "%"+storeName+"%").
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetTotalSpent 获取用户总消费金额
func (r *purchaseRecordRepository) GetTotalSpent(ctx context.Context, userID uint) (float64, error) {
	var totalSpent float64
	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(SUM(purchase_records.price), 0)").
		Scan(&totalSpent).Error
	return totalSpent, err
}

// GetSpentByMonth 获取按月份分组的消费统计
func (r *purchaseRecordRepository) GetSpentByMonth(ctx context.Context, userID uint, year int) (map[string]float64, error) {
	var results []struct {
		Month string  `json:"month"`
		Total float64 `json:"total"`
	}

	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Select("DATE_FORMAT(purchase_date, '%Y-%m') as month, COALESCE(SUM(price), 0) as total").
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND YEAR(purchase_records.purchase_date) = ?", userID, year).
		Group("DATE_FORMAT(purchase_date, '%Y-%m')").
		Order("month").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	monthlySpent := make(map[string]float64)
	for _, result := range results {
		monthlySpent[result.Month] = result.Total
	}

	return monthlySpent, nil
}

// GetSpentByCategory 获取按分类分组的消费统计
func (r *purchaseRecordRepository) GetSpentByCategory(ctx context.Context, userID uint) (map[string]float64, error) {
	var results []struct {
		CategoryName string  `json:"category_name"`
		Total        float64 `json:"total"`
	}

	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Select("clothing_categories.name as category_name, COALESCE(SUM(purchase_records.price), 0) as total").
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Joins("JOIN clothing_categories ON clothing_items.category_id = clothing_categories.id").
		Where("clothing_items.user_id = ?", userID).
		Group("clothing_categories.id, clothing_categories.name").
		Order("total DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	categorySpent := make(map[string]float64)
	for _, result := range results {
		categorySpent[result.CategoryName] = result.Total
	}

	return categorySpent, nil
}

// GetSpentByStore 获取按商店分组的消费统计
func (r *purchaseRecordRepository) GetSpentByStore(ctx context.Context, userID uint) (map[string]float64, error) {
	var results []struct {
		StoreName string  `json:"store_name"`
		Total     float64 `json:"total"`
	}

	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Select("purchase_records.store as store_name, COALESCE(SUM(purchase_records.price), 0) as total").
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.store != ''", userID).
		Group("purchase_records.store").
		Order("total DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	storeSpent := make(map[string]float64)
	for _, result := range results {
		if result.StoreName != "" {
			storeSpent[result.StoreName] = result.Total
		}
	}

	return storeSpent, nil
}

// GetAverageItemPrice 获取平均商品价格
func (r *purchaseRecordRepository) GetAverageItemPrice(ctx context.Context, userID uint) (float64, error) {
	var avgPrice float64
	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(AVG(purchase_records.price), 0)").
		Scan(&avgPrice).Error
	return avgPrice, err
}

// GetMostExpensiveItem 获取最贵的商品
func (r *purchaseRecordRepository) GetMostExpensiveItem(ctx context.Context, userID uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("purchase_records.price DESC").
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetDiscountStats 获取折扣统计
func (r *purchaseRecordRepository) GetDiscountStats(ctx context.Context, userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 简化模型中没有折扣字段，返回空统计
	stats["total_discount"] = 0.0
	stats["average_discount_rate"] = 0.0
	stats["discounted_count"] = int64(0)

	return stats, nil
}

// GetWarrantyInfo 获取保修信息
func (r *purchaseRecordRepository) GetWarrantyInfo(ctx context.Context, userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry IS NOT NULL", userID).
		Order("purchase_records.warranty_expiry ASC").
		Find(&records).Error
	return records, err
}

// GetActiveWarranties 获取有效保修记录
func (r *purchaseRecordRepository) GetActiveWarranties(ctx context.Context, userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	now := time.Now()
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry > ?", userID, now).
		Order("purchase_records.warranty_expiry ASC").
		Find(&records).Error
	return records, err
}

// GetExpiredWarranties 获取过期保修记录
func (r *purchaseRecordRepository) GetExpiredWarranties(ctx context.Context, userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	now := time.Now()
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry <= ?", userID, now).
		Order("purchase_records.warranty_expiry DESC").
		Find(&records).Error
	return records, err
}

// GetPurchasesByPaymentMethod 根据支付方式获取购买记录
func (r *purchaseRecordRepository) GetPurchasesByPaymentMethod(ctx context.Context, userID uint, paymentMethod string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.WithContext(ctx).Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.payment_method = ?", userID, paymentMethod).
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetPaymentMethodStats 获取支付方式统计
func (r *purchaseRecordRepository) GetPaymentMethodStats(ctx context.Context, userID uint) (map[string]int64, error) {
	var results []struct {
		PaymentMethod string `json:"payment_method"`
		Count         int64  `json:"count"`
	}

	err := r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Select("payment_method, COUNT(*) as count").
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND payment_method != ''", userID).
		Group("payment_method").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	paymentStats := make(map[string]int64)
	for _, result := range results {
		paymentStats[result.PaymentMethod] = result.Count
	}

	return paymentStats, nil
}

// GetPurchaseStats 获取购买统计信息
func (r *purchaseRecordRepository) GetPurchaseStats(ctx context.Context, userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总消费
	totalSpent, err := r.GetTotalSpent(ctx, userID)
	if err != nil {
		return nil, err
	}
	stats["total_spent"] = totalSpent

	// 平均价格
	avgPrice, err := r.GetAverageItemPrice(ctx, userID)
	if err != nil {
		return nil, err
	}
	stats["average_price"] = avgPrice

	// 购买数量
	var totalCount int64
	err = r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	stats["total_count"] = totalCount

	// 本月消费
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	var monthlySpent float64
	err = r.db.WithContext(ctx).Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.purchase_date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
		Select("COALESCE(SUM(purchase_records.price), 0)").
		Scan(&monthlySpent).Error
	if err != nil {
		return nil, err
	}
	stats["monthly_spent"] = monthlySpent

	return stats, nil
}
