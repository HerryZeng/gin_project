package ch04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/bind_json.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "绑定失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "绑定成功",
		})
	}
}
