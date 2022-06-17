package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"powerless_web/service"
)

func UserController(ctx *gin.Context,db *gorm.DB)  {
	fmt.Println("UserController Begin:")
	service.UserPage(1,1,db)
}