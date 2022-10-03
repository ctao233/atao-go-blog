package utils

import (
	"atao-go-blog/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

// 初始化连接
func init() {
	// 获取数据库配置
	mysqlcfg := &config.Cfg.MysqlConfig

	dataSorceName := fmt.Sprintf("%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlcfg.UserPassword, mysqlcfg.IpPort, mysqlcfg.Database)
	// dataSorceName := fmt.Sprintf("root:root@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))

	// 配置数据库类型 ，连接地址
	db, err := sql.Open("mysql", dataSorceName)

	if err != nil {
		log.Println("数据库连接异常")
		panic(err)
	}

	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)

	err = db.Ping()

	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}

	DBConn = db

}
