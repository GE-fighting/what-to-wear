package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"what-to-wear/server/config"
	"what-to-wear/server/database"
)

func main() {
	// 定义命令行参数
	var (
		action = flag.String("action", "migrate", "操作类型: migrate, seed, reset, status, drop")
		seeder = flag.String("seeder", "", "指定要运行的种子数据 (categories, tags)")
	)
	flag.Parse()

	// 加载配置并连接数据库
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatalf("加载数据库配置失败: %v", err)
	}

	db, err := config.ConnectDatabase(dbConfig)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 根据操作类型执行相应的命令
	switch *action {
	case "migrate":
		fmt.Println("开始数据库迁移...")
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("数据库迁移失败: %v", err)
		}
		fmt.Println("数据库迁移完成!")

	case "seed":
		fmt.Println("开始运行种子数据...")
		if *seeder == "" {
			// 运行所有种子数据
			if err := database.RunAllSeeders(db); err != nil {
				log.Fatalf("运行种子数据失败: %v", err)
			}
		} else {
			// 运行指定的种子数据
			if err := database.RunSeeder(db, *seeder); err != nil {
				log.Fatalf("运行种子数据失败: %v", err)
			}
		}
		fmt.Println("种子数据运行完成!")

	case "reset":
		fmt.Println("开始重置数据库...")
		if err := database.DropAllTables(db); err != nil {
			log.Fatalf("删除表失败: %v", err)
		}
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("重新迁移失败: %v", err)
		}
		if err := database.RunAllSeeders(db); err != nil {
			log.Fatalf("重新初始化种子数据失败: %v", err)
		}
		fmt.Println("数据库重置完成!")

	case "status":
		fmt.Println("检查数据库状态...")
		if err := database.CheckMigrationStatus(db); err != nil {
			log.Fatalf("数据库状态检查失败: %v", err)
		}
		fmt.Println("数据库状态正常!")

	case "drop":
		fmt.Println("警告: 即将删除所有表!")
		fmt.Print("确认删除? (y/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" {
			if err := database.DropAllTables(db); err != nil {
				log.Fatalf("删除表失败: %v", err)
			}
			fmt.Println("所有表已删除!")
		} else {
			fmt.Println("操作已取消")
		}

	default:
		fmt.Printf("未知操作: %s\n", *action)
		fmt.Println("可用操作: migrate, seed, reset, status, drop")
		os.Exit(1)
	}

	os.Exit(0)
}