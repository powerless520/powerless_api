package main

import (
	"github.com/gin-gonic/gin"
	"powerless_web/controller"
	"powerless_web/mysql"
)

func main() {

	db := mysql.Init()
	ctx := &gin.Context{}
	controller.UserController(ctx, db)

}
