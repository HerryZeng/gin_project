package middleware

import "github.com/gin-gonic/gin"

func CROSMiddleWare(ctx *gin.Context)  {
	ctx.Header("Access-Control-Allow-Origin","*")
	ctx.Next()
}
