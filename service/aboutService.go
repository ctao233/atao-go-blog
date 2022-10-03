package service

import (
	"atao-go-blog/dao"
	"atao-go-blog/vo"
)

func AboutService() *vo.LayoutRes {
	lr := Layout()

	lr.Data = dao.Getdescription()
	return lr
}
