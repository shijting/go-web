package middlewares

import "github.com/gin-gonic/gin"

func CheckAuthorization() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
