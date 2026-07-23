package main

import (
	"io"
)

type T []byte

func (t *T) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func foo(w io.Writer) { _ = "STUB: not implemented"; return }

func main() {
	x := T{}
	foo(&x)
}
