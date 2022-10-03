package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

// 标签数据库查询
var tagsearchFeild string = " id,name,createTime,updateTime  "

// 查询tag总数
func GetTagCount() int {
	row := utils.DBConn.QueryRow("select count(1) from tag where deleted=0")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var count int
	row.Scan(&count)
	return count

}

// 查询所有未删除的tag
func GetAllTag() ([]entity.Tag, error) {
	rows, err := utils.DBConn.Query("select " + tagsearchFeild + "from tag where deleted=0")
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var tags []entity.Tag
	for rows.Next() {
		var tag entity.Tag

		err := rows.Scan(&tag.Id, &tag.Name, &tag.CreateTime, &tag.UpdateTime)
		if err != nil {
			log.Println("tag取值出错", err)
			return nil, err
		}

		tags = append(tags, tag)
	}
	return tags, nil
}
