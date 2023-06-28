package middleware

import (
	commons "blog-go/src/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		mc, err := commons.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		c.Set("loginUser", mc.UserId)
		c.Next()
	}
}
