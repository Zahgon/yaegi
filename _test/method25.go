package main

import (
	"fmt"
	"sync"
)

func (p Pool) Get() *Buffer { _ = "STUB: not implemented"; return nil }

func NewPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }

type Buffer struct {
	bs   []byte
	pool Pool
}

type Pool struct {
	p *sync.Pool
}

var (
	_pool = NewPool()
	Get   = _pool.Get
)

func main() {
	fmt.Println(_pool)
	fmt.Println(Get())
}
