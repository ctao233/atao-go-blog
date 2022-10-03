package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

func GetSite() entity.Site {
	row := utils.DBConn.QueryRow("select * from site where id = 1")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var site entity.Site
	row.Scan(&site.Name, &site.Id, &site.Description, &site.Logo, &site.Favicon, &site.Notice, &site.BlogLink)
	return site
}
