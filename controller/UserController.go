package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"powerless_api/service"
)

func UserController(ctx *gin.Context)  {
	fmt.Println("UserController Begin:")

	service.UserPage(1,1)

	ctx.Status(200)
}