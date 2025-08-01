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
	UserRepo repositories.UserRepository

	// Services
	AuthService services.AuthService
	UserService services.UserService

	// Controllers
	AuthController *controllers.AuthController
	UserController *controllers.UserController
}

// NewContainer 创建容器实例
func NewContainer(db *gorm.DB) *Container {
	// 创建 Repositories
	userRepo := repositories.NewUserRepository(db)

	// 创建 Services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	// 创建 Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	return &Container{
		// Repositories
		UserRepo: userRepo,

		// Services
		AuthService: authService,
		UserService: userService,

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
