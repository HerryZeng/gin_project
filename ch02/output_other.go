package ch02

import (
	"fmt"
	user "gin_project/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OutJson(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutAsciiJson(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutJsonp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutPureJson(ctx *gin.Context) {
	ctx.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutSecureJson(ctx *gin.Context) {
	ctx.SecureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutXml(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutYaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello,World!</b>",
	})
}

func OutProto(ctx *gin.Context) {
	userData := &user.User{Name: "dlzeng", Age: 22}
	ctx.ProtoBuf(http.StatusOK, userData)
}

func RedirectA(ctx *gin.Context) {
	fmt.Println("这是A路由")

	ctx.Redirect(http.StatusFound, "/redirect_b")
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("这是B路由")
	ctx.String(http.StatusOK, "这是B路由")
}
