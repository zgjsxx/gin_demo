package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/hello",helloHandler)
	e.GET("/comment",commentHandler)
}