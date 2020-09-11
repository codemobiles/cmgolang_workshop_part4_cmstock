package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor1 - call this methos to add interceptor
func GeneralInterceptor1(c *gin.Context){
	token := c.Query("token")
	if (token == "1234"){
		c.Next()
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}