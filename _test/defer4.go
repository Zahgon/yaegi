package main

import "sync"

type T struct {
	mu   sync.RWMutex
	name string
}

func (t *T) get() string { _ = "STUB: not implemented"; return "" }

var d = T{name: "test"}

func main() {
	println(d.get())
}
