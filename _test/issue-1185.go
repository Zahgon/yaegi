package main

import "io"

type B []byte

func (b B) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	b := B{}
	a := make([]io.Writer, 0)
	a = append(a, b)
	println(len(a))
}
