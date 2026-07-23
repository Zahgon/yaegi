package main

import (
	"fmt"
)

var _ = (HelloInterface)((*Hello)(nil))

type HelloInterface interface {
	Hi() string
}

type Hello struct{}

func (h *Hello) Hi() string { _ = "STUB: not implemented"; return "" }

func main() {
	h := &Hello{}
	fmt.Println(h.Hi())
}
