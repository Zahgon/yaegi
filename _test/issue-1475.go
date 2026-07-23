package main

type T uint16

func f() T { _ = "STUB: not implemented"; return *new(T) }

func main() {
	println(f())
}
