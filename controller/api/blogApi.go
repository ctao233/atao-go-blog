package api

import (
	"atao-go-blog/service"
	"atao-go-blog/vo"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (*Api) GetBlogPage(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	pr := service.GetblogPageService(page, pageSize)

	if pr != nil {
		pr.Status = 200
		b, err := json.Marshal(pr)
		if err != nil {
			log.Println("json 解析出错")
		}
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.Write(b)
	}
}

func (*Api) GetarchivesBlogPage(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	pr := service.GetArchIvesBlogByYM(page, pageSize)

	if pr != nil {
		pr.Status = 200
		b, err := json.Marshal(pr)
		if err != nil {
			log.Println("json 解析出错")
		}
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.Write(b)
	}
}

func (*Api) GetBlogByTagsApi(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	condition := r.FormValue("tag")

	pr := service.GetblogPageByTagsService(page, pageSize, condition)
	if pr != nil {
		pr.Status = 200
		b, err := json.Marshal(pr)
		if err != nil {
			log.Println("json 解析出错")
		}
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.Write(b)
	}
}

func (*Api) GetBlogByCategoryApi(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	condition := r.FormValue("category")

	pr := service.GetblogPageByCategoryService(page, pageSize, condition)
	if pr != nil {
		pr.Status = 200
		b, err := json.Marshal(pr)
		if err != nil {
			log.Println("json 解析出错")
		}
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.Write(b)
	}
}

// 文章搜索
func (*Api) SearchBlogBytext(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("searchText")

	sbv := service.SearchBlogBytext(s)

	if sbv != nil {
		var result vo.Result
		result.Code = 200
		result.Data = sbv
		b, err := json.Marshal(result)
		if err != nil {
			log.Println("json 解析出错")
		}
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.Write(b)
	}

}
