package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func client(uri string) { _ = "STUB: not implemented"; return }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client(server.URL)
}
