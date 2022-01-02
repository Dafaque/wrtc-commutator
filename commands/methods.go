package commands

import (
	"bytes"
	"commutator/connection"
	"commutator/messages"
	"errors"
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
	if bytes.EqualFold(to, VAL_STAR) {
		return errors.New("invalid target")
	}

	println("SendOffer", "to:", string(to), "with:", string(p))

	messages.NewPublsher().Broadcast(
		messages.NewMessage(
			to,
			ws.ID,
			p,
			[]byte{},
		),
	)

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
	if bytes.EqualFold(to, VAL_STAR) {
		return errors.New("invalid target")
	}
	println("SendAnswer", "to:", string(to), "with:", string(p))

	messages.NewPublsher().Broadcast(
		messages.NewMessage(
			to,
			ws.ID,
			p,
			[]byte{},
		),
	)
	return nil
}

func Online(ws *connection.Connection, args []byte) error {
	//TODO validate args
	if len(ws.ID) > 0 {
		return errors.New("already online")
	}
	{
		id, e := parseArg(ARG_WITH, &args)
		if e != nil {
			return e
		}
		ws.ID = id
	}

	println("Online as", string(ws.ID))

	go func() {
		pub := messages.NewPublsher()
		con := messages.NewConsumer()
		pub.Consume(con)
		defer pub.Unconsume(con)
		ws.AddCloseHandler(func() {
			pub.Unconsume(con)
		})
		for {
			msg, err := con.Read()
			if err != nil {
				ws.Close(err)
				println("err read:", err.Error())
				break
			}
			if bytes.EqualFold(msg.To, ws.ID) || bytes.EqualFold(msg.To, VAL_STAR) {
				b, errMarshal := Serialize(msg)
				if err != nil {
					ws.Close(errMarshal)
					break
				}

				if errWriteMessage := ws.WriteMessage(b); errWriteMessage != nil {
					ws.Close(errWriteMessage)
					println("err", errWriteMessage)
					break
				}
				con.Confirm()
			}
		}
	}()

	return nil
}
