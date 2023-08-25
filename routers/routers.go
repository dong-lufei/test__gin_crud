package routers

import (
	middleware "test/middleWare"

	"time"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.Default()
	// r := gin.New()
	// gin.SetMode(gin.ReleaseMode)

	// 使用自定义中间件
	r.Use(middleware.MiddleWare1(), middleware.MiddleWare2())

	r.Static("/static", "./assets")

	// 加载html模板路径
	r.LoadHTMLGlob("template/*")
	r.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.HTML(200, "index.html", gin.H{"title": "首页", "home": "访问成功"})
	})
	for _, opt := range options {
		opt(r)
	}
	return r
}
