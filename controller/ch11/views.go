package ch11

import (
	"fmt"
	"gin_project/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func ApiAxios(ctx *gin.Context) {
	// 结构体|map
	user := models.User{Id: 2, Name: "dlzeng", Age: 22}
	// 数组|切片
	array1 := []int{1, 2, 3, 4, 5, 6}
	// 结构体数组
	arrayStruct := []models.User{
		{Id: 2, Name: "Herry", Age: 32},
		{Id: 3, Name: "ddd", Age: 23},
	}
	// 结构体map
	mapStruct := map[string]models.User{
		"user1": {Id: 1, Name: "abc", Age: 21},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"msg":         "成功",
		"user":        user,
		"array1":      array1,
		"arrayStruct": arrayStruct,
		"mapStruct":   mapStruct,
	})
}

var books = []Book{
	{Id: 1, Name: "西游记", Author: "吴承恩"},
	{Id: 2, Name: "水浒传", Author: "施耐庵"},
	{Id: 3, Name: "红楼梦", Author: "曹雪芹"},
	{Id: 4, Name: "三国演义", Author: "罗贯中"},
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func GetBookDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	fmt.Println("GetBookDetail")
	fmt.Println(id)
	id_new, _ := strconv.Atoi(id)

	book := books[id_new-1]
	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetPostData(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	fmt.Println(name, password)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
	})
}

func GetFile(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	fmt.Println(name, password)
	file, _ := ctx.FormFile("upload_file")
	filePath := "upload/" + file.Filename
	_ = ctx.SaveUploadedFile(file, filePath)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
	})
}

func GetFiles(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	fmt.Println(name, password)
	form,_ := ctx.MultipartForm()
	files := form.File["upload_files"]
	for _,file := range files {
		filePath := "upload/" + file.Filename
		_ = ctx.SaveUploadedFile(file, filePath)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
	})
}
