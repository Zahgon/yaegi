package main

import (
	"net/http"
	"net/http/httptest"
)

type T struct {
	http.ResponseWriter
}

type mw1 struct {
	next http.Handler
}

func (m *mw1) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	_ = "STUB: not implemented"
	return
}

func main() {
	m1 := &mw1{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", m1.ServeHTTP)

	server := httptest.NewServer(mux)
	defer server.Close()

	client(server.URL)
}

func client(uri string) { _ = "STUB: not implemented"; return }
