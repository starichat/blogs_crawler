package crawel

import (
	"blogs_crawler/dao"
	"blogs_crawler/global"
	"blogs_crawler/model"
	"blogs_crawler/util"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)


var info = &model.ArticleInfo{}

// 基于goquery爬虫页面实现
// 获取gityuan博客文章
func Getartices(url string) {

	// 请求
	resp, err := util.GetRequest(url)
	global.CheckErr(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	global.CheckErr(err)
	log.Println("parse response successful!")
	// 解析文章列表数据
	doc.Find(".container .post-preview").Each(func(i int, s *goquery.Selection) {
		info.Name = s.Find("a").Text()
		info.Url, _ = s.Find("a").Attr("href")
		if !strings.Contains(info.Url, "http") {
			info.Url = "http://gityuan.com" + info.Url
		}
		// 插入数据库
		dao.InsertData(info)
	})


}

func GetarticesForCSDN(url string) {
	resp, err := util.GetRequest(url)
	global.CheckErr(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	global.CheckErr(err)
	log.Println("parse response successful!")
	doc.Find("main .article-list .article-item-box").Each(func(i int, s *goquery.Selection) {
		info.Name = s.Find("h4").Find("a").Text()
		span := s.Find("h4").Find("a").Find("span").Text()
		info.Url, _ = s.Find("h4").Find("a").Attr("href")
		// 判断文章标题是否有<span>的内容，若有则去掉
		info.Name = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(info.Name), span))

		dao.InsertData(info)

	})

}

