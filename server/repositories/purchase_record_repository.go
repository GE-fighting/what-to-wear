package repositories

import (
	"time"
	"what-to-wear/server/models"

	"gorm.io/gorm"
)

// purchaseRecordRepository 购买记录仓库实现
type purchaseRecordRepository struct {
	db *gorm.DB
}

// NewPurchaseRecordRepository 创建购买记录仓库实例
func NewPurchaseRecordRepository(db *gorm.DB) PurchaseRecordRepository {
	return &purchaseRecordRepository{db: db}
}

// Create 创建购买记录
func (r *purchaseRecordRepository) Create(record *models.PurchaseRecord) error {
	return r.db.Create(record).Error
}

// GetByID 根据ID获取购买记录
func (r *purchaseRecordRepository) GetByID(id uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetByClothingItemID 根据衣物ID获取购买记录
func (r *purchaseRecordRepository) GetByClothingItemID(clothingItemID uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.Where("clothing_item_id = ?", clothingItemID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 更新购买记录
func (r *purchaseRecordRepository) Update(record *models.PurchaseRecord) error {
	return r.db.Save(record).Error
}

// Delete 删除购买记录
func (r *purchaseRecordRepository) Delete(id uint) error {
	return r.db.Delete(&models.PurchaseRecord{}, id).Error
}

// GetByUserID 根据用户ID获取购买记录
func (r *purchaseRecordRepository) GetByUserID(userID uint, limit int) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	query := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("purchase_records.purchase_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&records).Error
	return records, err
}

// GetByDateRange 根据日期范围获取购买记录
func (r *purchaseRecordRepository) GetByDateRange(userID uint, startDate, endDate string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.purchase_date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetByStore 根据商店获取购买记录
func (r *purchaseRecordRepository) GetByStore(userID uint, storeName string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND (purchase_records.store_name LIKE ? OR purchase_records.online_store LIKE ?)", 
			userID, "%"+storeName+"%", "%"+storeName+"%").
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetTotalSpent 获取用户总消费金额
func (r *purchaseRecordRepository) GetTotalSpent(userID uint) (float64, error) {
	var totalSpent float64
	err := r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(SUM(purchase_records.purchase_price), 0)").
		Scan(&totalSpent).Error
	return totalSpent, err
}

// GetSpentByMonth 获取按月份分组的消费统计
func (r *purchaseRecordRepository) GetSpentByMonth(userID uint, year int) (map[string]float64, error) {
	var results []struct {
		Month string  `json:"month"`
		Total float64 `json:"total"`
	}

	err := r.db.Model(&models.PurchaseRecord{}).
		Select("DATE_FORMAT(purchase_date, '%Y-%m') as month, COALESCE(SUM(purchase_price), 0) as total").
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
func (r *purchaseRecordRepository) GetSpentByCategory(userID uint) (map[string]float64, error) {
	var results []struct {
		CategoryName string  `json:"category_name"`
		Total        float64 `json:"total"`
	}

	err := r.db.Model(&models.PurchaseRecord{}).
		Select("clothing_categories.name as category_name, COALESCE(SUM(purchase_records.purchase_price), 0) as total").
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
func (r *purchaseRecordRepository) GetSpentByStore(userID uint) (map[string]float64, error) {
	var results []struct {
		StoreName string  `json:"store_name"`
		Total     float64 `json:"total"`
	}

	err := r.db.Model(&models.PurchaseRecord{}).
		Select("COALESCE(NULLIF(store_name, ''), online_store) as store_name, COALESCE(SUM(purchase_price), 0) as total").
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND (store_name != '' OR online_store != '')", userID).
		Group("COALESCE(NULLIF(store_name, ''), online_store)").
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
func (r *purchaseRecordRepository) GetAverageItemPrice(userID uint) (float64, error) {
	var avgPrice float64
	err := r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Select("COALESCE(AVG(purchase_records.purchase_price), 0)").
		Scan(&avgPrice).Error
	return avgPrice, err
}

// GetMostExpensiveItem 获取最贵的商品
func (r *purchaseRecordRepository) GetMostExpensiveItem(userID uint) (*models.PurchaseRecord, error) {
	var record models.PurchaseRecord
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ?", userID).
		Order("purchase_records.purchase_price DESC").
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetDiscountStats 获取折扣统计
func (r *purchaseRecordRepository) GetDiscountStats(userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总折扣金额
	var totalDiscount float64
	err := r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.original_price > purchase_records.purchase_price", userID).
		Select("COALESCE(SUM(purchase_records.original_price - purchase_records.purchase_price), 0)").
		Scan(&totalDiscount).Error
	if err != nil {
		return nil, err
	}
	stats["total_discount"] = totalDiscount

	// 平均折扣率
	var avgDiscountRate float64
	err = r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.discount > 0", userID).
		Select("COALESCE(AVG(purchase_records.discount), 0)").
		Scan(&avgDiscountRate).Error
	if err != nil {
		return nil, err
	}
	stats["average_discount_rate"] = avgDiscountRate

	// 有折扣的商品数量
	var discountedCount int64
	err = r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.discount > 0", userID).
		Count(&discountedCount).Error
	if err != nil {
		return nil, err
	}
	stats["discounted_count"] = discountedCount

	return stats, nil
}

// GetWarrantyInfo 获取保修信息
func (r *purchaseRecordRepository) GetWarrantyInfo(userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry IS NOT NULL", userID).
		Order("purchase_records.warranty_expiry ASC").
		Find(&records).Error
	return records, err
}

// GetActiveWarranties 获取有效保修记录
func (r *purchaseRecordRepository) GetActiveWarranties(userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	now := time.Now()
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry > ?", userID, now).
		Order("purchase_records.warranty_expiry ASC").
		Find(&records).Error
	return records, err
}

// GetExpiredWarranties 获取过期保修记录
func (r *purchaseRecordRepository) GetExpiredWarranties(userID uint) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	now := time.Now()
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.warranty_expiry <= ?", userID, now).
		Order("purchase_records.warranty_expiry DESC").
		Find(&records).Error
	return records, err
}

// GetPurchasesByPaymentMethod 根据支付方式获取购买记录
func (r *purchaseRecordRepository) GetPurchasesByPaymentMethod(userID uint, paymentMethod string) ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	err := r.db.Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.payment_method = ?", userID, paymentMethod).
		Order("purchase_records.purchase_date DESC").
		Find(&records).Error
	return records, err
}

// GetPaymentMethodStats 获取支付方式统计
func (r *purchaseRecordRepository) GetPaymentMethodStats(userID uint) (map[string]int64, error) {
	var results []struct {
		PaymentMethod string `json:"payment_method"`
		Count         int64  `json:"count"`
	}

	err := r.db.Model(&models.PurchaseRecord{}).
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
func (r *purchaseRecordRepository) GetPurchaseStats(userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总消费
	totalSpent, err := r.GetTotalSpent(userID)
	if err != nil {
		return nil, err
	}
	stats["total_spent"] = totalSpent

	// 平均价格
	avgPrice, err := r.GetAverageItemPrice(userID)
	if err != nil {
		return nil, err
	}
	stats["average_price"] = avgPrice

	// 购买数量
	var totalCount int64
	err = r.db.Model(&models.PurchaseRecord{}).
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
	err = r.db.Model(&models.PurchaseRecord{}).
		Joins("JOIN clothing_items ON purchase_records.clothing_item_id = clothing_items.id").
		Where("clothing_items.user_id = ? AND purchase_records.purchase_date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
		Select("COALESCE(SUM(purchase_records.purchase_price), 0)").
		Scan(&monthlySpent).Error
	if err != nil {
		return nil, err
	}
	stats["monthly_spent"] = monthlySpent

	return stats, nil
}