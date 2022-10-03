package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

// 友情链接数据

func GetAllLink() ([]entity.Link, error) {
	rows, err := utils.DBConn.Query("select * from link")
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var links []entity.Link
	for rows.Next() {
		var link entity.Link

		err := rows.Scan(&link.Id, &link.Hphoto, &link.Url, &link.Name, &link.Introduce, &link.CreateTime)
		if err != nil {
			log.Println("tag取值出错", err)
			return nil, err
		}

		links = append(links, link)
	}
	return links, nil
}
