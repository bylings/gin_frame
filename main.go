package main

import (
	_ "github.com/bylings/gin/app"
	"github.com/bylings/gin/routes"
)

func main() {
	// 热加载   bee run
	routes.InitRouter().Run(":8081")
}
