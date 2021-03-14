package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_gin/cors"
	"go_gin/dao/connect"
	"go_gin/router"
	"go_gin/task"
)



func main(){
	s := gin.Default()

	// 注册中间件
	s.Use(cors.Cors())
	//s.Use(infoMiddleware.PrintInfo())

	// 设置静态目录
	s.Static("/static", "./static")


	// 初始化连接数据库
	connect.InitDbConnect()


	// 初始化路由
	router.InitGinRouter(s)

	// 初始化定时任务
	task.InitTasks()
	// 数据库无数据时开启爬取
	go task.InitDbDatas()


	// 开启服务
	err := s.Run(":3000")
	if err != nil {
		fmt.Println("错误")
		return
	}
}
