package logsource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

func LoadLogConfig() *LogConfig {
	log_conf := LogConfig{}
	file, err := os.Open("conf/logs_conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, ioerr := ioutil.ReadAll(file)
	if ioerr != nil {
		panic(ioerr)
	}
	jsonerr := json.Unmarshal(b, &log_conf)
	if jsonerr != nil {
		panic(jsonerr)
	}
	//fmt.Println(log_conf)
	return &log_conf
}
