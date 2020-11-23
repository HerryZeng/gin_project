package ch09

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionTest(ctx *gin.Context) {
	// 初始化session对象
	session := sessions.Default(ctx)
	//设置session
	session.Set("name", "dlzeng")
	//获取session
	name := session.Get("name")
	fmt.Println("Session开始打印")
	fmt.Println(name)
}
