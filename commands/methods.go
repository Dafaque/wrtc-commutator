package commands

import (
	"commutator/global"

	"github.com/gorilla/websocket"
)

func SendOffer(ws *websocket.Conn, args []byte) error {
	to, e := parseArg("to", &args)
	if e != nil {
		return e
	}
	p, e := parseArg("with", &args)
	if e != nil {
		return e
	}
	println("SendOffer", "to:", to, "with:", p)
	return nil
}

func SendAnswer(ws *websocket.Conn, args []byte) error {
	to, e := parseArg("to", &args)
	if e != nil {
		return e
	}
	p, e := parseArg("with", &args)
	if e != nil {
		return e
	}
	println("SendAnswer", "to:", to, "with:", p)
	return nil
}

func Online(ws *websocket.Conn, args []byte) error {
	//TODO validate args
	id, e := parseArg("with", &args)
	if e != nil {
		return e
	}
	global.Connections[id] = ws

	println("Online", "with:", id)
	return nil
}
