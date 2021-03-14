package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Conn struct {
	Mysql *gorm.DB
}

type Videos struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null;unique"`
	Timer string `gorm:"not null"`
	ImgUri string `gorm:"not null"`
	VideoUri string `gorm:"not null"`
	VideoClass string `gorm:"type:varchar(20);not null"`
}

func (Videos)TableName()string{
	return "videos"
}

var Dbs *Conn

// 连接mysql
func (c *Conn) MysqlConnect(){
	// 改成自己的数据库
	dbConfig := "root:@root1234@tcp(127.0.0.1:3306)/colorvideo?charset=utf8&parseTime=true&loc=Asia%2fShanghai"
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	sqlDb, err := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(1000)
	if err != nil {
		log.Println(err)
		return
	}
	c.Mysql = db
}


// 查看连接是否成功
func (c *Conn) MysqlPing(){
	if c.Mysql == nil {
		log.Println("mysql未连接")
		return
	}
	log.Println("mysql已连接")
}


// 初始化连接  未连接进行连接
func InitDbConnect(){
	if Dbs == nil {
		Dbs = new(Conn)
		Dbs.MysqlConnect()
		Dbs.Mysql.AutoMigrate(&Videos{})

	}
}