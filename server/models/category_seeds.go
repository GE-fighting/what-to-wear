package models

import (
	"gorm.io/gorm"
)

// SystemCategory 系统预设分类结构
type SystemCategory struct {
	Name        string
	Description string
	ParentName  string // 父分类名称，用于建立层级关系
	Icon        string
	SortOrder   int
}

// GetSystemCategories 获取系统预设分类
func GetSystemCategories() []SystemCategory {
	return []SystemCategory{
		// 一级分类
		{Name: "上衣", Description: "各类上身衣物", Icon: "👕", SortOrder: 1},
		{Name: "下装", Description: "各类下身衣物", Icon: "👖", SortOrder: 2},
		{Name: "鞋子", Description: "各类鞋履", Icon: "👟", SortOrder: 3},
		{Name: "配饰", Description: "各类配饰用品", Icon: "👜", SortOrder: 4},
		{Name: "内衣", Description: "贴身衣物", Icon: "🩲", SortOrder: 5},
		{Name: "外套", Description: "各类外套", Icon: "🧥", SortOrder: 6},

		// 上衣二级分类
		{Name: "T恤", Description: "短袖T恤衫", ParentName: "上衣", Icon: "👕", SortOrder: 1},
		{Name: "衬衫", Description: "各类衬衫", ParentName: "上衣", Icon: "👔", SortOrder: 2},
		{Name: "毛衣", Description: "针织毛衣", ParentName: "上衣", Icon: "🧶", SortOrder: 3},
		{Name: "背心", Description: "无袖背心", ParentName: "上衣", Icon: "🎽", SortOrder: 4},
		{Name: "吊带", Description: "吊带上衣", ParentName: "上衣", Icon: "👙", SortOrder: 5},
		{Name: "卫衣", Description: "休闲卫衣", ParentName: "上衣", Icon: "👘", SortOrder: 6},

		// 下装二级分类
		{Name: "牛仔裤", Description: "各类牛仔裤", ParentName: "下装", Icon: "👖", SortOrder: 1},
		{Name: "休闲裤", Description: "休闲长裤", ParentName: "下装", Icon: "👖", SortOrder: 2},
		{Name: "西裤", Description: "正装西裤", ParentName: "下装", Icon: "👔", SortOrder: 3},
		{Name: "短裤", Description: "各类短裤", ParentName: "下装", Icon: "🩳", SortOrder: 4},
		{Name: "裙子", Description: "各类裙装", ParentName: "下装", Icon: "👗", SortOrder: 5},
		{Name: "运动裤", Description: "运动长裤", ParentName: "下装", Icon: "🏃", SortOrder: 6},

		// 鞋子二级分类
		{Name: "运动鞋", Description: "各类运动鞋", ParentName: "鞋子", Icon: "👟", SortOrder: 1},
		{Name: "皮鞋", Description: "正装皮鞋", ParentName: "鞋子", Icon: "👞", SortOrder: 2},
		{Name: "靴子", Description: "各类靴子", ParentName: "鞋子", Icon: "👢", SortOrder: 3},
		{Name: "凉鞋", Description: "夏季凉鞋", ParentName: "鞋子", Icon: "👡", SortOrder: 4},
		{Name: "拖鞋", Description: "居家拖鞋", ParentName: "鞋子", Icon: "🩴", SortOrder: 5},
		{Name: "高跟鞋", Description: "女式高跟鞋", ParentName: "鞋子", Icon: "👠", SortOrder: 6},

		// 配饰二级分类
		{Name: "包包", Description: "各类包袋", ParentName: "配饰", Icon: "👜", SortOrder: 1},
		{Name: "帽子", Description: "各类帽子", ParentName: "配饰", Icon: "👒", SortOrder: 2},
		{Name: "围巾", Description: "围巾丝巾", ParentName: "配饰", Icon: "🧣", SortOrder: 3},
		{Name: "手表", Description: "各类手表", ParentName: "配饰", Icon: "⌚", SortOrder: 4},
		{Name: "首饰", Description: "项链耳环等", ParentName: "配饰", Icon: "💍", SortOrder: 5},
		{Name: "眼镜", Description: "眼镜墨镜", ParentName: "配饰", Icon: "👓", SortOrder: 6},
		{Name: "腰带", Description: "各类腰带", ParentName: "配饰", Icon: "👔", SortOrder: 7},

		// 内衣二级分类
		{Name: "内裤", Description: "各类内裤", ParentName: "内衣", Icon: "🩲", SortOrder: 1},
		{Name: "文胸", Description: "女式文胸", ParentName: "内衣", Icon: "👙", SortOrder: 2},
		{Name: "保暖内衣", Description: "保暖内衣", ParentName: "内衣", Icon: "🧥", SortOrder: 3},
		{Name: "袜子", Description: "各类袜子", ParentName: "内衣", Icon: "🧦", SortOrder: 4},

		// 外套二级分类
		{Name: "夹克", Description: "各类夹克", ParentName: "外套", Icon: "🧥", SortOrder: 1},
		{Name: "大衣", Description: "长款大衣", ParentName: "外套", Icon: "🧥", SortOrder: 2},
		{Name: "羽绒服", Description: "保暖羽绒服", ParentName: "外套", Icon: "🧥", SortOrder: 3},
		{Name: "西装", Description: "正装西装", ParentName: "外套", Icon: "🤵", SortOrder: 4},
		{Name: "风衣", Description: "时尚风衣", ParentName: "外套", Icon: "🧥", SortOrder: 5},
		{Name: "开衫", Description: "针织开衫", ParentName: "外套", Icon: "🧶", SortOrder: 6},
	}
}

// SeedSystemCategories 初始化系统预设分类
func SeedSystemCategories(db *gorm.DB) error {
	systemCategories := GetSystemCategories()
	categoryMap := make(map[string]uint) // 用于存储分类名称到ID的映射

	// 首先创建所有一级分类
	for _, category := range systemCategories {
		if category.ParentName == "" { // 一级分类
			var existingCategory ClothingCategory
			err := db.Where("name = ? AND parent_id IS NULL", category.Name).First(&existingCategory).Error

			if err == gorm.ErrRecordNotFound {
				newCategory := ClothingCategory{
					Name:        category.Name,
					Description: category.Description,
					Icon:        category.Icon,
					SortOrder:   category.SortOrder,
					IsActive:    true,
				}

				if err := db.Create(&newCategory).Error; err != nil {
					return err
				}
				categoryMap[category.Name] = newCategory.ID
			} else {
				categoryMap[category.Name] = existingCategory.ID
			}
		}
	}

	// 然后创建所有二级分类
	for _, category := range systemCategories {
		if category.ParentName != "" { // 二级分类
			parentID, exists := categoryMap[category.ParentName]
			if !exists {
				continue // 父分类不存在，跳过
			}

			var existingCategory ClothingCategory
			err := db.Where("name = ? AND parent_id = ?", category.Name, parentID).First(&existingCategory).Error

			if err == gorm.ErrRecordNotFound {
				newCategory := ClothingCategory{
					Name:        category.Name,
					Description: category.Description,
					ParentID:    &parentID,
					Icon:        category.Icon,
					SortOrder:   category.SortOrder,
					IsActive:    true,
				}

				if err := db.Create(&newCategory).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// GetSystemCategoriesByParent 根据父分类获取子分类
func GetSystemCategoriesByParent(parentName string) []SystemCategory {
	allCategories := GetSystemCategories()
	var filteredCategories []SystemCategory

	for _, category := range allCategories {
		if category.ParentName == parentName {
			filteredCategories = append(filteredCategories, category)
		}
	}

	return filteredCategories
}
