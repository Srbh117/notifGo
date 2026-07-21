package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	listenAddr string
}

func NewHTTPProdcuer(listenAddr string) *HTTPProducer {
	return &HTTPProducer{
		listenAddr: listenAddr,
	}
}

func (p *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func (p *HTTPProducer) Start() error {
	slog.Info("Start function has started")
	return http.ListenAndServe(p.listenAddr, p)
}
