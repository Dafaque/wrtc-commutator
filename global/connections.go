package global

import "github.com/gorilla/websocket"

var Connections map[string]*websocket.Conn = make(map[string]*websocket.Conn)
