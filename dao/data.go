package dao

import (
	"blogs_crawler/global"
	"blogs_crawler/model"
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 爬去结果存入数据库
func InsertData(info *model.ArticleInfo) {
	db, err := sql.Open("mysql", "root:111111Aa@/blogs_crawler?charset=utf8")
	global.CheckErr(err)
	defer func() {
		log.Println("close database")
		db.Close()
	}()
	sql := "INSERT INTO article_info(name,url) VALUES(?,?)"
	log.Println(strings.TrimSpace(info.Name))
	result, err := db.Exec(sql, strings.TrimSpace(info.Name), strings.TrimSpace(info.Url))
	global.CheckErr(err)
	id, _ := result.LastInsertId()
	log.Println(id)
}
