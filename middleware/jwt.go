package middleware

import (
	"strings"
	"net/http"
	"gokit/services"
	"github.com/gin-gonic/gin"
)

//TODO: check if JWT token exists on cache.
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		Authorization := c.Request.Header.Get("Authorization")
		if Authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Authorization not found"})
			c.Abort()
			return
		}

		AuthorizationValues := strings.Split(Authorization, "Bearer ");

		if  len(AuthorizationValues) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "Invalid Token Format"})
			c.Abort()
			return
		}
		
		if err := services.ValidateToken(strings.TrimSpace(AuthorizationValues[1])); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
