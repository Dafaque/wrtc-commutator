package commands

import (
	"commutator/connection"
	"commutator/errcodes"
	"os"
	"strconv"
)

const (
	METHOD_ONLINE     rune = 43
	METHOD_OFFER      rune = 62
	METHOD_ANSWER     rune = 60
	METHOD_CANDIDATES rune = 64

	RESULT_ERROR       byte = 33
	RESULT_ONLINE      byte = 64
	RESULT_SDP_MESSAGE byte = 42

	MODE_OFFER  byte = 0
	MODE_ANSWER byte = 1
)

var (
	ARG_TO   []byte = []byte{116, 111}
	ARG_WITH []byte = []byte{119, 105, 116, 104}
	ARG_SIGN []byte = []byte{115, 105, 103, 110}
)

var (
	NETWORK            string = "tcp6"
	CONNECTION_TIMEOUT int    = 5
)

func init() {
	if v, exist := os.LookupEnv("APP.NETWORK"); exist {
		NETWORK = v
	}
	if v, exist := os.LookupEnv("APP.CONNECTION_TIMEOUT"); exist {
		i, err := strconv.Atoi(v)
		if err != nil {
			println("err init vars: ", err.Error())
			return
		}
		CONNECTION_TIMEOUT = i
	}
	println("NETWORK:", NETWORK)
	println("CONNECTION_TIMEOUT:", CONNECTION_TIMEOUT)

}

type WSHandler func(*connection.Connection, []byte) errcodes.ErrorCode

var exec map[rune]WSHandler = map[rune]WSHandler{
	METHOD_ONLINE:     Online,
	METHOD_OFFER:      SendOffer,
	METHOD_ANSWER:     SendAnswer,
	METHOD_CANDIDATES: SendCandidates,
}
