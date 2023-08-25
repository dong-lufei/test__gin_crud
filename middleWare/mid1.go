package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义中间件1
func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("调用中间件1")
	}
}
