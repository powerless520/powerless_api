package main

import (
	"github.com/gin-gonic/gin"
	"powerless_api/controller"
	"powerless_api/mysql"
)

func main() {
	_ = mysql.Init()

	r := gin.Default()

	r.GET("/user",controller.UserController)

	r.Run(":9090")
}
