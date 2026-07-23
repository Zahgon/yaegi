package main

import (
	"net/http"
	"net/http/httptest"
)

func mock(name string) http.HandlerFunc { _ = "STUB: not implemented"; return *new(http.HandlerFunc) }

func client(uri string) { _ = "STUB: not implemented"; return }

func main() {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	mux.Handle("/", mock("foo"))
	client(server.URL)
}
