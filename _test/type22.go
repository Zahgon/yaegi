package main

type T1 T

func foo() T1 { _ = "STUB: not implemented"; return *new(T1) }

type T struct {
	Name string
}

func main() {
	println(foo().Name)
}
