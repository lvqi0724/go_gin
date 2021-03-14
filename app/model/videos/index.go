package videos

import (
	"go_gin/dao/connect"
	"log"
)

// 获取是否存在 不存在则插入数据库
func GetVideo(value *connect.Videos){
	v := &connect.Videos{}
	tb := connect.Dbs.Mysql.Where("title = ?",value.Title ).Find(v)
	if tb.Error != nil {
		log.Println(tb.Error)
		return
	}
	if v.Title != "" {
		return
	}
	InsetVideo(value)
}

// 插入 接受指针
func InsetVideo (value *connect.Videos){
	tb := connect.Dbs.Mysql.Create(value)
	if tb.Error != nil {
		log.Println("插入失败", tb.Error)
		return
	}
}