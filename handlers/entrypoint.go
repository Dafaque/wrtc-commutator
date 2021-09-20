package handlers

import (
	servertools "commutator/server_tools"
	"net/http"

	"github.com/gorilla/websocket"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// TODO: figure out ws protocol method
	case http.MethodGet:
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var ws *websocket.Conn
	// TODO response header

	ws, err := servertools.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO
	ws.SetCloseHandler(nil)
	ws.SetPingHandler(nil)
	ws.SetPongHandler(nil)

	for {
		// _, payload, err := ws.ReadMessage()
	}
}
