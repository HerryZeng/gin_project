package router

import (
	"gin_project/controller/ch01"
	"gin_project/controller/ch02"
	"gin_project/controller/ch03"
	"gin_project/controller/ch04"
	"gin_project/controller/ch05"
	"gin_project/controller/ch06"
	"gin_project/controller/ch07"
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

	ch01.Router(ch01Rg)
	ch02.Router(ch02Rg)
	ch03.Router(ch03Rg)
	ch04.Router(ch04Rg)
	ch05.Router(ch05Rg)
	ch06.Router(ch06Rg)
	ch07.Router(ch07Rg)
}
