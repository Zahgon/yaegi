package main

import (
	"fmt"
	"sync"
)

type Hello struct {
	mu sync.Mutex
}

func (h *Hello) Hi() string { _ = "STUB: not implemented"; return "" }

func main() {
	a := &Hello{}

	fmt.Println(a.Hi())
}
