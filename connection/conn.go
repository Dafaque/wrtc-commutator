package connection

import (
	"commutator/errcodes"
	servertools "commutator/server_tools"
	"net/http"

	"github.com/gorilla/websocket"
)

const RESULT_ERROR byte = 33

type Connection struct {
	conn          *websocket.Conn
	ID            []byte
	Tag           []byte
	closeHandlers []func()
}

func (c *Connection) ReadMessage() (int, []byte, error) {
	return c.conn.ReadMessage()
}

func (c *Connection) Error(reason errcodes.ErrorCode) error {
	if reason != errcodes.ERROR_CODE_NONE {
		var msg []byte = []byte{RESULT_ERROR}
		c.WriteMessage(append(msg, byte(reason)))
	}
	if reason < 100 {
		return nil
	}
	for _, fn := range c.closeHandlers {
		fn()
	}
	return c.conn.Close()
}

func (c *Connection) WriteMessage(payload []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, payload)
}

func (c *Connection) AddCloseHandler(fn func()) {
	c.closeHandlers = append(c.closeHandlers, fn)
}

func NewConnection(w http.ResponseWriter, r *http.Request) (conn *Connection, err error) {
	ws, err := servertools.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		println("cannot upgrade:", err.Error())
		return
	}
	// TODO
	ws.SetCloseHandler(nil)
	ws.SetPingHandler(nil)
	ws.SetPongHandler(nil)
	conn = &Connection{}
	conn.conn = ws

	conn.AddCloseHandler(func() {
		servertools.ConnectionClosed()
	})
	return
}
