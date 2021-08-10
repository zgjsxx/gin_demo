package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func blogHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello blog handler",
	})
}