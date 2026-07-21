package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Memory
//Server (TCP,HTTP)

func def() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		mgs := WSMessage{
			Action: "subscribe",
			Topic:  "Saurabh",
		}
		b, err := json.Marshal(mgs)
		if err != nil {
			log.Fatal(err)
		}
		conn.WriteJSON(b)
		fmt.Println(conn)
	}
}

func main() {
	cfg := &Config{ListenerAddr: ":42069", WSListenAddr: ":4000", StoreProducerFunc: func() Storer {
		return NewMemoryStorage()
	}}
	s := ReturnServer(*cfg)
	go s.Start()
	// slog.Info("Does my program come till here?")
	go def()
	select {}
}
