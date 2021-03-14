package router

import (
	"github.com/gin-gonic/gin"

	"go_gin/router/video"
)



// 初始化路由
func InitGinRouter (s *gin.Engine){

	// 注册路由

	video.HandelBindGroup(s)

}

