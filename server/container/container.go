package container

import (
	"log"

	"what-to-wear/server/config"
	"what-to-wear/server/controllers"
	"what-to-wear/server/repositories"
	"what-to-wear/server/services"

	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	Config                *config.Config
	DB                    *gorm.DB
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
	OSSService            services.OSSService

	// Controllers
	AuthController     *controllers.AuthController
	UserController     *controllers.UserController
	ClothingController *controllers.ClothingController
	OSSController      *controllers.OSSController
}

// NewContainer 创建容器实例
func NewContainer(cfg *config.Config, db *gorm.DB) *Container {
	// 创建 Repositories
	userRepo := repositories.NewUserRepository(db)
	outfitRepo := repositories.NewOutfitRepository(db)
	outfitItemRepo := repositories.NewOutfitItemRepository(db)
	clothingItemRepo := repositories.NewClothingItemRepository(db)
	clothingCategoryRepo := repositories.NewClothingCategoryRepository(db)
	clothingTagRepository := repositories.NewClothingTagRepository(db)
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
	clothingCategoryService := services.NewCategoryService(clothingCategoryRepo)
	clothingTagService := services.NewClothingTagService(clothingTagRepository)

	// 创建 OSS Service（传入 config）
	ossService, err := services.NewOSSService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize OSS service: %v", err)
	}

	// 创建 Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	clothingController := controllers.NewClothingController(
		clothingItemService,
		clothingCategoryService,
		clothingTagService,
		wearRecordService,
	)
	ossController := controllers.NewOSSController(ossService)

	return &Container{
		Config:              cfg,
		DB:                  db,
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
		OSSService:            ossService,

		// Controllers
		AuthController:     authController,
		UserController:     userController,
		ClothingController: clothingController,
		OSSController:      ossController,
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

// GetClothingController 获取衣物控制器
func (c *Container) GetClothingController() *controllers.ClothingController {
	return c.ClothingController
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

// GetOSSController 获取OSS控制器
func (c *Container) GetOSSController() *controllers.OSSController {
	return c.OSSController
}
