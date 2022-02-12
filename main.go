package main

import (
	_ "commutator/commands"
	"commutator/handlers"
	"commutator/logger"
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
		Addr:              "0.0.0.0:8080",
		Handler:           mux,
		TLSConfig:         &tls.Config{},
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		// TODO: calculate this number
		MaxHeaderBytes: 20000,
		// TODO: wtf?
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
		ConnState:    servertools.OnConnectionStateChanged,
		// TODO: logger
		ErrorLog:    &log.Logger{},
		BaseContext: servertools.MakeListenerContext,
		ConnContext: servertools.MakeConnectionContext,
	}
	logger.Println("Server started")
	panic(server.ListenAndServe())
}
