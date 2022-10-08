package dao

import (
	"atao-go-blog/utils"
	"log"
)

// 标签数据库查询

// 查询tag总数
func GetTagCount() int {
	row := utils.DBConn.QueryRow("SELECT count(DISTINCT SUBSTRING_INDEX(SUBSTRING_INDEX(A.tags,',',help_topic_id+1),',',-1)) AS count " +
		"FROM" +
		" (SELECT tags ,del_flag FROM blog) a JOIN " +
		" mysql.help_topic b WHERE " +
		" b.help_topic_id < LENGTH(A.tags)-LENGTH(REPLACE(A.tags,',',''))+1 AND a.del_flag=0")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var count int
	row.Scan(&count)
	return count

}

// 查询所有未删除的tag
func GetAllTag() ([]string, error) {
	rows, err := utils.DBConn.Query(
		"SELECT DISTINCT SUBSTRING_INDEX(SUBSTRING_INDEX(A.tags,',',help_topic_id+1),',',-1) AS tag " +
			"FROM" +
			" (SELECT tags ,del_flag FROM blog) a JOIN " +
			" mysql.help_topic b WHERE " +
			" b.help_topic_id < LENGTH(A.tags)-LENGTH(REPLACE(A.tags,',',''))+1 AND a.del_flag=0")

	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var tags []string
	for rows.Next() {
		var str string

		err := rows.Scan(&str)
		if err != nil {
			log.Println("tag取值出错", err)
			return nil, err
		}

		tags = append(tags, str)
	}
	return tags, nil
}
