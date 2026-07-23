package main

import "fmt"

const (
	Foo = iota
	Bar
	Baz
)

const (
	Asm = iota
	C
	Java
	Go
)

func main() {
	fmt.Println(Foo, Bar, Baz)
	fmt.Println(Asm, C, Java, Go)
}
