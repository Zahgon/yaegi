package main

import (
	"net/http"
)

var version string = "1.0"

type Middleware struct {
	Name string
}

func (m *Middleware) Handler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func main() {
	m := &Middleware{"Test"}
	http.HandleFunc("/", m.Handler)
	http.ListenAndServe(":8080", nil)
}
