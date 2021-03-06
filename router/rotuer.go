package router

import (
	"gin_project/controller/ch01"
	"gin_project/controller/ch02"
	"gin_project/controller/ch03"
	"gin_project/controller/ch04"
	"gin_project/controller/ch05"
	"gin_project/controller/ch06"
	"gin_project/controller/ch07"
	"gin_project/controller/ch09"
	"gin_project/controller/ch11"
	"gin_project/middleware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	ch01Rg := r.Group("/ch01")
	ch02Rg := r.Group("/ch02")
	ch03Rg := r.Group("/ch03")
	ch04Rg := r.Group("/ch04")
	ch05Rg := r.Group("/ch05")
	ch06Rg := r.Group("/ch06")
	ch07Rg := r.Group("/ch07")
	ch09Rg := r.Group("/ch09")
	ch11Rg := r.Group("/ch11")
	ch11Rg.Use(middleware.CROSMiddleWare)

	ch01.Router(ch01Rg)
	ch02.Router(ch02Rg)
	ch03.Router(ch03Rg)
	ch04.Router(ch04Rg)
	ch05.Router(ch05Rg)
	ch06.Router(ch06Rg)
	ch07.Router(ch07Rg)
	ch09.Router(ch09Rg)
	ch11.Router(ch11Rg)
}
