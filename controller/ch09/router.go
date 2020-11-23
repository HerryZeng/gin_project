package ch09

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup) {
	rg.GET("/session_test", SessionTest)
}
