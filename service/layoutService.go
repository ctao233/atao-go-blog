package service

import (
	"atao-go-blog/dao"
	"atao-go-blog/vo"
	"log"
)

func Layout() *vo.LayoutRes {

	blogs, _ := dao.GetAllBlog(0, 3)
	blogCout := dao.GetBlogCount()
	tags, _ := dao.GetAllTag()
	tagCount := dao.GetTagCount()
	categorycount := dao.GetCategoryCount()
	site := dao.GetSite()
	infor, _ := dao.GetInformation()
	footer, _ := dao.GetFootInfo()

	//  右侧扩展数据
	var Right vo.Right
	Right.HotArticles = dao.GetHotArticle()

	m, err := dao.GetCategoryBlogCount()
	if err != nil {
		log.Println("数据组合错误")
		return nil
	}
	result := &vo.LayoutRes{

		Avatar:  infor.Photo,
		Name:    infor.Name,
		Sentenc: infor.Sentence,

		BlogCount:     blogCout,
		TagCount:      tagCount,
		CategoryCount: categorycount,
		Categories:    m,

		Github: infor.Github,
		Email:  infor.Email,
		Site:   site,
		Blogs:  blogs,
		Tags:   tags,
		Foot:   footer,
		Right:  Right,
	}
	return result
}
