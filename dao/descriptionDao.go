package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

// 网站关于表

func Getdescription() entity.Description {
	row := utils.DBConn.QueryRow("select * from description where id = 1")
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var description entity.Description
	row.Scan(&description.Id, &description.Name, &description.Constellation,
		&description.Experience, &description.Sentence, &description.Idol,
		&description.Character, &description.Contact, &description.Information, &description.Language)
	return description
}
