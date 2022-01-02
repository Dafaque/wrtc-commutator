package handlers

import (
	"commutator/commands"
	"commutator/connection"
	"net/http"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	conn, err := connection.NewConnection(w, r)
	if err != nil {
		return
	}

	for {
		_, payload, errReadMessage := conn.ReadMessage()
		if errReadMessage != nil {
			conn.Close(errReadMessage)
			return
		}
		commands.Exec(conn, payload)
	}
}
