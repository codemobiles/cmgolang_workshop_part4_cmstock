package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func myInterceptor(c *gin.Context){
	token := c.Query("token")
	if (token == "1234"){
		c.Next()
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product",myInterceptor, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}
