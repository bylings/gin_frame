package routes

import (
	"github.com/bylings/gin/middleware"
	"github.com/gin-gonic/gin"
)

// 约定自动注册的类型
type Router func(engine *gin.Engine)

var routers = []Router{}

func RegisterRoute(routes ...Router) {
	routers = append(routers, routes...)
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.InitApp())
	//r.Use(middleware.Login())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index",
			"code":    200,
		})
	})

	r.GET("/user", middleware.Auth(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user",
			"code":    200,
		})
	})
	for _, route := range routers {
		route(r) // 加载route
	}
	return r
}
