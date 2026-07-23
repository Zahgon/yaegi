package main

import (
	"net/http"
	"net/http/httptest"
)

type mw1 struct {
	next http.Handler
}

func (m *mw1) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	_ = "STUB: not implemented"
	return
}

type mw0 struct{}

func (m *mw0) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	m0 := &mw0{}
	m1 := &mw1{next: m0}

	mux := http.NewServeMux()
	mux.HandleFunc("/", m1.ServeHTTP)

	server := httptest.NewServer(mux)
	defer server.Close()

	client(server.URL)
}

func client(uri string) { _ = "STUB: not implemented"; return }
