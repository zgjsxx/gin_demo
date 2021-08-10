package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello handler",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "comment handler",
	})
}