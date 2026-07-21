package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

type Config struct {
	ListenerAddr      string
	WSListenAddr      string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	*Config

	mu        sync.RWMutex
	topics    map[string]Storer
	consumers []Consumer
	peers     map[Peer]bool
	producers []Producer

	quitch chan struct{}
}

func ReturnServer(cfg Config) *Server {
	s := &Server{Config: &cfg,
		topics:    make(map[string]Storer),
		quitch:    make(chan struct{}),
		producers: []Producer{NewHTTPProdcuer(cfg.ListenerAddr)},
	}
	s.consumers = []Consumer{NewWSConsumer(cfg.WSListenAddr, s)}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
	w.Write([]byte(r.URL.Path))

}

func (s *Server) Start() {
	for _, consumer := range s.consumers {
		go func(c Consumer) {
			if err := consumer.Start(); err != nil {
				fmt.Println(err)
			}
		}(consumer)
	}

	for _, producer := range s.producers {
		go func(p Producer) {
			if err := producer.Start(); err != nil {
				fmt.Println(err)
			}
		}(producer)
	}
	<-s.quitch
	s.loop()
}

func (s *Server) CreateNewTopic(topicName string) bool {
	_, ok := s.topics[topicName]
	if ok {
		return false
	}
	s.topics[topicName] = ReturnNewMemoryStore()
	return true
}

func (s *Server) SearchByTopic(topicName string) (Storer, error) {
	val, ok := s.topics[topicName]
	if !ok {
		return &MemoryStorage{}, fmt.Errorf("Topic-name doesn't exist")
	} else {
		return val, nil
	}
}

func (s *Server) AddConn(c Peer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	slog.Info("added new peer", "peer", c)
	s.peers[c] = true
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
			return

		}
	}
}
