package ch05

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var mapData map[string]interface{} = map[string]interface{}{
	"zs": gin.H{"age": 22, "address": "xxx"},
	"ls": gin.H{"age": 25, "address": "yyy"},
	"ww": gin.H{"age": 36, "address": "zzz"},
}

func AuthTest(ctx *gin.Context) {
	userName := ctx.Query("username")
	data, ok := mapData[userName]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": userName, "data": data})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{"user": userName, "data": data})
	}
}

func WarpTest(w http.ResponseWriter,r *http.Request)  {

}