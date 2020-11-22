package logsource

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	log_conf := LoadLogConfig()

	file, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	Log.Out = file
	level_mapping := map[string]logrus.Level{
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(level_mapping[log_conf.LogLevel])

	Log.SetFormatter(&logrus.TextFormatter{})
}
