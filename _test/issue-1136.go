package main

import (
	"fmt"
	"io"
)

type T struct {
	r io.Reader
}

func (t *T) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	x := io.LimitedReader{}
	y := io.Reader(&x)
	y = &T{y}
	fmt.Println(y.Read([]byte("")))
}
