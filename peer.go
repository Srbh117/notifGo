package main

import (
	"github.com/gorilla/websocket"
)

type Peer interface {
	Send([]byte) error
}

type WSPeer struct {
	conn *websocket.Conn
}

func NewWSPeer(conn *websocket.Conn) *WSPeer {
	return &WSPeer{
		conn: conn,
	}
}
func (p *WSPeer) Send(b []byte) (err error) {
	return p.conn.WriteMessage(websocket.BinaryMessage, b)
}
