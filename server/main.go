package main

import (
	"what-to-wear/server/config"
	"what-to-wear/server/container"
	"what-to-wear/server/logger"
	"what-to-wear/server/middleware"
	"what-to-wear/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志系统
	config.InitLogger()
	logger.Info("Starting What-to-Wear server")

	// 连接数据库
	config.ConnectDatabase()
	logger.Info("Database connected successfully")

	// 创建依赖注入容器
	appContainer := container.NewContainer(config.DB)
	logger.Info("Dependency injection container initialized")

	// 创建Gin引擎
	r := gin.New() // 使用gin.New()而不是gin.Default()来避免默认日志

	// 添加中间件
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.ErrorLoggingMiddleware())
	r.Use(gin.Recovery())

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1420", "https://tauri.localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
		AllowCredentials: true,
	}))

	// 配置路由
	routes.SetupRoutes(r, appContainer)
	logger.Info("Routes configured successfully")

	// 启动服务
	logger.Info("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", logger.Fields{
			"error": err.Error(),
		})
	}
}
