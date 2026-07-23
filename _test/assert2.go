package main

import (
	"sync"
)

type stringBuilder interface {
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (int, error)
	Reset()
	Grow(n int)
	String() string
}

var builderPool = sync.Pool{New: func() interface{} {
	return newStringBuilder()
}}

func newStringBuilder() stringBuilder { _ = "STUB: not implemented"; return *new(stringBuilder) }

func main() {
	i := builderPool.Get()
	sb := i.(stringBuilder)
	_, _ = sb.WriteString("hello")

	println(sb.String())

	builderPool.Put(i)
}
