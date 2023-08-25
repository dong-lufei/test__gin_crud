package shop

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/zoos", getZoos)
	e.POST("/zoos", addZoos)
}
