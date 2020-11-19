package ch03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id    int
	Title string
	Desc  string
}

func TestSyntaxTpl(ctx *gin.Context) {

	name := "dlzeng"
	arr := []int{1, 2, 3, 4, 5}
	article := Article{Id: 1, Title: "西游记", Desc: "西游记外传,不容错过!"}
	mapData := map[string]interface{}{
		"name":    name,
		"arr":     arr,
		"article": article,
	}
	ctx.HTML(http.StatusOK, "ch03/test.html", mapData)
}
