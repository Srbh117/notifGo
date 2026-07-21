package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
}

type Consumer interface {
	Start() error
}

type WSConsumer struct {
	ListenAddr string
	server     *Server
}

func NewWSConsumer(listenAddr string, server *Server) *WSConsumer {
	return &WSConsumer{ListenAddr: listenAddr, server: server}
}

func (ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(err.Error())
	} else {
		p := NewWSPeer(conn)
		ws.server.AddConn(p)
	}
}

func (ws *WSConsumer) Start() error {
	slog.Info("websocket consumer started", "port:", ws.ListenAddr)
	return http.ListenAndServe(ws.ListenAddr, ws)
}

func foo() {
	websocket.DefaultDialer.Dial("ws:/oo", nil)

}
