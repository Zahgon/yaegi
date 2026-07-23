package main

import (
	"net/http"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
})

type T1 struct {
	Name string
}

func (t *T1) Handler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func main() {
	t := &T1{"myName"}
	handler := t.Handler(myHandler)
	http.ListenAndServe(":8080", handler)
}
