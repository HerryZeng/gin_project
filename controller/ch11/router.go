package ch11

import (
	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	//rg.Use(middleware.CROSMiddleWare)
	rg.GET("/api_axios", ApiAxios)
	rg.GET("/get_books", GetBooks)
	rg.GET("/book_detail", GetBookDetail)
	rg.POST("/post_test", GetPostData)
}
