package ch07

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup) {
	rg.GET("/log_test", LogTest)
}
