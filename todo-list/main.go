package main

import (
    "github.com/gin-gonic/gin"
    "todo-list/controllers"
    "todo-list/middlewares"
    "todo-list/config"
    "todo-list/models"
    "time"
)

func main() {
    time.Local = time.FixedZone("CST", 8*3600)
    r := gin.Default()

    // 初始化数据库
    config.ConnectDatabase()  
    config.DB.AutoMigrate(&models.Todo{})
    r.POST("/users/register", controllers.Register)
    r.POST("/users/login", controllers.Login)

    protected := r.Group("/todos")
    protected.Use(middlewares.AuthMiddleware)
    {
        protected.GET("/", controllers.GetTodos)
        protected.POST("/", controllers.CreateTodo)
        protected.PUT("/:id", controllers.UpdateTodo)
        protected.DELETE("/:id", controllers.DeleteTodo)
    }

    r.Run()
}
