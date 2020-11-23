package ch11

import (
	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	//rg.Use(middleware.CROSMiddleWare)
	rg.GET("/api_axios", ApiAxios)
}
