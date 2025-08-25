package container

import (
	"what-to-wear/server/controllers"
	"what-to-wear/server/repositories"
	"what-to-wear/server/services"

	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	// Repositories
	UserRepo             repositories.UserRepository
	OutfitRepo           repositories.OutfitRepository
	OutfitItemRepo       repositories.OutfitItemRepository
	ClothingItemRepo     repositories.ClothingItemRepository
	ClothingCategoryRepo repositories.ClothingCategoryRepository
	AttachmentRepo       repositories.AttachmentRepository
	PurchaseRecordRepo   repositories.PurchaseRecordRepository
	WearRecordRepo       repositories.WearRecordRepository

	// Services
	AuthService           services.AuthService
	UserService           services.UserService
	OutfitService         services.OutfitService
	PurchaseRecordService services.PurchaseRecordService
	WearRecordService     services.WearRecordService
	ClothingItemService   services.ClothingItemService

	// Controllers
	AuthController *controllers.AuthController
	UserController *controllers.UserController
}

// NewContainer 创建容器实例
func NewContainer(db *gorm.DB) *Container {
	// 创建 Repositories
	userRepo := repositories.NewUserRepository(db)
	outfitRepo := repositories.NewOutfitRepository(db)
	outfitItemRepo := repositories.NewOutfitItemRepository(db)
	clothingItemRepo := repositories.NewClothingItemRepository(db)
	clothingCategoryRepo := repositories.NewClothingCategoryRepository(db)
	attachmentRepo := repositories.NewAttachmentRepository(db)
	purchaseRecordRepo := repositories.NewPurchaseRecordRepository(db)
	wearRecordRepo := repositories.NewWearRecordRepository(db)

	// 创建 Services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	outfitService := services.NewOutfitService(
		outfitRepo,
		outfitItemRepo,
		clothingItemRepo,
		clothingCategoryRepo,
		attachmentRepo,
	)
	purchaseRecordService := services.NewPurchaseRecordService(
		purchaseRecordRepo,
		clothingItemRepo,
		clothingCategoryRepo,
	)
	wearRecordService := services.NewWearRecordService(
		wearRecordRepo,
		clothingItemRepo,
	)
	clothingItemService := services.NewClothingItemService(
		clothingItemRepo,
		clothingCategoryRepo,
		attachmentRepo,
		purchaseRecordRepo,
		wearRecordRepo,
	)

	// 创建 Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	return &Container{
		// Repositories
		UserRepo:             userRepo,
		OutfitRepo:           outfitRepo,
		OutfitItemRepo:       outfitItemRepo,
		ClothingItemRepo:     clothingItemRepo,
		ClothingCategoryRepo: clothingCategoryRepo,
		AttachmentRepo:       attachmentRepo,
		PurchaseRecordRepo:   purchaseRecordRepo,
		WearRecordRepo:       wearRecordRepo,

		// Services
		AuthService:           authService,
		UserService:           userService,
		OutfitService:         outfitService,
		PurchaseRecordService: purchaseRecordService,
		WearRecordService:     wearRecordService,
		ClothingItemService:   clothingItemService,

		// Controllers
		AuthController: authController,
		UserController: userController,
	}
}

// GetAuthController 获取认证控制器
func (c *Container) GetAuthController() *controllers.AuthController {
	return c.AuthController
}

// GetUserController 获取用户控制器
func (c *Container) GetUserController() *controllers.UserController {
	return c.UserController
}

// GetOutfitService 获取穿搭服务
func (c *Container) GetOutfitService() services.OutfitService {
	return c.OutfitService
}

// GetWearRecordService 获取穿着记录服务
func (c *Container) GetWearRecordService() services.WearRecordService {
	return c.WearRecordService
}

// GetClothingItemService 获取衣物服务
func (c *Container) GetClothingItemService() services.ClothingItemService {
	return c.ClothingItemService
}
