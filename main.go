package main

import (
	_ "commutator/commands"
	"commutator/handlers"
	servertools "commutator/server_tools"
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", handlers.Entrypoint)

	var server *http.Server = &http.Server{
		// TODO: domain name, i guess
		Addr:      "0.0.0.0:8080",
		Handler:   mux,
		TLSConfig: &tls.Config{},
		// TODO: zeros to numbers
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       1 * time.Minute,
		// TODO: calculate this number
		MaxHeaderBytes: 20000,
		// TODO: wtf?
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
		ConnState:    servertools.OnConnectionStateChanged,
		// TODO: own logger
		ErrorLog:    &log.Logger{},
		BaseContext: servertools.MakeContextForListener,
		ConnContext: servertools.MakeContextForConnection,
	}
	panic(server.ListenAndServe())
}
