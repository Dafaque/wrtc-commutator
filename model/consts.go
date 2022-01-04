package model

import "os"

var (
	SECRET []byte = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func init() {
	if v, exist := os.LookupEnv("APP.SECRET"); exist {
		SECRET = []byte(v)
	}
}
