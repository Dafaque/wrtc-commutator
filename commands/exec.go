package commands

import (
	"commutator/connection"
	"commutator/errcodes"
	"unicode/utf8"
)

func Exec(ws *connection.Connection, payload []byte) {
	a, s := utf8.DecodeRune(payload)
	if fn, exists := exec[a]; exists {
		if err := fn(ws, payload[s:]); err != errcodes.ERROR_CODE_NONE {
			ws.Error(err)
		}
	}
}
