package config

import (
	"gopkg.in/ini.v1"
	"strings"
)

var AppConfig *ini.File

func init() {
	AppConfig, _ = ini.Load("./config.ini")
	//if err != nil {
	//	//失败
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exit(1)
	//}
	//fmt.Println("AppConfig   ", AppConfig)
	////获取ini里面的配置
	//fmt.Println("App Mode:", config.Section("").Key("app_name").String())
	//fmt.Println("App Mode:", AppConfig.Section("mysql").Key("host").String())
	//fmt.Println("App Mode:", AppConfig.Section("redis").Key("ip").String())
	////给ini写入数据
}

func GetConfig(section, key string) string {
	return strings.Trim(string(AppConfig.Section(section).Key(key).String()), "")
}
