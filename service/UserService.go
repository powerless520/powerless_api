package service

import (
	"fmt"
	"gorm.io/gorm"
	"powerless_web/dao"
)

func UserPage(page int,pageSize int,db *gorm.DB)  {
	fmt.Println("page : ", page, "pageSize : ", pageSize)
	dao.UserQuery(page,pageSize,db)
}