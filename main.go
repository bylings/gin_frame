package main

import (
	_ "github.com/bylings/gin/app"
	_ "github.com/bylings/gin/config"
	"github.com/bylings/gin/routes"
)

func main() {
	//config, err := ini.Load("./config/config.ini.example")
	//if err != nil {
	//	//失败
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exit(1)
	//}
	//fmt.Println("config",config)
	// 热加载   bee run
	routes.InitRouter().Run(":8081")
}
