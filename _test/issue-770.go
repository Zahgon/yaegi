package main

import "reflect"

type I interface {
	Foo() int
}

type T struct {
	Name string
}

func (t T) Foo() int { _ = "STUB: not implemented"; return 0 }

func f(v reflect.Value) int { _ = "STUB: not implemented"; return 0 }

func main() {
	println("hello")
}
