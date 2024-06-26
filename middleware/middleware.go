package middleware

import (
	"fmt"
	"net/http"

	token "github.com/alikhanMuslim/ecommerce/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		fmt.Println("checking")
		if ClientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No auth token found"})
			c.Abort()
			return
		}
		fmt.Println("passed")

		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
