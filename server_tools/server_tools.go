package servertools

import (
	"context"
	"net"
	"net/http"
)

func OnConnectionStateChanged(c net.Conn, s http.ConnState) {
	// TODO
	println(c.RemoteAddr().String(), " changed conn state to", s)
}

func MakeContextForListener(l net.Listener) context.Context {
	return context.Background()
}

func MakeContextForConnection(ctx context.Context, c net.Conn) context.Context {
	// TODO
	return ctx
}
