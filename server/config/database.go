package config

import (
    "fmt"
    "log"
    "os"
    "what-to-wear/server/models"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // 加载 .env 文件
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

    // 从环境变量读取数据库配置
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")

    // 设置默认值
    if host == "" {
        host = "localhost"
    }
    if user == "" {
        user = "admin"
    }
    if password == "" {
        password = "admin"
    }
    if dbname == "" {
        dbname = "what_to_wear"
    }
    if port == "" {
        port = "5432"
    }
    if sslmode == "" {
        sslmode = "disable"
    }

    // 构建数据库连接字符串
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode)

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 自动迁移
    err = database.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    DB = database
    fmt.Println("Database connected successfully")
}