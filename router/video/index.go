package video

import (
	"github.com/gin-gonic/gin"
	"go_gin/app/api/video"
)

func HandelBindGroup(c *gin.Engine){
	// video组
	g := c.Group("/video"  )


	// 获取首页数据接口
	g.GET("/getHot", video.ApiVideoGetHot)

	// 播放页likeAndMore 查询
	g.POST("/getVideoByIdAndLike", video.ApiVideoGetVideoByIdAndLike)

	// 搜索
	g.POST("/searchByLike", video.ApiVideoSearchByLike)

	// 分类查询
	g.POST("/getVideoByTypeAndPage", video.ApiVideoGetVideoByTypeAndPage)
}