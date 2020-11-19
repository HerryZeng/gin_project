package ch01

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup) {
	rg.GET("/", Index)
}
