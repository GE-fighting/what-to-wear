package repositories

import (
	"what-to-wear/server/models"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	// 创建用户
	Create(user *models.User) error

	// 根据用户名查找用户
	GetByUsername(username string) (*models.User, error)

	// 根据邮箱查找用户
	GetByEmail(email string) (*models.User, error)

	// 根据ID查找用户
	GetByID(id uint) (*models.User, error)

	// 更新用户信息
	Update(user *models.User) error

	// 删除用户
	Delete(id uint) error

	// 检查用户名是否存在
	ExistsByUsername(username string) (bool, error)

	// 检查邮箱是否存在
	ExistsByEmail(email string) (bool, error)
}

// AttachmentRepository 附件仓库接口
type AttachmentRepository interface {
	// 基础CRUD操作
	Create(attachment *models.Attachment) error
	GetByID(id uint) (*models.Attachment, error)
	Update(attachment *models.Attachment) error
	Delete(id uint) error

	// 查询操作
	GetByEntityID(entityType models.EntityType, entityID uint) ([]models.Attachment, error)
	GetByUserID(userID uint, limit int) ([]models.Attachment, error)
	GetByType(attachmentType models.AttachmentType, limit int) ([]models.Attachment, error)

	// 统计操作
	GetTotalSize(userID uint) (int64, error)
	GetCountByType(userID uint) (map[models.AttachmentType]int64, error)
}

// OutfitRepository 穿搭数据访问接口
type OutfitRepository interface {
	// 创建穿搭记录
	Create(outfit *models.Outfit) error

	// 根据用户ID获取穿搭历史
	GetByUserID(userID uint, limit, offset int) ([]*models.Outfit, error)

	// 根据ID获取穿搭记录
	GetByID(id uint) (*models.Outfit, error)

	// 更新穿搭记录
	Update(outfit *models.Outfit) error

	// 删除穿搭记录
	Delete(id uint) error
}

// OutfitItemRepository 穿搭单品仓库接口
type OutfitItemRepository interface {
	// 基础CRUD操作
	Create(item *models.OutfitItem) error
	GetByID(id uint) (*models.OutfitItem, error)
	Update(item *models.OutfitItem) error
	Delete(id uint) error

	// 查询操作
	GetByOutfitID(outfitID uint) ([]models.OutfitItem, error)
	GetByClothingItemID(clothingItemID uint) ([]models.OutfitItem, error)
	GetByRole(outfitID uint, role string) ([]models.OutfitItem, error)

	// 批量操作
	CreateBatch(items []models.OutfitItem) error
	DeleteByOutfitID(outfitID uint) error
	UpdateLayerOrders(outfitID uint, items []models.OutfitItem) error

	// 统计操作
	GetItemUsageCount(clothingItemID uint) (int64, error)
	GetPopularItems(userID uint, limit int) ([]models.ClothingItem, error)
}
