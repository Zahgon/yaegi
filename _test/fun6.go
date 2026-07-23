package main

import (
	"fmt"
	"sync"
)

func NewPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }

type Pool struct {
	p *sync.Pool
}

var _pool = NewPool()

func main() {
	fmt.Println(_pool)
}
