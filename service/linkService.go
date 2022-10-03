package service

import (
	"atao-go-blog/dao"
	"atao-go-blog/vo"
	"log"
)

func LinkService() *vo.LayoutRes {
	lr := Layout()
	links, err := dao.GetAllLink()
	if err != nil {
		log.Println("LinkService:  运行出错", err)
		return nil
	}
	lr.Data = links
	return lr
}
