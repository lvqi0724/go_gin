package video

import (
	"github.com/gin-gonic/gin"
	"go_gin/dao/connect"
	"log"
	"math/rand"
	"net/http"
	"time"
)


// 获取首页数据
func ApiVideoGetHot(c *gin.Context){
	t := make([]connect.Videos, 20)
	tb := connect.Dbs.Mysql.Order("timer desc").Limit(20).Find(&t)
	if tb.Error != nil {
		log.Println(tb.Error)
		c.JSON(500, gin.H{
			"code": 500,
			"msg": "内部服务器出错",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": t,
	})
}

// 播放页 like
func ApiVideoGetVideoByIdAndLike(c *gin.Context){
	id := make(map[string]int)
	err := c.BindJSON(&id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg": "ID无效",
		})
		c.Abort()
		return
	}
	t := &connect.Videos{}
	tx := connect.Dbs.Mysql.First(&t, id["id"])
	if tx.Error != nil {
		log.Println(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg": "不存在的Id",
		})
		c.Abort()
		return
	}
	str := []rune(t.Title)
	rand.Seed(time.Now().Unix())
	len := rand.Intn(len(str))
	tt := make([]connect.Videos, 10)
	connect.Dbs.Mysql.Where("Title LIKE  ?", "%" + string(str[len])+ "%").Limit(10).Find(&tt)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": t,
		"more": tt,
	})
}


// 搜素

func ApiVideoSearchByLike(c *gin.Context)  {
	like := make(map[string]string)
	err := c.ShouldBindJSON(&like)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "无查询",
		})
		c.Abort()
		return
	}
	tt := make([]connect.Videos, 10)
	ttt := make([]connect.Videos, 10)
	connect.Dbs.Mysql.Where("Title LIKE ?", "%" + like["like"] + "%").Find(&tt)
	rand.Seed(time.Now().Unix())
	for i:=0; i < 10; i++ {
		ttt[i] = tt[rand.Intn(len(tt))]
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": ttt,
	})
}


func ApiVideoGetVideoByTypeAndPage(c *gin.Context){
	obj := make(map[string]interface{})
	err := c.BindJSON(&obj)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": "查询数据错误",
		})
		c.Abort()
		return
	}
	if obj["type"] != "" && obj["limit"].(float64) > 0 {
		t := make([]connect.Videos, 20)
		tx := connect.Dbs.Mysql.Where("video_class = ?", obj["type"]).Order("timer DESC").Offset((int(obj["limit"].(float64)) -1 ) * 20).Limit(20).Find(&t)
		var count int64
		txx := connect.Dbs.Mysql.Model(&connect.Videos{}).Where("video_class = ?", obj["type"]).Count(&count)
		if txx.Error != nil || tx.Error != nil{
			log.Println(tx.Error)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg": "查询数据错误",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": t,
			"count": count,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg": "查询数据错误",
	})
}