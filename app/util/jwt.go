package util

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 获取userId
func GetUserId(c *gin.Context) uint {
	claims := jwt.ExtractClaims(c)
	userId, ok := claims["userId"]
	if ok {
		return uint(userId.(float64))
	}

	return 0
}

// 获取auth
func GetAuth(c *gin.Context) uint {
	claims := jwt.ExtractClaims(c)
	auth, ok := claims["auth"]
	if ok {
		return uint(auth.(float64))
	}

	return 0
}
