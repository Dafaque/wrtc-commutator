package commands

import (
	"commutator/connection"
	"unicode/utf8"
)

func Exec(ws *connection.Connection, payload []byte) {
	a, s := utf8.DecodeRune(payload)
	if fn, exists := exec[a]; exists {
		if err := fn(ws, payload[s:]); err != nil {
			ws.WriteMessage([]byte(err.Error()))
			ws.Close()
		}
	}
}
