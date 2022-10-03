package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"log"
)

func GetFootInfo() (entity.Footer, error) {
	row := utils.DBConn.QueryRow("SELECT * FROM Footer")
	if row.Err() != nil {
		log.Println(row.Err())
		return entity.Footer{}, row.Err()
	}

	var footer entity.Footer
	row.Scan(&footer.Id, &footer.About, &footer.Number, &footer.Copyright, &footer.Powerby, &footer.Byurl)
	return footer, nil

}
