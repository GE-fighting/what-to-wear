package routes

import (
	"what-to-wear/server/controllers"
	"what-to-wear/server/middleware"

	"github.com/gin-gonic/gin"
)

// SetupClothingRoutes 设置衣物管理相关路由
func SetupClothingRoutes(router *gin.RouterGroup, clothingController *controllers.ClothingController) {
	// 衣物管理API组，需要认证
	clothingAPI := router.Group("/clothing")
	clothingAPI.Use(middleware.AuthMiddleware())
	{
		// 衣物CRUD操作
		clothingAPI.POST("/item", clothingController.CreateClothingItem)
		clothingAPI.GET("/items", clothingController.GetClothingItems)
		clothingAPI.GET("/items/:id", clothingController.GetClothingItem)
		clothingAPI.PUT("/items/:id", clothingController.UpdateClothingItem)
		clothingAPI.DELETE("/items/:id", clothingController.DeleteClothingItem)

		// 衣物统计
		clothingAPI.GET("/stats", clothingController.GetClothingStats)

		// 穿着记录
		clothingAPI.POST("/items/:id/wear", clothingController.RecordWear)

		// 分类管理
		clothingAPI.GET("/categories", clothingController.GetCategories)
		clothingAPI.GET("/categories/tree", clothingController.GetCategoryTree)

		// 标签管理
		clothingAPI.GET("/tags", clothingController.GetTags)
		clothingAPI.GET("/tags/:type", clothingController.GetTagsByType)
	}

	// 公开API组，不需要认证（用于获取系统预设数据）
	publicAPI := router.Group("/public/clothing")
	{
		publicAPI.GET("/categories", clothingController.GetCategories)
		publicAPI.GET("/categories/tree", clothingController.GetCategoryTree)
		publicAPI.GET("/tags/system", clothingController.GetTags) // 只返回系统标签

		// 系统标签枚举API（从内存获取，无需数据库查询）
		publicAPI.GET("/tags/enums/all", clothingController.GetAllSystemTagEnums)      // 获取所有系统标签枚举
		publicAPI.GET("/tags/enums/:type", clothingController.GetSystemTagEnumsByType) // 根据类型获取系统标签枚举
	}
}

// 扩展路由配置，包含更多功能
func SetupExtendedClothingRoutes(router *gin.Engine, clothingController *controllers.ClothingController) {
	clothingAPI := router.Group("/api/clothing")
	clothingAPI.Use(middleware.AuthMiddleware())
	{
		// 高级搜索和筛选
		clothingAPI.GET("/search", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "搜索功能待实现"})
		})

		// 推荐系统
		clothingAPI.GET("/recommendations", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "推荐功能待实现"})
		})

		// 衣物保养
		maintenanceGroup := clothingAPI.Group("/maintenance")
		{
			maintenanceGroup.POST("/items/:id/records", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "保养记录功能待实现"})
			})
			maintenanceGroup.GET("/items/:id/records", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "保养记录功能待实现"})
			})
			maintenanceGroup.GET("/reminders", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "保养提醒功能待实现"})
			})
		}

		// 穿着分析
		analyticsGroup := clothingAPI.Group("/analytics")
		{
			analyticsGroup.GET("/wear-frequency", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "穿着频率分析待实现"})
			})
			analyticsGroup.GET("/comfort-analysis", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "舒适度分析待实现"})
			})
			analyticsGroup.GET("/cost-per-wear", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "成本分析待实现"})
			})
		}

		// 购买记录
		purchaseGroup := clothingAPI.Group("/purchases")
		{
			purchaseGroup.POST("/items/:id/records", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "购买记录功能待实现"})
			})
			purchaseGroup.GET("/records", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "购买记录功能待实现"})
			})
			purchaseGroup.GET("/stats", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "购买统计功能待实现"})
			})
		}

		// 标签管理（用户自定义标签）
		tagGroup := clothingAPI.Group("/tags")
		{
			tagGroup.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "创建标签功能待实现"})
			})
			tagGroup.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新标签功能待实现"})
			})
			tagGroup.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除标签功能待实现"})
			})
		}

		// 分类管理（管理员功能）
		categoryGroup := clothingAPI.Group("/categories")
		{
			categoryGroup.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "创建分类功能待实现"})
			})
			categoryGroup.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新分类功能待实现"})
			})
			categoryGroup.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除分类功能待实现"})
			})
		}

		// 批量操作
		batchGroup := clothingAPI.Group("/batch")
		{
			batchGroup.POST("/items/tags", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "批量标签功能待实现"})
			})
			batchGroup.DELETE("/items/tags", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "批量标签功能待实现"})
			})
			batchGroup.PUT("/items/category", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "批量分类功能待实现"})
			})
			batchGroup.DELETE("/items", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "批量删除功能待实现"})
			})
		}

		// 导入导出
		importExportGroup := clothingAPI.Group("/import-export")
		{
			importExportGroup.POST("/import", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "导入功能待实现"})
			})
			importExportGroup.GET("/export", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "导出功能待实现"})
			})
		}
	}
}
