package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "todo-list/models"
    "todo-list/config"  // 导入 config 包，获取数据库实例
)

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    userID, _ := c.Get("user_id")
    config.DB.Where("user_id = ?", userID).Find(&todos)  // 使用 config.DB
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var input struct {
        Title       string `json:"title"`
        Description string `json:"description"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("user_id")
    todo := models.Todo{Title: input.Title, Description: input.Description, UserID: userID.(uint)}
    config.DB.Create(&todo)  // 使用 config.DB
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")
    
    if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    var input struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Status      string `json:"status"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo.Title = input.Title
    todo.Description = input.Description
    todo.Status = input.Status

    config.DB.Save(&todo)

    c.JSON(http.StatusOK, todo)
}

// 删除待办事项
func DeleteTodo(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")

    if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    config.DB.Delete(&todo)

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}