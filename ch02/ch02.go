package ch02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Id      int
	Name    string
	Age     int
	Address string
}

func User(ctx *gin.Context) {
	user := "DLZENG"
	ctx.HTML(200, "user/user.html", user)
}

func UserInfoStruct(ctx *gin.Context) {
	userInfo := userInfo{Id: 1, Name: "Herry", Age: 34, Address: "中国深圳市宝安区西乡固戍东山得宝楼2栋一单元1103"}
	ctx.HTML(http.StatusOK, "ch02/userinfo.html", userInfo)
}

func ArrayController(ctx *gin.Context) {
	arr1 := [...]int{1, 3, 4, 9}
	ctx.HTML(http.StatusOK, "ch02/array.html", arr1)
}

func ArrayStruct(ctx *gin.Context) {
	arrayStruct := [3]userInfo{
		{Id: 1, Name: "Harry", Age: 34, Address: "中国深圳市宝安区西乡固戍东山得宝楼2栋一单元1103"},
		{Id: 2, Name: "Hank", Age: 40, Address: "中国深圳市宝安区西乡固戍东山得宝楼1栋一单元1103"},
		{Id: 3, Name: "Meris", Age: 29, Address: "中国深圳市宝安区西乡固戍东山得宝楼2栋一单元1106"},
	}

	ctx.HTML(http.StatusOK, "ch02/array_struct.html", arrayStruct)
}
