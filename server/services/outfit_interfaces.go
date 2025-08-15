package services

import (
	"what-to-wear/server/dto"
	"what-to-wear/server/models"
)

// OutfitService 穿搭服务接口
type OutfitService interface {
	// 创建穿搭记录
	CreateOutfit(userID uint, req *dto.CreateOutfitRequest) (*models.Outfit, error)

	// 获取用户穿搭历史
	GetUserOutfits(userID uint, page, pageSize int) ([]*models.Outfit, int64, error)

	// 获取穿搭推荐
	GetOutfitRecommendation(userID uint, weather string) (*models.OutfitRecommendation, error)

	// 评价穿搭
	RateOutfit(userID, outfitID uint, rating int, notes string) error
}
