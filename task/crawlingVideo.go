package task

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron/v3"
	"go_gin/app/model/videos"
	"go_gin/dao/connect"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var baseUrl = "http://ais97.com/"
var videoBaseUrl = "http://ais96.com"

func Crawlings() {

	//uri, err := url.Parse("http://127.0.0.1:10809")
	//if err != nil {
	//	log.Println(err)
	//}
	client := http.Client{
		//Transport: &http.Transport{
		//	Proxy: http.ProxyURL(uri),
		//},
	}
	res, err := client.Get("http://ais97.com/shipin/list-all.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("status code error: %d %s %v", res.StatusCode, res.Status, "列表")
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}
	doc.Find(".hy-video-list .item .clearfix li").Each(func(i int, selection *goquery.Selection) {
		v := connect.Videos{}
		v.Title, _ = selection.Find(".videopic").Attr("title")
		v.VideoClass = selection.Find(".textbg").Text()
		v.ImgUri, _ = selection.Find(".videopic").Attr("data-original")
		href, _ := selection.Find(".videopic").Attr("href")
		go GetTimer(baseUrl+href, &client, v)
	})


}

func GetTimer(href string, client *http.Client, v connect.Videos) {

	res, err := client.Get(href)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("status code error: %d %s %v", res.StatusCode, res.Status, "详情")
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	v.Timer = doc.Find(".hy-video-details .clearfix  .content .clearfix ul li").Eq(1).Text()
	videoHref, _ := doc.Find(".hy-layout  .tab-content .hy-play-list .playlistlink-4 tbody tr td a").Attr("href")

		go GetVideoUri(baseUrl+videoHref, client, v )
}

func GetVideoUri(href string, client *http.Client, v connect.Videos) {

	res, err := client.Get(href)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("status code error: %d %s %v", res.StatusCode, res.Status, "播放")
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	html := doc.Find("script").Last()
	htmls := html.Text()
	reg := regexp.MustCompile(`video = '(.*?)';`)
	x := reg.FindStringSubmatch(htmls)
	uri := x[1]
	if ok := strings.HasSuffix(uri, ".m3u8"); ok {
		uri = videoBaseUrl + uri
	}
	v.VideoUri = uri


	videos.GetVideo(&v)
}


func InitTasks() {
	c := cron.New(cron.WithSeconds())
	defer c.Stop()

	c.AddFunc("0 */59 * * * ?", Crawlings)
	go c.Start()
}
