package models

import (
	"fmt"
	"github.com/bylings/gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dns string

func init() {
	host := config.GetConfig("mysql", "host")
	port := config.GetConfig("mysql", "port")
	user := config.GetConfig("mysql", "user")
	password := config.GetConfig("mysql", "password")
	database := config.GetConfig("mysql", "database")
	charset := config.GetConfig("mysql", "charset")
	dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, password, host, port, database, charset)
}

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表前缀
			SingularTable: false, // 是否开启表前缀,true开启，false关闭
		},
	})
	fmt.Println("err ", err)
	return db
}
