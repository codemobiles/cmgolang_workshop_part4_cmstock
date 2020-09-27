package main

import (
	"fmt"
	"main/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")

	api.Setup(router)
	// router.Run(":8081")

	// In case of running on Heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("Running on Heroku using random PORT")
		router.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		router.Run(":" + port)
	}
}
