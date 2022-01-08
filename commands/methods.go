package commands

import (
	"commutator/connection"
	"commutator/model"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func SendOffer(ws *connection.Connection, args []byte) error {
	if len(ws.ID) == 0 {
		return errors.New("not online")
	}
	to, e := parseArg(ARG_TO, &args)
	if e != nil {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != nil {
		return e
	}
	s, e := parseArg(ARG_SIGN, &args)
	if e != nil {
		return e
	}

	println(
		"SendOffer",
		"to:", string(to),
		"with:", string(p),
		"signature:", string(s),
	)

	msg := model.NewSDP(ws.ID, p, MODE_OFFER, s)
	if !msg.Verify(to) {
		return errors.New("bad signature")
	}
	return Dial(string(to), msg)
}

func SendAnswer(ws *connection.Connection, args []byte) error {
	if len(ws.ID) == 0 {
		return errors.New("not online")
	}
	to, e := parseArg(ARG_TO, &args)
	if e != nil {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != nil {
		return e
	}

	s, e := parseArg(ARG_SIGN, &args)
	if e != nil {
		return e
	}

	println(
		"SendAnswer",
		"to:", string(to),
		"with:", string(p),
		"signature:", string(s),
	)

	msg := model.NewSDP(ws.ID, p, MODE_ANSWER, s)
	if !msg.Verify(to) {
		return errors.New("bad signature")
	}
	return Dial(string(to), msg)
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
		l, err := net.Listen(NETWORK, ":0")
		if err != nil {
			panic(err)
		}

		ws.AddCloseHandler(func() {
			if e := l.Close(); e != nil {
				println("listener already closed", ws.ID)
			}
			println("listener closed", ws.ID)
		})
		defer l.Close()
		{
			port := l.Addr().(*net.TCPAddr).Port
			var b []byte = []byte{RESULT_ONLINE}
			b = append(b, PortToHex(port)...)
			errSendID := ws.WriteMessage(b)
			if errSendID != nil {
				println("err send online ID:", errSendID)
				return
			}
		}

		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("err accept conn", err)
				break
			}

			conn.SetDeadline(time.Now().Add(time.Duration(CONNECTION_TIMEOUT) * time.Second))

			data, errReadFromConn := io.ReadAll(conn)
			if errReadFromConn != nil {
				println("err read tcp message: ", errReadFromConn.Error())
			}
			conn.Close()
			errWrite := ws.WriteMessage(data)
			if errWrite != nil {
				ws.Close(errWrite)
				break
			}
		}
	}()
	return nil
}

func PortToHex(p int) []byte {
	return []byte(fmt.Sprintf("%x", p))
}

func HexToPort(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func Dial(hexPort string, sdp *model.SDP) error {
	port, err := HexToPort(hexPort)
	if err != nil {
		return err
	}

	conn, errDial := net.Dial(NETWORK, fmt.Sprintf(":%d", port))
	if errDial != nil {
		return errDial
	}
	b, errSerialize := Serialize(sdp)
	if errSerialize != nil {
		return errSerialize
	}
	conn.Write(b)
	return conn.Close()
}
