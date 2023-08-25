package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义中间件2
func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("调用中间件2")
	}
}
