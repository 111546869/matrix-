package config

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "todo-list/models"  // 导入 models 包
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "pan:123456@tcp(127.0.0.1:3306)/todo_db?parseTime=true&loc=Asia%2FShanghai"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database!", err)
    }

    DB = database
    // config.DB.AutoMigrate(&models.Todo{})
    // 自动迁移数据库模型
    DB.AutoMigrate(&models.User{}, &models.Todo{})
}
