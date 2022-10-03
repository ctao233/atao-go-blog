package service

import (
	"atao-go-blog/common"
	"atao-go-blog/dao"
	"atao-go-blog/define"
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"atao-go-blog/vo"
	"context"
	"fmt"
	"strconv"
	"time"
)

// 分页
func GetblogPageService(page, pageSize int) *vo.PageResult {

	blogs, err := dao.GetAllBlog(((page - 1) * pageSize), pageSize)

	if err != nil {
		return nil
	}
	// 遍历文章 展示部分内容 与日期格式化
	common.WordLimit(blogs)

	// 编写页码逻辑
	total := dao.GetBlogCount()
	//总页码数
	pagesCount := (total-1)/pageSize + 1

	pageResult := &vo.PageResult{
		Page:      page,
		PageSize:  pageSize,
		PageCount: pagesCount,
		Total:     total,
		Result:    blogs,
	}

	return pageResult
}

// 根据id 查询文章业务
func GetBlogByIdService(id int) *vo.LayoutRes {
	b, err := dao.GetBlogById(id)
	if err != nil {
		return nil
	}
	i, err2 := utils.RDB.PFCount(context.Background(), define.BlogViewKey+strconv.Itoa(id)).Result()
	if err2 != nil {
		panic(err2)
	}
	b.View = i
	// 拿到通用模板数据
	lr := Layout()
	lr.Data = b
	return lr
}

// 归档查询 日期降序
func GetArchIvesBlogByYM(page, pageSize int) *vo.PageResult {

	blogs, _ := dao.GetArchIvesBlog(((page - 1) * pageSize), pageSize)
	ym, _ := dao.GetArchIveYearMonth()

	testRsult := make([]map[string][]entity.Blog, 0)
	for j := 0; j < len(ym); j++ {
		mresult := make(map[string][]entity.Blog)
		var mblog []entity.Blog
		for _, v := range blogs {

			// 年月归档
			if v.YearMonth == ym[j] {
				content := []rune(v.Content)
				if len(content) > 100 {
					content = content[0:100]
				}
				v.Content = string(content)
				mblog = append(mblog, v)
			}
		}
		// 有数据不为空直接添加
		if mblog != nil {
			mresult[ym[j]] = mblog
			testRsult = append(testRsult, mresult)
		}
	}

	//按年归档
	// result := make(map[string][]map[string][]entity.Blog)

	// for j := 0; j < len(ym); j += 2 {
	// 	year := strings.Split(ym[j], "-")[0]

	// 	var mblogs []map[string][]entity.Blog
	// 	for k, v := range mresult {
	// 		if strings.Split(v[0].YearMonth, "-")[0] == year {
	// 			m := map[string][]entity.Blog{k: v}
	// 			mblogs = append(mblogs, m)
	// 		}

	// 	}
	// 	if mblogs != nil {
	// 		fmt.Printf("year: %v\n", year)
	// 		result[year] = mblogs
	// 	}
	// }

	// 倒叙
	// 编写页码逻辑
	total := dao.GetBlogCount()
	//总页码数
	pagesCount := (total-1)/pageSize + 1

	pageResult := &vo.PageResult{
		Page:      page,
		PageSize:  pageSize,
		PageCount: pagesCount,
		Total:     total,
		Result:    testRsult,
	}

	return pageResult

}

// 文章搜索
func SearchBlogBytext(text string) []vo.SearchBlogVo {
	sbv, err := dao.GetBlogBytext(text)

	if err != nil {
		fmt.Println(common.DateDay(time.Now()) + " 检索异常:" + err.Error())
		panic(err)
	}

	return sbv
}
