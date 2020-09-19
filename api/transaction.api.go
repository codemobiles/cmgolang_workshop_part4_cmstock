package api

import (
	"main/db"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine)  {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)
	}
}


func getTransaction(c *gin.Context) {
	c.String(http.StatusOK, "List Transaction")
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction	
	if err := c.ShouldBind(&transaction); err == nil {		
		transaction.CreatedAt = time.Now()		
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(404, gin.H{"result": "nok"})
	}
}