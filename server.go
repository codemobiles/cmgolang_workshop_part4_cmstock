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

	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port In Heroku")
		router.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		router.Run(port)
	}
}
