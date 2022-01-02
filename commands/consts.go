package commands

import (
	"commutator/connection"
)

const (
	METHOD_ONLINE rune = 43
	METHOD_OFFER  rune = 62
	METHOD_ANSWER rune = 60
)

var (
	ARG_TO   []byte = []byte{116, 111}
	ARG_WITH []byte = []byte{119, 105, 116, 104}

	VAL_STAR []byte = []byte{42}
)

type WSHandler func(*connection.Connection, []byte) error

var exec map[rune]WSHandler = map[rune]WSHandler{
	METHOD_ONLINE: Online,
	METHOD_OFFER:  SendOffer,
	METHOD_ANSWER: SendAnswer,
}
