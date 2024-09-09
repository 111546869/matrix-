package controllers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "todo-list/models"
    "todo-list/config"  // 导入 config 包，获取数据库实例
    "todo-list/utils"
	"fmt"
)

func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 使用 bcrypt 对密码进行哈希加密
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := models.User{Username: input.Username, Password: string(hashedPassword)}

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


func Login(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // 输出用于调试的密码信息
    fmt.Println("Stored password hash:", user.Password)
    fmt.Println("Input password:", input.Password)

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
