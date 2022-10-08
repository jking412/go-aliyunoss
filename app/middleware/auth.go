package middleware

import (
	"aliyunoss/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.AuthJWT(c.Request)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "token错误",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.JWTCustomClaims); ok && token.Valid {
			c.Set("userId", claims.UserId)
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "token错误",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
	}
}
