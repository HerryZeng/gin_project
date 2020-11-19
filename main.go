package main

import (
	"gin_project/controller/ch03"
	"gin_project/controller/ch04"
	"gin_project/controller/ch05"
	"gin_project/router"
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
	//r := gin.New()
	//r.Use(gin.Logger(), gin.Recovery())
	//r.Use(ch05.MiddleWare01, ch05.MiddleWare03())
	r.Use(ch05.MiddleWare01)

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
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/ch01")
	})
	// 路由分组
	router.Router(r)

	s := &http.Server{
		Addr:         ":9000",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
