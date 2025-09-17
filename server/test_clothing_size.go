package main

import (
	"fmt"
	"what-to-wear/server/api"
	"what-to-wear/server/models"
)

// 测试 ClothingItem 的 Size 字段简化
func main() {
	// 测试创建衣物实例
	item := &models.ClothingItem{
		UserID:      1,
		CategoryID:  1,
		Name:        "测试T恤",
		Brand:       "优衣库",
		Color:       "白色",
		Size:        "XL", // 现在是简单的字符串
		Material:    "棉",
		Description: "简单舒适的T恤",
		Price:       99.00,
		Condition:   api.ClothingStatusActive,
		WearCount:   0,
		IsActive:    true,
		IsFavorite:  false,
	}

	// 测试不同的尺码格式
	testSizes := []string{
		"XS", "S", "M", "L", "XL", "XXL", // 通用尺码
		"38", "39", "40", "41", "42", // 欧码/中国码
		"8", "10", "12", "14", // 美码/英码
		"One Size", // 均码
	}

	fmt.Println("=== 测试衣物尺码字段简化 ===")
	fmt.Printf("衣物名称: %s\n", item.Name)
	fmt.Printf("当前尺码: %s\n", item.Size)
	fmt.Println()

	fmt.Println("=== 测试不同尺码格式 ===")
	for _, size := range testSizes {
		item.Size = size
		fmt.Printf("尺码: %-8s -> 存储值: %s\n", size, item.Size)
	}

	fmt.Println()
	fmt.Println("=== 尺码字段类型验证 ===")
	fmt.Printf("Size 字段类型: %T\n", item.Size)
	fmt.Printf("Size 字段值: %q\n", item.Size)

	// 测试尺码的智能识别
	fmt.Println()
	fmt.Println("=== 尺码系统智能识别示例 ===")
	sizeExamples := map[string]string{
		"XS":       "通用尺码系统",
		"S":        "通用尺码系统",
		"M":        "通用尺码系统",
		"L":        "通用尺码系统",
		"XL":       "通用尺码系统",
		"38":       "欧码/中国码",
		"39":       "欧码/中国码",
		"40":       "欧码/中国码",
		"8":        "美码/英码",
		"10":       "美码/英码",
		"12":       "美码/英码",
		"One Size": "均码",
	}

	for size, system := range sizeExamples {
		fmt.Printf("尺码 %-8s -> 识别为: %s\n", size, system)
	}

	fmt.Println()
	fmt.Println("✅ 尺码字段简化测试完成！现在可以直接使用字符串表示尺码，无需复杂的结构体。")
}
