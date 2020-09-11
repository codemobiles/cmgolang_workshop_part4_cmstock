package api

import (
	"main/interceptor"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.GeneralInterceptor1, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}
