package handlers

import (
	"commutator/commands"
	servertools "commutator/server_tools"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
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
	ws.SetPingHandler(func(appData string) error {
		return ws.WriteControl(websocket.PongMessage, nil, time.Now().Add(1*time.Second))
	})
	ws.SetPongHandler(func(appData string) error {
		return ws.WriteControl(websocket.PingMessage, nil, time.Now().Add(1*time.Second))
	})

	for {
		_, payload, err := ws.ReadMessage()
		if err != nil {
			println("err in read loop", err.Error())
			break
		}
		if len(payload) == 0 {
			continue
		}
		commands.Exec(ws, payload)
	}
}
