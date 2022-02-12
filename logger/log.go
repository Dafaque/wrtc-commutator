package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	if _, exist := os.LookupEnv("APP.ENABLE_LOG"); exist {
		logger = log.Default()
		println("LOGGER ENABLED")
	}
}

func Println(v ...interface{}) {
	if logger != nil {
		logger.Println(v...)
	}
}

func Printf(format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}
