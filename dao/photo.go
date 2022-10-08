package dao

import (
	"atao-go-blog/utils"
	"log"
)

func GetPagePhoto(pageName string) string {
	row := utils.DBConn.QueryRow("SELECT " + pageName + " FROM photo")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var photoLink string
	row.Scan(&photoLink)
	return photoLink
}
