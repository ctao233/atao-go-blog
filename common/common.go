package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"sync"
	"text/template"
)

// 模板存放文件目录
const (
	ListDir      = 0x0001
	TEMPLATE_DIR = "./template/blog"
)

//模板缓存
var Templates = make(map[string]*template.Template)

func LoadTemplate() {

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	Check(err)

	// 以文件名为key 存储模板
	var templateName, templatePath string

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		for _, fileInfo := range fileInfoArr {
			templateName = fileInfo.Name()

			// 返回路径的扩展名判断是否为html 文件
			if ext := path.Ext(templateName); ext != ".html" {
				continue
			}

			// 当前路径下公共的模板路径
			path, _ := os.Getwd()
			right := path + "/template/blog/layout/right.html"
			header := path + "/template/blog/layout/header.html"
			foot := path + "/template/blog/layout/foot.html"
			templatePath = TEMPLATE_DIR + "/" + templateName

			log.Println("Loading template:", templatePath)
			// 加载模板要用到得方法
			t := template.New(templateName)

			// template.Must()确保了模板不能解析成功时，一定会触发错误处理流程。
			t.Funcs(template.FuncMap{"DateDay": DateDay, "NowDate": NowDate, "GetSplit": GetSplit, "MdSweapHtml": MdSweapHtml, "StrCount": StrCount})
			initT := template.Must(t.ParseFiles(templatePath, header, foot, right))

			Templates[templateName] = initT
		}
		w.Done()
	}()
	w.Wait()

}

// 模板渲染方法
func RenderHtml(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl += ".html"
	// 通过名字取路径
	err := Templates[tmpl].Execute(w, data)
	Check(err)
}

// 用于统一捕获50x系列的服务端内部错误
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	把业务逻辑处理函数作为参数传入到safeHandler()方法中，
	这样任何一个错误处理流程向上回溯的时候，我们都能对其进行拦截处理，从而也能避免程序停止运行

	巧妙地使用了defer关键字搭配recover()方法终结panic的肆行。
	safeHandler()接收一个业务逻辑处理函数作为参数，同时调用这个业务逻辑处理函数。
	该业务逻辑函数执行完毕后，safeHandler()中defer指定的匿名函数会执行。
	倘若业务逻辑处理函数里边引发了panic，则调用recover()对其进行检测，若为一般性的错误，
	则输出HTTP 50x出错信息并记录日志，而程序将继续良好运行。
*/
func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				// 或者自定义50x错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error",e)
				// logging

				log.Println("WARN: panic fired in %v.panic - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

// 动静分离
func StaticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		fmt.Printf("file: %v\n", file)
		if (flags & ListDir) == 0 {
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

// request json 数据传递
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
