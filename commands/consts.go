package commands

import "github.com/gorilla/websocket"

const (
	METHOD_ONLINE rune = 43
	METHOD_OFFER  rune = 62
	METHOD_ANSWER rune = 60
)

type WSHandler func(*websocket.Conn, []byte) error

var exec map[rune]WSHandler = map[rune]WSHandler{
	METHOD_ONLINE: Online,
	METHOD_OFFER:  SendOffer,
	METHOD_ANSWER: SendAnswer,
}
