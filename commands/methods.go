package commands

import (
	"commutator/connection"
	"commutator/errcodes"
	"commutator/logger"
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

	logger.Println(
		"SendOffer",
		"to:", string(to),
		"with:", string(p),
	)
	msg := model.NewSDP(ws.ID, p, MODE_OFFER, s)
	if !msg.Verify(to) {
		return errcodes.ERROR_CODE_BAD_SIGNATURE
	}

	// TODO: DRY
	b, errSerialize := Serialize(msg)
	if errSerialize != nil {
		logger.Println("err serialize:", errSerialize.Error())
		return errcodes.ERROR_CODE_UNKNOWN
	}
	var sdpData []byte = []byte{RESULT_SDP_MESSAGE}
	return Dial(string(to), append(sdpData, b...))
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

	logger.Println(
		"SendAnswer",
		"to:", string(to),
		"with:", string(p),
	)

	msg := model.NewSDP(ws.ID, p, MODE_ANSWER, s)
	if !msg.Verify(to) {
		return errcodes.ERROR_CODE_BAD_SIGNATURE
	}

	// TODO: DRY
	b, errSerialize := Serialize(msg)
	if errSerialize != nil {
		logger.Println("err serialize:", errSerialize.Error())
		return errcodes.ERROR_CODE_UNKNOWN
	}
	var sdpData []byte = []byte{RESULT_SDP_MESSAGE}

	return Dial(string(to), append(sdpData, b...))
}

func SendCandidates(ws *connection.Connection, args []byte) errcodes.ErrorCode {
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

	logger.Println(
		"SendCandidates",
		"to:", string(to),
		"with:", string(p),
	)

	msg := model.NewCandidates(ws.ID, p, s)
	if !msg.Verify(to) {
		return errcodes.ERROR_CODE_BAD_SIGNATURE
	}

	b, errSerialize := Serialize(msg)
	if errSerialize != nil {
		logger.Println("err serialize:", errSerialize.Error())
		return errcodes.ERROR_CODE_UNKNOWN
	}
	var sdpData []byte = []byte{RESULT_CANDIDATES_MESSAGE}

	return Dial(string(to), append(sdpData, b...))
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

	go func() {
		l, err := net.Listen(NETWORK, ":0")
		if err != nil {
			panic(err)
		}

		ws.AddCloseHandler(func() {
			if l != nil {
				l.Close()
			}
		})
		defer func() {
			if l != nil {
				l.Close()
			}
		}()
		{
			port := l.Addr().(*net.TCPAddr).Port
			var b []byte = []byte{RESULT_ONLINE}
			ws.ID = PortToHex(port)
			b = append(b, ws.ID...)
			errSendID := ws.WriteMessage(b)
			if errSendID != nil {
				logger.Println("cannot send ID:", errSendID.Error())
				ws.Error(errcodes.ERROR_CODE_CANT_WRITE_BACK)
				return
			}
		}

		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Println("cannot accept connection:", err.Error())
				break
			}
			{
				errSetDeadline := conn.SetDeadline(
					time.Now().Add(
						time.Duration(CONNECTION_TIMEOUT) * time.Second,
					),
				)
				if errSetDeadline != nil {
					logger.Println("cannot set deadline:", errSetDeadline.Error())
				}
			}

			data, errReadFromConn := io.ReadAll(conn)
			if errReadFromConn != nil {
				logger.Println("err read tcp message:", errReadFromConn.Error())
				continue
			}
			conn.Close()
			errWrite := ws.WriteMessage(data)
			if errWrite != nil {
				logger.Println("cannot write message:", errWrite.Error())
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

func Dial(hexPort string, data []byte) errcodes.ErrorCode {
	port, err := HexToPort(hexPort)
	if err != nil {
		logger.Println("cannot cast hexPort to int:", err.Error())
		return errcodes.ERROR_CODE_UNKNOWN
	}

	conn, errDial := net.Dial(NETWORK, fmt.Sprintf(":%d", port))
	if errDial != nil {
		logger.Println("err dial:", port, errDial.Error())
		return errcodes.ERROR_CODE_TARGET_UNACCESSABLE
	}

	_, errWrite := conn.Write(data)
	if errWrite != nil {
		logger.Println("err write:", port, errWrite.Error())
		return errcodes.ERROR_CODE_TARGET_UNACCESSABLE
	}
	if errClose := conn.Close(); errClose != nil {
		logger.Println("err write:", port, errClose.Error())
		return errcodes.ERROR_CODE_UNKNOWN
	}
	return errcodes.ERROR_CODE_NONE
}
