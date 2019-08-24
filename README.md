# 基于go实现的 blogs 信息爬虫项目

# 技术栈选型
go + goquery + colly + mysql + bootstrap + jquery

## 项目结构

## to do list
-  数据库表格的设计(已完成)
-  基于 goquery 爬取 CSDN 和 android 博客内容(已完成)
-  实现自动分页获取数据的能力
-  使用协程并发爬取并控制并发数量同时定时爬取
-  前端页面展示
2.基于goquery的网页爬取实现


func WithRequest(url string) (*http.Response,error) {
	req, err := http.NewRequest("GET",url,nil)
	if err != nil{
		return nil,err
		}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36")
	log.Print("Success Set Header")
	return http.DefaultClient.Do(req)
}
