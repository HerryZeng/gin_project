package ch05

import "github.com/gin-gonic/gin"

func Router(rg *gin.RouterGroup) {
	rg.Use(MiddleWare2())
	rg.GET("/middleware2", MiddleWareDemo)
	rg.GET("/auth_test", gin.BasicAuth(gin.Accounts{
		"zs": "123456",
		"ls": "1234",
		"ww": "12345678",
	}), gin.WrapF(WarpTest), AuthTest)

	//rg.GET("/wrapF", WarpTest)
}
