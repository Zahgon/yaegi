package main

import "unsafe"

type T struct {
	X uint64
	Y uint64
}

func f(off uintptr) { _ = "STUB: not implemented"; return }

func main() {
	f(unsafe.Offsetof(T{}.Y))
}
