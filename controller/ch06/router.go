package ch06

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup) {
	rg.GET("/gorm_test", GormTest)
}
