package util

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserEmailFromClaims(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	email, ok := claims["email"]
	if ok {
		return email.(string)
	}

	return "no email"
}
