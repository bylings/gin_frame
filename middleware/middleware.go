package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin框架的中间件
// 定位：全局
// 位置：所有路由执行前
func InitApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("InitApp ====== >>>  start")
		fmt.Println("InitApp ====== >>>  end")
	}
}

// gin框架的中间件
// 定位：全局
// 执行：所有路有前  =》 执行请求  =》所有路由后
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Login ====== >>>  start")
		c.Next()
		fmt.Println("Login ====== >>>  end")
	}
}

// gin框架中的中间件
// 定位：局部，因使用关系设定
// 执行：执行：所有路有前  =》 执行请求  =》所有路由后
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Auth ====== >>>  start")
		c.Next()
		fmt.Println("Auth ====== >>>  end")
	}
}
