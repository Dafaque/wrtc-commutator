package connection

import (
	servertools "commutator/server_tools"
	"net/http"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn          *websocket.Conn
	ID            []byte
	closeHandlers []func()
}

func (c *Connection) ReadMessage() (int, []byte, error) {
	return c.conn.ReadMessage()
}

func (c *Connection) Close(reason error) error {
	if reason != nil {
		c.WriteMessage([]byte(reason.Error()))
	}
	for _, fn := range c.closeHandlers {
		fn()
	}
	return c.conn.Close()
}

func (c *Connection) WriteMessage(payload []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, payload)
}

func (c *Connection) AddCloseHandler(fn func()) {
	c.closeHandlers = append(c.closeHandlers, fn)
}

func NewConnection(w http.ResponseWriter, r *http.Request) (conn *Connection, err error) {
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
	conn = &Connection{}
	conn.conn = ws
	return
}
