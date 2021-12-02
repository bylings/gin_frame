package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dns string = "root:root@tcp(192.168.86.129:3306)/go?charset=utf8"
)

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "", // 表前缀
			SingularTable: false, // 是否开启表前缀,true开启，false关闭
		},
	})
	fmt.Println("err ", err)
	return db
}
