package commands

import (
	"bytes"
	"commutator/connection"
	"commutator/messages"

	"github.com/vmihailenco/msgpack/v5"
)

func SendOffer(ws *connection.Connection, args []byte) error {
	to, e := parseArg(ARG_TO, &args)
	if e != nil {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != nil {
		return e
	}
	println("SendOffer", "to:", string(to), "with:", string(p))

	messages.NewPublsher().Broadcast(messages.NewMessage(to, []byte(ws.ID), p, []byte{}))

	return nil
}

func SendAnswer(ws *connection.Connection, args []byte) error {
	to, e := parseArg(ARG_TO, &args)
	if e != nil {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != nil {
		return e
	}
	println("SendAnswer", "to:", string(to), "with:", string(p))
	return nil
}

func Online(ws *connection.Connection, args []byte) error {
	//TODO validate args
	id, e := parseArg(ARG_WITH, &args)
	if e != nil {
		return e
	}
	ws.ID = id
	println("Online", "with:", string(id))

	pub := messages.NewPublsher()
	con := messages.NewConsumer()
	pub.Consume(con)
	defer pub.Unconsume(con)
	for {
		println("cycle")
		msg, err := con.Read()
		if err != nil {
			ws.Close()
			break
		}
		if bytes.EqualFold(msg.To, ws.ID) || bytes.EqualFold(msg.To, VAL_STAR) {
			b, err := msgpack.Marshal(msg)
			if err != nil {
				ws.WriteMessage([]byte(err.Error()))
			}
			if errWriteMessage := ws.WriteMessage(b); errWriteMessage != nil {
				ws.Close()
				println("err", errWriteMessage)
			}
		}
	}
	return nil
}
