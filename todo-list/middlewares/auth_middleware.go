package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "todo-list/utils"
    "strings"
)




func AuthMiddleware(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
        c.Abort()
        return
    }
    // 确保 token 具有 "Bearer " 前缀
    if !strings.HasPrefix(tokenString, "Bearer ") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // 去掉 "Bearer " 前缀
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")
    userID, err := utils.ParseToken(tokenString)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    c.Set("user_id", userID)
    c.Next()
}
