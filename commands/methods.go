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
	if _, exists := global.Users[global.UserIdentifyer(id)]; exists {
		return UserNameAleadyExists()
	}
	global.Users[global.UserIdentifyer(id)] = ws

	ws.SetCloseHandler(func(code int, text string) error {
		delete(global.Users, global.UserIdentifyer(id))
		return ws.Close()
	})
	return nil
}
