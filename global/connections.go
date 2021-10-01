package global

import "github.com/gorilla/websocket"

type UserIdentifyer string

var Users map[UserIdentifyer]*websocket.Conn = make(map[UserIdentifyer]*websocket.Conn)
