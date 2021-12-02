package goods

import (
	"fmt"
	"github.com/bylings/gin/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Goods struct {
	Id         int
	GoodsName  string `form:"name" json:"name" binding:"required"`
	GoodsPrice int    `form:"price" json:"price" binding:"required,GoodsCheckPrice"`
	GoodsNum   int    `form:"num" json:"num" binding:"required,gt=5"`
}

// 自定义验证器（价格验证器）
var GoodsCheckPrice validator.Func = func(fl validator.FieldLevel) bool {
	fmt.Println("fl ========>> start ==========>>")
	fieldValue := fl.Field().Interface().(int) // 根据验证的字段进行定义
	fmt.Println("field : ", fieldValue)
	// todo 做验证逻辑
	return true
}

func init() {
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("GoodsCheckPrice", GoodsCheckPrice)
	}
}

func GetGoods(c *gin.Context) {
	id := c.Query("id")
	retionInfo := services.GetRegionsDetail(map[string]interface{}{
		"id": id,
	})
	c.JSON(200, gin.H{
		"code":    200,
		"message": "goods index",
		"data":    retionInfo,
	})
}

// json参数
func AddGoods(c *gin.Context) {
	var json Goods
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "AddGoods",
		"json":    json,
	})
}

// 表单参数
func SaveGoods(c *gin.Context) {
	var form Goods
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "AddGoods",
		"json":    form,
	})
}
