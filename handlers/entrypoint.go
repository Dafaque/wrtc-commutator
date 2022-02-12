package handlers

import (
	"commutator/commands"
	"commutator/connection"
	"commutator/errcodes"
	"commutator/logger"
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
		logger.Println("cannot create connection:", err.Error())
		return
	}

	for {
		_, payload, errReadMessage := conn.ReadMessage()
		if errReadMessage != nil {
			logger.Println("cannot read message:", errReadMessage.Error())
			conn.Error(errcodes.ERROR_CODE_MUMBLING)
			return
		}
		commands.Exec(conn, payload)
	}
	// TODO: make shure cycle is breaks
}
