package ch02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Id      int    `form:"id"`
	Name    string `form:"name"`
	Age     int    `form:"age"`
	Address string `form:"address"`
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

func MapController(ctx *gin.Context) {
	mapData1 := map[string]string{"Name": "Harry", "Age": "18"}
	mapData2 := map[string]int{"Age": 18}
	mapData3 := map[string]interface{}{"mapdata1": mapData1, "mapdata2": mapData2}
	ctx.HTML(http.StatusOK, "ch02/map.html", mapData3)
}

func MapStruct(ctx *gin.Context) {
	mapStruct := map[string]userInfo{
		"user1": {Id: 1, Name: "Harry", Age: 34, Address: "中国深圳市宝安区西乡固戍东山得宝楼2栋一单元1103"},
		"user2": {Id: 2, Name: "Hank", Age: 40, Address: "中国深圳市宝安区西乡固戍东山得宝楼1栋一单元1103"},
		"user3": {Id: 3, Name: "Meris", Age: 29, Address: "中国深圳市宝安区西乡固戍东山得宝楼2栋一单元1106"},
	}
	ctx.HTML(http.StatusOK, "ch02/map_struct.html", mapStruct)
}

func SliceController(ctx *gin.Context) {
	sliceData := []int{1, 2, 3, 4, 5, 6}
	ctx.HTML(http.StatusOK, "ch02/slice.html", sliceData)
}

func Param1(ctx *gin.Context) {
	Id := ctx.Param("Id")
	ctx.String(http.StatusOK, "Hello,%s", Id)
}

func Param2(ctx *gin.Context) {
	Id := ctx.Param("Id")
	ctx.String(http.StatusOK, "Hello,%s", Id)
}

func GetQueryData(ctx *gin.Context) {
	Id := ctx.Query("id")
	name := ctx.DefaultQuery("name", "DLZENG")
	fmt.Println(name)
	ctx.String(http.StatusOK, "DLZENG,%s", Id)
}

func GetQueryArrayData(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	ctx.String(http.StatusOK, "DLZENG,%s", ids)
}

func GetQueryMapData(ctx *gin.Context) {
	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK, "DLZENG,%s", user)
}

func ToUserAdd(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "ch02/user_add.html", nil)
}

func DoUserAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age", "18")
	love := ctx.PostFormArray("love")
	user := ctx.PostFormMap("user")
	fmt.Println(username, password, age, user)
	fmt.Println(love)
	ctx.String(http.StatusOK, "用户名是:%s,密码是:%s,年龄是:%s,爱好是:%v", username, password, age, love)
}

func ToUserAdd3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add3.html", nil)
}

func DoUserAdd3(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age", "18")

	fmt.Println(username, password, age)
	mapData := map[string]interface{}{
		"code": 200,
		"msg":  "成功",
	}
	ctx.JSON(http.StatusOK, mapData)
}

func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add4.html", nil)
}

func DoUserAdd4(ctx *gin.Context) {
	var userInfo userInfo
	err := ctx.ShouldBind(&userInfo)
	fmt.Println(err)
	fmt.Println(userInfo)
	ctx.String(http.StatusOK, "绑定成功")
}
