package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DbCon *gorm.DB

func InitDb() error {
	dbCon, error := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if error != nil {
		fmt.Println("数据库链接失败")
		return error

	}
	DbCon = dbCon
	DbCon.DB().SetMaxOpenConns(2000)
	DbCon.DB().SetMaxIdleConns(1000)
	DbCon.SingularTable(true)
	fmt.Println("init mysql success")
	return nil

}
