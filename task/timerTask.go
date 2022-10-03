package task

import (
	"atao-go-blog/common"
	"atao-go-blog/dao"
	"atao-go-blog/define"
	"atao-go-blog/utils"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func SyncRdbtoMysql() {

	keys, err := utils.RDB.Keys(context.Background(), define.BlogViewKey+"*").Result()
	updatValue := ""
	upateId := "-1"
	for _, v := range keys {
		id := strings.Split(v, ":")[2]
		ic := utils.RDB.PFCount(context.Background(), v)
		i, _ := ic.Result()
		updatValue += " when " + id + " then " + strconv.Itoa((int(i)))
		upateId += "," + id
	}
	if err != nil {
		panic(err)
	}
	i, err := dao.UpdateBlogViewbyId(updatValue, upateId)

	if err != nil {
		panic(err)
	}
	if i > 0 {
		fmt.Print(common.DateDay(time.Now()) + " 以更新:" + strconv.Itoa(i) + "条数据")
	}

}

func StartTimer(f func()) {
	go func() {
		for {
			// 取得当前时间
			f()
			t := time.Now()
			next := t.Add(time.Hour * 24)
			// 计算下一个0点
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			// 计算时间差
			duration := next.Sub(time.Now())

			fmt.Print(next)
			t2 := time.NewTicker(duration)
			<-t2.C
		}
	}()
}
