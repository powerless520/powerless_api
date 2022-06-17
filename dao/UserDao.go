package dao

import (
	"fmt"
	"gorm.io/gorm"
	"powerless_web/model"
)

func UserQuery(page, pageSize int, t *gorm.DB) {

	var userModels []model.SysUser
	err := t.Limit(pageSize).Offset(page).Find(&userModels).Error
	if err != nil {
		fmt.Println("UserQuery:" + err.Error())
		return
	}

	fmt.Println("userModels:", userModels)
}