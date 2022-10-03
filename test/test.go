package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// func main() {
// 	mkStr := "|数据类型|功能|\n|----|----|\n|blob|blob 最大长度为64kb|\n|mediumblob|blob 最大长度为16mb|\n|longblob|blob 最大长度为4gb|\n|tinyblob|blob 最大长度为255字节|"

// 	output := markdown.ToHTML([]byte(mkStr), nil, nil)
// 	s := string(output)
// 	fmt.Printf("s: %v\n", s)

// 	// fmt.Println("n=-1: ", strings.Replace(mkStr, "#", "", -1))
// 	// str := "/details/tag/go"
// 	// s := strings.Split(str, "/")
// 	// fmt.Print(s)
// 	// fmt.Print(s[1])

// }
var ctx = context.Background()

func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "120.79.203.17:6379",
		Password: "ataonihao.com", // no password set
		DB:       0,               // use default DB
	})

	s, err := rdb.Keys(ctx, "blog:view:*").Result()
	if err != nil {
		panic(err)
	}

	updatValue := ""
	upateId := "-1"
	for _, v := range s {
		id := strings.Split(v, ":")[2]
		ic := rdb.PFCount(context.Background(), v)
		i, _ := ic.Result()
		updatValue += " when " + id + " then " + strconv.Itoa((int(i)))
		upateId += "," + id
		fmt.Println(i)
	}

	fmt.Println(updatValue)
	fmt.Println(upateId)

	// for i := 0; i < 10; i++ {

	// 	ic := rdb.PFAdd(ctx, "myset", fmt.Sprint(i))
	// 	i2, err2 := ic.Result()
	// 	if err2 == nil {
	// 		fmt.Println(i2)
	// 	}

	// }

	// card, err := rdb.PFCount(ctx, "myset").Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("set cardinality", card)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
}

// func init() {
// 	startTimer(func() { fmt.Println("阿滔你好") })
// }

// func main() {
// 	//创建一个周期性的定时器

// 	for {
// 		time.Sleep(time.Second * 1)
// 	}
// }

func startTimer(f func()) {
	go func() {
		for {
			// 取得当前时间
			f()
			t := time.Now()
			next := t.Add(time.Hour * 24)
			fmt.Println(t)
			// 计算下一个0点
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			fmt.Println(t.Second())
			// 计算时间差
			duration := next.Sub(time.Now())
			fmt.Printf("duration  type: %T,\t val: %v\n", duration, duration)

			fmt.Print(next)
			t2 := time.NewTicker(duration)
			<-t2.C
		}
	}()
}
