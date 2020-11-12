package ch01

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context) {
	index := "Hello Gin"
	ctx.HTML(200, "index/index.html", index)
}

