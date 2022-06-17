package dao

import (
	"fmt"
	"powerless_api/model"
	"powerless_api/mysql"
)

func UserQuery(page, pageSize int) {
	// 获取mysql连接
	t := mysql.GetMysqlEngine()

	var userModels []model.SysUser
	err := t.Limit(pageSize).Offset(page).Find(&userModels).Error
	if err != nil {
		fmt.Println("UserQuery:" + err.Error())
		return
	}

	fmt.Println("userModels:", userModels)
}