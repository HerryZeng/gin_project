package main

import (
	"gin_project/ch01"
	"gin_project/ch02"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "static")

	r.GET("/", ch01.Index)
	r.GET("/user", ch02.User)
	r.GET("/userinfo", ch02.UserInfoStruct)
	r.GET("/array", ch02.ArrayController)
	r.GET("/array_struct", ch02.ArrayStruct)
	r.Run(":8080")
}
