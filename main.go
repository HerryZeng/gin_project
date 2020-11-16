package main

import (
	"gin_project/ch01"
	"gin_project/ch02"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "static")

	r.GET("/", ch01.Index)
	r.GET("/user", ch02.User)
	r.GET("/userinfo", ch02.UserInfoStruct)
	r.GET("/array", ch02.ArrayController)
	r.GET("/array_struct", ch02.ArrayStruct)
	r.GET("/map", ch02.MapController)
	r.GET("/map_struct", ch02.MapStruct)
	r.GET("/slice", ch02.SliceController)
	r.GET("/param1/:Id", ch02.Param1)
	r.GET("/param2/*Id", ch02.Param2)
	r.GET("/query", ch02.GetQueryData)
	r.GET("/query_array", ch02.GetQueryArrayData)
	r.GET("/query_map", ch02.GetQueryMapData)
	r.GET("/to_user_add", ch02.ToUserAdd)
	r.POST("/do_user_add", ch02.DoUserAdd)

	r.GET("/to_user_add3", ch02.ToUserAdd3)
	r.POST("/do_user_add3", ch02.DoUserAdd3)

	r.GET("/to_user_add4", ch02.ToUserAdd4)
	r.POST("/do_user_add4", ch02.DoUserAdd4)

	r.GET("/test_to_upload1", ch02.ToUpload1)
	r.POST("/test_do_upload1", ch02.DoUpload1)

	r.GET("/test_to_upload2", ch02.ToUpload2)
	r.POST("/test_do_upload2", ch02.DoUpload2)

	r.GET("/test_to_upload4", ch02.ToUploadFile4)
	r.POST("/test_do_upload4", ch02.DoUploadFile4)

	_ = r.Run(":8080")
}
