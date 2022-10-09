package views

import (
	"atao-go-blog/common"
	"atao-go-blog/dao"
	"atao-go-blog/define"
	"atao-go-blog/service"
	"atao-go-blog/utils"
	"atao-go-blog/vo"
	"context"
	"net/http"
	"strconv"
	"strings"
)

// 首页映射处理函数
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	// 写入数据
	result := service.Layout()
	result.PagePhoto = dao.GetPagePhoto("indexPhoto")
	common.RenderHtml(w, "index", result)
}

// 时间轴
func (*HTMLApi) ArchivesIndex(w http.ResponseWriter, r *http.Request) {
	result := service.Layout()
	result.PagePhoto = dao.GetPagePhoto("archivesPhoto")
	common.RenderHtml(w, "archives", result)
}

// // 友情链接
// func (*HTMLApi) LinkIndex(w http.ResponseWriter, r *http.Request) {
// 	lr := service.LinkService()
// 	err := common.Templates["link.html"].Execute(w, lr)
// 	common.Check(err)

// }

// // 关于
// func (*HTMLApi) AboutIndex(w http.ResponseWriter, r *http.Request) {
// 	lr := service.AboutService()
// 	err := common.Templates["about.html"].Execute(w, lr)
// 	common.Check(err)
// }

func (*HTMLApi) NOtFound(w http.ResponseWriter, r *http.Request) {
	common.RenderHtml(w, "404", nil)
}

// 分类详情
func (*HTMLApi) DetailCategory(w http.ResponseWriter, r *http.Request) {
	result := service.Layout()
	result.PagePhoto = dao.GetPagePhoto("categoryPhoto")
	searchField := r.FormValue("category")

	categoryCount := dao.GetBlogByCategoryCount(searchField)
	result.Data = vo.TagOrCategory{searchField, categoryCount}

	common.RenderHtml(w, "detailsCategory", result)
}

// 标签详情
func (*HTMLApi) DetailTag(w http.ResponseWriter, r *http.Request) {
	result := service.Layout()
	result.PagePhoto = dao.GetPagePhoto("tagPhoto")
	searchField := r.FormValue("tag")

	tagCount := dao.GetBlogByTagsCount(searchField)
	result.Data = vo.TagOrCategory{searchField, tagCount}

	common.RenderHtml(w, "detailsTag", result)
}

// 博客观览页面
func (*HTMLApi) Blog(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("id")
	s := r.RemoteAddr

	if str != "" {

		id, err := strconv.Atoi(str)
		common.Check(err)
		lr := service.GetBlogByIdService(id)
		lr.PagePhoto = dao.GetPagePhoto("blogPhoto")
		common.RenderHtml(w, "blog", lr)

		if err := utils.RDB.PFAdd(context.Background(), define.BlogViewKey+str, strings.Split(s, ":")[0]).Err(); err != nil {
			panic(err)
		}
	}

}

// // 标签于分类详情页
// func (*HTMLApi) Details(w http.ResponseWriter, r *http.Request) {

// 	var detailsParam struct {
// 		Genre string
// 		param string
// 	}

// 	result := service.Layout()

// 	detailsParam.Genre = r.FormValue("searchType")
// 	detailsParam.param = r.FormValue("searchField")
// 	result.Data = detailsParam

// 	fmt.Print(result)
// 	common.RenderHtml(w, "details", result)

// }
