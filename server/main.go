package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 引擎
	r := gin.Default()

	// 设置一个 API 路由组
	api := r.Group("/api")
	{
		// 定义一个 GET 请求的接口 /api/ping
		api.GET("/ping", func(c *gin.Context) {
			// 返回一个 JSON 响应
			c.JSON(http.StatusOK, gin.H{
				"message": "pong from go server!",
			})
		})
	}

	// 启动服务，监听在 8080 端口
	r.Run(":8080")
}
