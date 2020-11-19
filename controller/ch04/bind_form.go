package ch04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name    string `form:"name" json:"name" uri:"name"`
	Age     int    `form:"age" json:"age" uri:"age"`
	Address string `form:"address" json:"address" uri:"address"`
}

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		ctx.String(http.StatusOK, "绑定失败")
	} else {
		ctx.String(http.StatusOK, "绑定成功")
	}
}
