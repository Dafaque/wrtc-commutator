package commands

import (
	"commutator/connection"
	"commutator/errcodes"
	"commutator/model"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func SendOffer(ws *connection.Connection, args []byte) errcodes.ErrorCode {
	if len(ws.ID) == 0 {
		return errcodes.ERROR_CODE_NOT_ONLINE
	}
	to, e := parseArg(ARG_TO, &args)
	if e != errcodes.ERROR_CODE_NONE {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != errcodes.ERROR_CODE_NONE {
		return e
	}
	s, e := parseArg(ARG_SIGN, &args)
	if e != errcodes.ERROR_CODE_NONE {
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
		return errcodes.ERROR_CODE_BAD_SIGNATURE
	}
	return Dial(string(to), msg)
}

func SendAnswer(ws *connection.Connection, args []byte) errcodes.ErrorCode {
	if len(ws.ID) == 0 {
		return errcodes.ERROR_CODE_NOT_ONLINE
	}
	to, e := parseArg(ARG_TO, &args)
	if e != errcodes.ERROR_CODE_NONE {
		return e
	}
	p, e := parseArg(ARG_WITH, &args)
	if e != errcodes.ERROR_CODE_NONE {
		return e
	}

	s, e := parseArg(ARG_SIGN, &args)
	if e != errcodes.ERROR_CODE_NONE {
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
		return errcodes.ERROR_CODE_BAD_SIGNATURE
	}
	return Dial(string(to), msg)
}

func Online(ws *connection.Connection, args []byte) errcodes.ErrorCode {
	//TODO validate args
	if len(ws.ID) > 0 {
		return errcodes.ERROR_CODE_ALREADY_ONLINE
	}
	{
		tag, e := parseArg(ARG_WITH, &args)
		if e != errcodes.ERROR_CODE_NONE {
			return e
		}
		ws.Tag = tag
	}

	println("Online as", string(ws.Tag))

	go func() {
		l, err := net.Listen(NETWORK, ":0")
		if err != nil {
			panic(err)
		}

		ws.AddCloseHandler(func() {
			if e := l.Close(); e != nil {
				println("listener already closed", string(ws.Tag))
			}
			println("listener closed", string(ws.Tag))
		})
		defer l.Close()
		{
			port := l.Addr().(*net.TCPAddr).Port
			var b []byte = []byte{RESULT_ONLINE}
			ws.ID = PortToHex(port)
			b = append(b, ws.ID...)
			errSendID := ws.WriteMessage(b)
			if errSendID != nil {
				println("err send online ID:", errSendID)
				ws.Error(errcodes.ERROR_CODE_CANT_WRITE_BACK)
				return
			}
		}

		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("err accept conn", err)
				break
			}

			conn.SetDeadline(
				time.Now().Add(
					time.Duration(CONNECTION_TIMEOUT) * time.Second,
				),
			)

			data, errReadFromConn := io.ReadAll(conn)
			if errReadFromConn != nil {
				println("err read tcp message: ", errReadFromConn.Error())
			}
			conn.Close()
			errWrite := ws.WriteMessage(data)
			if errWrite != nil {
				ws.Error(errcodes.ERROR_CODE_CANT_WRITE_BACK)
				break
			}
		}
	}()
	return errcodes.ERROR_CODE_NONE
}

func PortToHex(p int) []byte {
	return []byte(fmt.Sprintf("%x", p))
}

func HexToPort(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func Dial(hexPort string, sdp *model.SDP) errcodes.ErrorCode {
	port, err := HexToPort(hexPort)
	if err != nil {
		// TODO
		return errcodes.ERROR_CODE_UNKNOWN
	}

	conn, errDial := net.Dial(NETWORK, fmt.Sprintf(":%d", port))
	if errDial != nil {
		return errcodes.ERROR_CODE_TARGET_UNACCESSABLE
	}
	b, errSerialize := Serialize(sdp)
	if errSerialize != nil {
		return errcodes.ERROR_CODE_UNKNOWN
	}
	var sdpData []byte = []byte{RESULT_SDP_MESSAGE}
	_, errWrite := conn.Write(append(sdpData, b...))
	if errWrite != nil {
		return errcodes.ERROR_CODE_TARGET_UNACCESSABLE
	}
	if errClose := conn.Close(); errClose != nil {
		// TODO
		return errcodes.ERROR_CODE_UNKNOWN
	}
	return errcodes.ERROR_CODE_NONE
}
