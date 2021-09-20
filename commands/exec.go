package commands

import (
	"unicode/utf8"

	"github.com/gorilla/websocket"
)

func Exec(ws *websocket.Conn, payload []byte) {
	a, s := utf8.DecodeRune(payload)
	if fn, exists := exec[a]; exists {
		fn(ws, payload[s:])
	}
}
