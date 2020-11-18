package ch04

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//type Article struct {
//	Id      int    `form:"-"`
//	Title   string `form:"title" binding:"required"`
//	Content string `form:"content" binding:"min=5"`
//	Desc    string `form:"desc" binding:"len_valid"`
//}

type Article struct {
	Id      int    `form:"-"`
	Title   string `form:"title" valid:"Required"`
	Content string `form:"content" valid:"Required"`
	Desc    string `form:"desc" valid:"Required"`
}

func ToValidData(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "ch04/valid_data.html", nil)
}

//func DoValidData(ctx *gin.Context) {
//	var article Article
//	err := ctx.ShouldBind(&article)
//	fmt.Println(article)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		ctx.String(http.StatusOK, "验证成功")
//	}
//}

func DoValidData(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBind(&article)
	fmt.Println(article)
	valid := validation.Validation{}
	b, _ := valid.Valid(&article)
	if !b {
		for _, err2 := range valid.Errors {
			fmt.Println(err2.Key)
			fmt.Println(err2.Message)
		}
	}
	if err != nil {
		fmt.Println(err)
	} else {
		ctx.String(http.StatusOK, "验证成功")
	}
}

var LenValid validator.Func = func(fl validator.FieldLevel) bool {
	//data := fl.Field().Interface().(string)
	data := fl.Field().Interface().(string)
	fmt.Println(data)
	if len(data) < 8 {
		return false
	} else {
		return true
	}
}
