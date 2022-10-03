package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

// 查询用户所有内容
func GetInformation() (entity.Information, error) {

	row := utils.DBConn.QueryRow("select *from information")
	if row.Err() != nil {
		log.Println(row.Err())
		return entity.Information{}, row.Err()
	}

	var infor entity.Information
	row.Scan(&infor.Id, &infor.Photo, &infor.Sentence, &infor.Email, &infor.Github, &infor.Name)
	return infor, nil
}
