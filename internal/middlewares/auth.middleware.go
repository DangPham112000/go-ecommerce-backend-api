package middlewares

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
