package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收注册数据的结构体
type Register struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段;还可验证字段
	Username   string `form:"username" json:"username" uri:"username" xml:"username" binding:"gte=3,lte=12"`
	Password   string `form:"password" json:"password" uri:"password" xml:"password" binding:"required,gte=6,lte=12"`
	RePassword string `form:"rePassword" json:"rePassword" uri:"rePassword" xml:"rePassword" binding:"eqfield=Password"`
	Email      string `form:"email" json:"email" uri:"email" xml:"email" binding:"required,email"`
}

// 定义接收登录数据的结构体
type Login struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	// Phone string `form:"phone" json:"phone" uri:"phone" xml:"phone" binding:"e164"`
	// Email string `form:"email" json:"email" uri:"email" xml:"email" binding:"required,email"`
}

// 登录接口:JSON数据请求
func loginJSON(c *gin.Context) {
	// 声明接收的变量
	var userJSON Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&userJSON); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if userJSON.Username != "root" || userJSON.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 登录接口:表单数据请求
func loginForm(c *gin.Context) {
	// 声明接收的变量
	var userForm Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if userForm.Username != "root" || userForm.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 登录接口:Uri方式
func loginUri(c *gin.Context) {
	// 声明接收的变量
	var userUri Login
	// 根据请求头中content-type自动推断
	if err := c.ShouldBindUri(&userUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if userUri.Username != "root" || userUri.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func Routers(e *gin.Engine) {
	e.POST("/loginJSON", loginJSON)
	e.POST("/loginForm", loginForm)
	e.GET("/:username/:password", loginUri)
}
