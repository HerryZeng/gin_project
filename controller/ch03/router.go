package ch03

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup)  {
	// 模板语法
	rg.GET("/test_tpl", TestSyntaxTpl)
	rg.GET("/test_func_tpl", FuncTpl)
}