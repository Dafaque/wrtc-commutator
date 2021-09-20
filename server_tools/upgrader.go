package servertools

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func onWsUpgradeError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	w.WriteHeader(http.StatusResetContent)
	w.Write([]byte(reason.Error()))
}

func checkOrigin(r *http.Request) bool {
	// TODO
	return true
}

var Upgrader *websocket.Upgrader = &websocket.Upgrader{
	// TODO: normalize this nums
	HandshakeTimeout: 10 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	Error:            onWsUpgradeError,
	CheckOrigin:      checkOrigin,
}
