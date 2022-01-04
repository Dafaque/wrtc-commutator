package handlers

import (
	"commutator/commands"
	"commutator/connection"
	servertools "commutator/server_tools"
	"net/http"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	if v, _ := r.Context().Value(servertools.ConnectionsOverflowKey).(bool); v {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}
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
	// TODO: make shure cycle is breaks
}
