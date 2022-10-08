package main

import (
	"atao-go-blog/common"
	"atao-go-blog/config"
	"atao-go-blog/router"
	"atao-go-blog/task"
	"log"
	"net/http"
)

func init() {

	// 初始化模板
	common.LoadTemplate()
	// 加载 定时任务
	task.StartTimer(task.SyncRdbtoMysql)
}
func main() {
	server := http.Server{
		Addr: config.Cfg.AppConfig.Address,
	}

	router.RouterInit()

	if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		log.Println(err)
	}
	log.Println("ip 端口", config.Cfg.AppConfig.Address)

}
