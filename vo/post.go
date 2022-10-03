package vo

import (
	"atao-go-blog/entity"
	"time"
)

// 导航 ，页脚 , 左侧需要的公共数据 以及一些不常变化的数据
type LayoutRes struct {
	Avatar  string
	Name    string
	Sentenc string

	BlogCount     int //  博客数量
	TagCount      int // 标签数量
	CategoryCount int // 分类数量

	Github string
	Email  string

	Site       entity.Site
	Blogs      []entity.Blog
	Categories map[string]int
	Tags       []entity.Tag
	Foot       entity.Footer
	Data       interface{}
	Right      Right
}

// 分页数据Api
type PageResult struct {
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
	PageCount int         `json:"pageCount"`
	Total     int         `json:"total"`
	Status    int         `json:"status"`
	Result    interface{} `json:"result"`
}

// 归档数据结构
type ArchIvesResul struct {
	Month string
	Blog  entity.Blog
}

// 分类标签返回数据模型
type TagOrCategory struct {
	SearchField string
	Count       int
}

// 文章搜索返回模型
type SearchBlogVo struct {
	Id      int
	Content string
	Title   string
}

// 同一返回
type Result struct {
	Code int
	Msg  string
	Data interface{}
}

// 热门文章 排xiang
type HotArticle struct {
	Id         int
	Title      string
	Photo      string
	CreateDate time.Time
}

type Right struct {
	// 热门文章
	HotArticles []HotArticle
}
