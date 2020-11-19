package ch02

import (
	"gin_project/controller/ch05"
	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	rg.GET("/user", User)
	rg.GET("/userinfo", UserInfoStruct)
	rg.GET("/array", ArrayController)
	rg.GET("/array_struct", ch05.MiddleWare03(), ArrayStruct)
	rg.GET("/map", MapController)
	rg.GET("/map_struct", MapStruct)
	rg.GET("/slice", SliceController)
	rg.GET("/param1/:Id", Param1)
	rg.GET("/param2/*Id", Param2)
	rg.GET("/query", GetQueryData)
	rg.GET("/query_array", GetQueryArrayData)
	rg.GET("/query_map", GetQueryMapData)
	rg.GET("/to_user_add", ToUserAdd)
	rg.POST("/do_user_add", DoUserAdd)

	rg.GET("/to_user_add3", ToUserAdd3)
	rg.POST("/do_user_add3", DoUserAdd3)

	rg.GET("/to_user_add4", ToUserAdd4)
	rg.POST("/do_user_add4", DoUserAdd4)

	rg.GET("/test_to_upload1", ToUpload1)
	rg.POST("/test_do_upload1", DoUpload1)

	rg.GET("/test_to_upload2", ToUpload2)
	rg.POST("/test_do_upload2", DoUpload2)

	rg.GET("/test_to_upload4", ToUploadFile4)
	rg.POST("/test_do_upload4", DoUploadFile4)

	rg.GET("/output_json", OutJson)
	rg.GET("/output_ascii_json", OutAsciiJson)
	rg.GET("/output_jsonp", OutJsonp)
	rg.GET("/output_pure_json", OutPureJson)
	rg.GET("/output_secure_json", OutSecureJson)
	rg.GET("/output_xml", OutXml)
	rg.GET("/output_yaml", OutYaml)
	rg.GET("/output_proto", OutProto)

	rg.GET("/redirect_a", RedirectA)
	rg.GET("/redirect_b", RedirectB)
}
