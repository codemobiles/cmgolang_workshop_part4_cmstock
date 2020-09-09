package api

import (
	"main/db"
	"main/model"
	"time"
	_ "time"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "login"})
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {		
		// user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()						
		if err := db.GetDB().Create(&user).Error; err != nil {			
			c.JSON(200, gin.H{"result": "nok", "error": err})
		}else{
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}		
}
