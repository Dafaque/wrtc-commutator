package servertools

import (
	"os"
	"strconv"
)

var (
	MAX_CONNECTIONS int = 2
)

func init() {
	if v, exist := os.LookupEnv("APP.MAX_CONNECTIONS"); exist {
		i, err := strconv.Atoi(v)
		if err != nil {
			println("err init consts:", err.Error())
			return
		}
		MAX_CONNECTIONS = i
	}
	println("MAX_CONNECTIONS:", MAX_CONNECTIONS)
}
