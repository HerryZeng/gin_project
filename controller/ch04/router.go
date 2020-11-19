package ch04

import (
	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	// 数据绑定
	rg.GET("/to_bind_form", ToBindForm)
	rg.POST("/do_bind_form", DoBindForm)
	rg.GET("/bind_query_string", BindQueryString)
	rg.GET("/to_bind_json", ToBindJson)
	rg.POST("/do_bind_json", DoBindJson)

	rg.GET("/bind_uri/:name/:age/:address", BindUri)

	rg.GET("/to_valid_data", ToValidData)
	rg.POST("/do_valid_data", DoValidData)
}
