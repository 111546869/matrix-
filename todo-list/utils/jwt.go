package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "fmt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
    UserID uint
    jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)

    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (uint, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        fmt.Println("Token parsing error:", err)
        return 0, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims.UserID, nil
    } else {
        return 0, err
    }
}
