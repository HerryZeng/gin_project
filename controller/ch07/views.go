package ch07

import (
	"gin_project/logsource"
	"github.com/gin-gonic/gin"
)

func LogTest(ctx *gin.Context) {
	logsource.Log.Info("这是一条Info信息日志")
	logsource.Log.Debug("这是一条Debug信息日志")
	logsource.Log.Warn("这是一条Warn信息日志")
}
