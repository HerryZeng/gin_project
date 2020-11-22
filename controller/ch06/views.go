package ch06

import (
	"gin_project/datasource"
	"gin_project/models"
	"github.com/gin-gonic/gin"
)

func GormTest(ctx *gin.Context) {
	user := models.User{
		Name: "dlzeng",
		Age:  19,
	}
	defer datasource.DB.Close()
	datasource.DB.Create(&user)
}
