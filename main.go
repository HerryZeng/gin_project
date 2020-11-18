package main

import (
	"gin_project/ch01"
	"gin_project/ch02"
	"gin_project/ch03"
	"gin_project/ch04"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"add":      ch03.AddNum,
		"subStr":   ch03.SubStr,
		"htmlSafe": ch03.HtmlSafe,
	})

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = v.RegisterValidation("len_valid", ch04.LenValid)
	}
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

	r.GET("/output_json", ch02.OutJson)
	r.GET("/output_ascii_json", ch02.OutAsciiJson)
	r.GET("/output_jsonp", ch02.OutJsonp)
	r.GET("/output_pure_json", ch02.OutPureJson)
	r.GET("/output_secure_json", ch02.OutSecureJson)
	r.GET("/output_xml", ch02.OutXml)
	r.GET("/output_yaml", ch02.OutYaml)
	r.GET("/output_proto", ch02.OutProto)

	r.GET("/redirect_a", ch02.RedirectA)
	r.GET("/redirect_b", ch02.RedirectB)

	// 模板语法
	r.GET("/test_tpl", ch03.TestSyntaxTpl)
	r.GET("/test_func_tpl", ch03.FuncTpl)

	// 数据绑定
	r.GET("/to_bind_form", ch04.ToBindForm)
	r.POST("/do_bind_form", ch04.DoBindForm)
	r.GET("/bind_query_string", ch04.BindQueryString)
	r.GET("/to_bind_json", ch04.ToBindJson)
	r.POST("/do_bind_json", ch04.DoBindJson)

	r.GET("/bind_uri/:name/:age/:address", ch04.BindUri)

	r.GET("/to_valid_data", ch04.ToValidData)
	r.POST("/do_valid_data", ch04.DoValidData)

	//_ = r.Run(":8080")
	s := &http.Server{
		Addr:         ":9000",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
