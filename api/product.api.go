package api

import (
	"main/interceptor"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product", "username": c.GetString("jwt_username"), "level": c.GetString("jwt_level")})
}

func createProduct(c *gin.Context) {

	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	image, _ := c.FormFile("image")	
	product.Image = image.Filename
	c.JSON(200, gin.H{"result": product})

}
