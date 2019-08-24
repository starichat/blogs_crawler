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

## 开发流程
1. 分析要获取的数据以及要展示的效果，最终简单抽取出要获取如下表格的数据展示

|Id | 文章名称 | 文章url地址 |
|:---:|:---:|:---:|
| 01 | 基于Goland的blogs爬虫项目|http://starichat.pro/*****|

基于此设计数据库sql语句如下：
```cassandraql

CREATE DATABASE IF NOT EXISTS `blogs_crawler` DEFAULT CHARACTER SET utf8mb4;

USE `blogs_crawler`;

CREATE TABLE IF NOT EXISTS `article_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文章名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '文章url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='文章信息表';
```


