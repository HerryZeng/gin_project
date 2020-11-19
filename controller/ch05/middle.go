package ch05

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare01(ctx *gin.Context) {
	startTime := time.Now()
	fmt.Println("自定义中间件1-开始")
	ctx.Next()
	timeCount := time.Since(startTime)
	fmt.Println(timeCount)
	fmt.Println("自定义中间件1-结束")
}

func MiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("这是自定义中间件2-开始")
		time.Sleep(3)
		context.Next()
		fmt.Println("这是自定义中间件2-结束")
	}
}

func MiddleWare03() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("这是自定义中间件3-开始")
		//if 3 < 4 {
		//	context.Abort()
		//}
		time.Sleep(3)
		context.Next()
		fmt.Println("这是自定义中间件3-结束")
	}
}

func MiddleWareDemo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是中间件返回的信息")
}
