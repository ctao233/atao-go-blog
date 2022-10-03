package dao

import (
	"atao-go-blog/utils"
	"log"
)

// 查询分类总数量
func GetCategoryCount() int {
	row := utils.DBConn.QueryRow("select count(1) from category  where del_flag=0 ")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var count int
	row.Scan(&count)
	return count

}

// 查询分类对应的博客数量
func GetCategoryBlogCount() (map[string]int, error) {
	rows, err := utils.DBConn.Query("select catalog, count(blog.catalog) as count from blog where del_flag=0  GROUP BY blog.catalog")
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}
	var cgcount = make(map[string]int)
	for rows.Next() {
		var catalog string
		var count int
		err := rows.Scan(&catalog, &count)
		if err != nil {
			log.Println("GetCategoryBlogCount取值出错", err)
			return nil, err
		}
		cgcount[catalog] = count

	}
	return cgcount, nil
}
