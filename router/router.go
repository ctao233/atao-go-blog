package router

import (
	"net/http"

	"atao-go-blog/controller/api"
	"atao-go-blog/controller/views"
)

func RouterInit() {

	/*
		页面
	*/
	http.HandleFunc("/AtaoBlog", views.HTML.Index)
	http.HandleFunc("/archives", views.HTML.ArchivesIndex)
	http.HandleFunc("/blog", views.HTML.Blog)
	//http.HandleFunc("/tag", views.HTML.TagIndex)
	// http.HandleFunc("/link", views.HTML.LinkIndex)
	http.HandleFunc("/detailCategory", views.HTML.DetailCategory)
	http.HandleFunc("/detailTag", views.HTML.DetailTag)
	http.HandleFunc("/404", views.HTML.NOtFound)
	// http.HandleFunc("/category", views.HTML.CategoryIndex)
	// http.HandleFunc("/about", views.HTML.AboutIndex)

	// // 数据
	http.HandleFunc("/api/blogpage", api.JsonApi.GetBlogPage)
	http.HandleFunc("/api/archivesBlog", api.JsonApi.GetarchivesBlogPage)
	http.HandleFunc("/api/detailTag", api.JsonApi.GetBlogByTagsApi)
	http.HandleFunc("/api/detailCategory", api.JsonApi.GetBlogByCategoryApi)
	http.HandleFunc("/api/search", api.JsonApi.SearchBlogBytext)
	// 静态资源
	http.Handle("/blog/", http.StripPrefix("/blog/", http.FileServer(http.Dir("public/blog/"))))
}
