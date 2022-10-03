package service

import (
	"atao-go-blog/common"
	"atao-go-blog/dao"
	"atao-go-blog/vo"
)

// 分类 业务
func GetblogPageByCategoryService(page, pageSize int, condition string) *vo.PageResult {

	blogs, err := dao.GetAllBlogByCategory(((page - 1) * pageSize), pageSize, condition)

	if err != nil {
		return nil
	}

	// 遍历文章 展示部分内容
	common.WordLimit(blogs)

	// 编写页码逻辑
	total := dao.GetBlogByCategoryCount(condition)
	//总页码数
	pagesCount := (total-1)/5 + 1

	pageResult := &vo.PageResult{
		Page:      page,
		PageSize:  pageSize,
		PageCount: pagesCount,
		Total:     total,
		Result:    blogs,
	}
	return pageResult
}
