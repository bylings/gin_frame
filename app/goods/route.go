package goods

import (
	"github.com/bylings/gin/routes"
	"github.com/gin-gonic/gin"
)

func init() { // 初始化的时候注册
	routes.RegisterRoute(Routes)
}

func Routes(g *gin.Engine) {

	goods := g.Group("goods")
	{
		goods.GET("get", GetGoods)
		goods.POST("add", AddGoods)
		goods.POST("save", SaveGoods)
	}

}