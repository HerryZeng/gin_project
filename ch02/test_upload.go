package ch02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ToUpload1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/test_upload1.html", nil)
}

func DoUpload1(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Filename)
	}
	timeUnixStr := strconv.FormatInt(time.Now().Unix(), 10)
	err = ctx.SaveUploadedFile(file, "upload/"+timeUnixStr+file.Filename)
	if err != nil {
		ctx.String(http.StatusOK, "文件保存失败", err)
		return
	}
	ctx.String(http.StatusOK, "文件上传成功!")
}

func ToUpload2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/test_upload2.html", nil)
}

func DoUpload2(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	files := form.File["file"]
	for _, file := range files {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(file.Filename)
		}
		timeUnixStr := strconv.FormatInt(time.Now().Unix(), 10)
		err = ctx.SaveUploadedFile(file, "upload/"+timeUnixStr+file.Filename)
		if err != nil {
			ctx.String(http.StatusOK, "文件保存失败", err)
			return
		}
	}
	ctx.String(http.StatusOK, "多文件上传成功!")
}
