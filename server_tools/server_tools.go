package servertools

import (
	"context"
	"net"
	"net/http"
	"sync/atomic"
)

var connectionsNumber uint32 = 0

type typeConnectionsOverflowKey string

var ConnectionsOverflowKey typeConnectionsOverflowKey = "overflow"

func OnConnectionStateChanged(c net.Conn, s http.ConnState) {
	if s == http.StateActive {
		atomic.AddUint32(&connectionsNumber, 1)
	}
	if s == http.StateClosed {
		ConnectionClosed()
	}
}

func MakeListenerContext(l net.Listener) context.Context {
	return context.Background()
}

func MakeConnectionContext(ctx context.Context, c net.Conn) context.Context {
	ctx = context.WithValue(ctx, ConnectionsOverflowKey, connectionsNumber > uint32(MAX_CONNECTIONS))
	return ctx
}

func ConnectionClosed() {
	if connectionsNumber > 0 {
		atomic.AddUint32(&connectionsNumber, ^uint32(0))
	}
}
