package task

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go_gin/dao/connect"
	"log"
	"net/http"
)

func InitDbDatas(){
	t := &connect.Videos{}
	tb := connect.Dbs.Mysql.First(t)
	if tb.Error != nil {
		log.Println(tb.Error)
	}
	if tb.RowsAffected < 1 {
		for i := 1; i < 300; i++ {
			//uri, err := url.Parse("http://127.0.0.1:10809")
			//if err != nil {
			//	log.Println(err)
			//}
			client := http.Client{
				//Transport: &http.Transport{
				//	Proxy: http.ProxyURL(uri),
				//},
			}
			res, err := client.Get("http://ais97.com/shipin/list-all-insert_time-"+ fmt.Sprint(i) + ".html")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s %v", res.StatusCode, res.Status, "列表")
				return
			}
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find(".hy-video-list .item .clearfix li").Each(func(i int, selection *goquery.Selection) {
				v := connect.Videos{}
				v.Title, _ = selection.Find(".videopic").Attr("title")
				v.VideoClass = selection.Find(".textbg").Text()
				v.ImgUri, _ = selection.Find(".videopic").Attr("data-original")
				href, _ := selection.Find(".videopic").Attr("href")
				go GetTimer(baseUrl+href, &client, v )
			})
		}
	}
}
