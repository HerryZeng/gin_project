package ch03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSyntaxTpl(ctx *gin.Context) {

	name := "dlzeng"
	arr := []int{1, 2, 3, 4, 5}
	mapData := map[string]interface{}{
		"name": name,
		"arr":  arr,
	}
	ctx.HTML(http.StatusOK, "ch03/test.html", mapData)
}
