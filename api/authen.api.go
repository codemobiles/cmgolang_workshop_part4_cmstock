package api

import (
	"main/db"
	"main/model"
	"time"
	_ "time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
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
	var user model.User
	
	if c.ShouldBind(&user) == nil {
		var queryUser model.User		
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {							
			c.JSON(200, gin.H{"result": "nok", "error": err})
		}else if (checkPasswordHash(user.Password, queryUser.Password) == false){
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		}else{		

			atClaims := jwt.MapClaims{}

			// Payload begin
			atClaims["id"] = queryUser.ID
			atClaims["username"] = queryUser.Username
			atClaims["level"] = queryUser.Level
			atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
			// Payload end

			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			token, _ := at.SignedString([]byte("1234"))

			c.JSON(200, gin.H{"result": "ok", "token": token})
		}
		
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}	
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {		
		user.Password, _ = hashPassword(user.Password)
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



func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}