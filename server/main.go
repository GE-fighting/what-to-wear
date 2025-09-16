package main

import (
	"what-to-wear/server/config"
	"what-to-wear/server/container"
	"what-to-wear/server/database"
	"what-to-wear/server/logger"
	"what-to-wear/server/middleware"
	"what-to-wear/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志系统
	config.InitLogger()
	log := logger.GetLogger()
	log.Info("Starting What-to-Wear server")

	// 初始化数据库（连接 + 迁移 + 种子数据）
	if err := database.Initialize(); err != nil {
		log.Fatal("Database initialization failed", logger.Fields{
			"error": err.Error(),
		})
	}
	log.Info("Database initialized successfully")

	// 创建依赖注入容器
	appContainer := container.NewContainer(database.GetDB())
	log.Info("Dependency injection container initialized")

	// 创建Gin引擎
	r := gin.New() // 使用gin.New()而不是gin.Default()来避免默认日志

	// 添加中间件
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.ErrorLoggingMiddleware())
	r.Use(gin.Recovery())

	// 配置CORS中间件
	r.Use(config.GetCORSMiddleware())

	// 配置路由
	routes.SetupRoutes(r, appContainer)
	log.Info("Routes configured successfully")

	// 启动服务
	log.Info("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", logger.Fields{
			"error": err.Error(),
		})
	}
}
