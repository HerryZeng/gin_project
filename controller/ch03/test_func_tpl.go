package ch03

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func FuncTpl(ctx *gin.Context) {
	name := "dlzeng"
	age := 19
	sliceData := []string{"张三", "李四", "王五"}

	//now_time := time.Now().Format("2006/01/02 15:04:05")
	now_time := time.Now()
	//fmt.Println(now_time)
	mapData := map[string]interface{}{
		"name":      name,
		"age":       age,
		"slicedata": sliceData,
		"now_time":  now_time,
	}
	ctx.HTML(http.StatusOK, "ch03/test_func_tpl.html", mapData)
}

func AddNum(n1, n2 int) int {
	return n1 + n2
}

func SubStr(str string, num int) string {
	if len(str) <= num {
		return str
	}
	return str[0:num] + "..."
}

func HtmlSafe(str string) template.HTML {
	return template.HTML(str)
}
