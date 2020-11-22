package datasource

import (
	"fmt"
	"gin_project/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var Err error

func init() {

	data_source := LoadMySQLConf()
	//fmt.Println(data_source)
	host := data_source.Host
	port := data_source.Port
	user_name := data_source.UserName
	password := data_source.Password
	database := data_source.Database
	log_mode := data_source.LogMode

	//connStr := "root:Kaka@2019@tcp(127.0.0.1:3306)/gin_project?charset=utf8&parseTime=True&loc=Local"
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user_name, password, host, port, database)
	DB, Err = gorm.Open("mysql", connStr)
	if Err != nil {
		panic(Err)
	}

	//defer DB.Close()

	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(50)
	DB.LogMode(log_mode)
	DB.AutoMigrate(&models.User{})
}
