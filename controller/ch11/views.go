package ch11

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiAxios(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
	})
}
