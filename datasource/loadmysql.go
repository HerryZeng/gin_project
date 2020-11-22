package datasource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MySQLConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Database string `json:"database"`
	LogMode  bool   `json:"log_mode"`
}

func LoadMySQLConf() MySQLConf {

	mysqlConf := MySQLConf{}
	file, err := os.Open("conf/mysql_conf.json")
	if err != nil {
		panic(err)
	}
	b, ioerr := ioutil.ReadAll(file)
	if ioerr != nil {
		panic(ioerr)
	}
	jsonerr := json.Unmarshal(b, &mysqlConf)
	if jsonerr != nil {
		panic(jsonerr)
	}
	return mysqlConf
}
