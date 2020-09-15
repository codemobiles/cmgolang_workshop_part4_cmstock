package api

import (
	"fmt"
	"main/db"
	_"main/interceptor"
	"main/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product" /*interceptor.JwtVerify,*/, getProduct)
		productAPI.GET("/product/:id" /*interceptor.JwtVerify,*/, getProductByID)
		productAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)
		productAPI.PUT("/product" /*interceptor.JwtVerify,*/, editProduct)
	}
}

func getProduct(c *gin.Context) {
	var product []model.Product
	db.GetDB().Find(&product)
	c.JSON(200, product)
}


func getProductByID(c *gin.Context) {
	var product model.Product
	db.GetDB().Where("id = ?", c.Param("id")).First(&product)
	c.JSON(200, product)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&product).Update("image", fileName)
	}
}

func createProduct(c *gin.Context) {

	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDB().Create(&product)
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)

	c.JSON(http.StatusOK, gin.H{"result": product})

}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	db.GetDB().Save(&product)
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})

}
