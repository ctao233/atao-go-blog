package common

import (
	"atao-go-blog/entity"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
)

func GetSplit(str, symbol string) []string {
	s := strings.Split(str, symbol)
	return s
}

//获得格式化后得当前日期
func NowDate(layout string) string {
	return time.Now().Format(layout)
}

/// 日期格式化
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

// markdown 转html
func MdSweapHtml(target string) string {

	d := markdown.NormalizeNewlines([]byte(target))
	output := markdown.ToHTML(d, nil, nil)

	return string(output)
}

// 字数统计
func StrCount(target string) int {
	return len([]rune(target))
}

// 字数控制
func WordLimit(blogs []entity.Blog) {
	for i, v := range blogs {
		content := []rune(v.Content)
		if len(content) > 150 {
			content = content[0:150]
		}
		// 补充正则替换
		blogs[i].Content = strings.Replace(string(content)+".....", "#", "", -1)
	}
}
