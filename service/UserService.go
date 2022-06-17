package service

import (
	"fmt"
	"powerless_api/dao"
)

func UserPage(page int,pageSize int)  {
	fmt.Println("page : ", page, "pageSize : ", pageSize)
	dao.UserQuery(page,pageSize)
}