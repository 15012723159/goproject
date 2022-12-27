package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"grpc/web/controller"
	"grpc/web/model"
)

func main() {

	// 初始化数据库链接
	model.InitDb()
	router := gin.Default()

	r1 := router.Group("api/v1/")
	{
		r1.GET("index/index", controller.Index)

		r1.POST("user/avatar", controller.PostUserAvatar)

		r1.POST("user/find", controller.SelectUser)
	}

	router.Run(":8088")
}
