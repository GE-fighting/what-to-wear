package main

import (
	"what-to-wear/server/config"
	"what-to-wear/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 创建Gin引擎
	r := gin.Default()

	// 配置路由
	routes.SetupRoutes(r)

	// 启动服务
	r.Run(":8080")
}
