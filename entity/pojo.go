package entity

import (
	"time"
)

/**
  Description: 个人博客表
 **/

type Blog struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Photo      string    `json:"photo"`
	View       int64     `json:"view"`
	Tags       string    `json:"tags"`
	Catalog    string    `json:"catalog"`
	Content    string    `json:"content"`
	UpLoadTime time.Time `json:"upLoadTime"`
	UpdateTime time.Time `json:"updateTime"`
	YearMonth  string
}

/*
 * Description: 个人目录表
 */
type Category struct {
	Id   int
	Name string
}

/**
 * Description: 评论配置信息表
 */

type Comment struct {
	Id          int
	Flag        bool
	AppId       string
	AppKey      int
	Placeholder string
	PlacePhoto  string
}

/**
关于个人描述页面表
*/

type Description struct {
	Id            int
	Name          string
	Constellation string
	Experience    string
	Sentence      string
	Idol          string
	Character     string
	Contact       string
	Information   string
	Language      string
}

/**
 * 页脚底部信息表
 */
type Footer struct {
	Id        int
	About     string
	Number    string
	Copyright string
	Powerby   string
	Byurl     string
}

/**
 * Description: 侧边栏的个人信息表
 */
type Information struct {
	Id       int
	Photo    string
	Name     string
	Sentence string
	Email    string
	Github   string
}

/**
 * Description: 友情链接表
 */
type Link struct {
	Id         int
	Hphoto     string
	Url        string
	Name       string
	Introduce  string
	CreateTime time.Time
}

/*
	Description:页面的展示图像
*/
type Photo struct {
	Id           int
	LinkPhoto    string
	AboutPhoto   string
	TagPhoto     string
	IndexPhoto   string
	CatalogPhoto string
}

/*
 * Description: 站点信息表
 */
type Site struct {
	Id          int
	Name        string
	Description string
	Logo        string
	Notice      string
	Favicon     string
	BlogLink    string
}

/*
 * Description: 标签表
 */

type Tag struct {
	Id         int
	Name       string
	CreateTime time.Time
	Deleted    int
	UpdateTime time.Time
}

/*
 * Description: 个人用户表
 */
type User struct {
	Id       int
	Account  string
	Password string
	Problem  string
	Answer   string
}
