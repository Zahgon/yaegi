package main

import "fmt"

type T struct {
	Name string
}

func (T) create() *T { _ = "STUB: not implemented"; return nil }

func main() {
	fmt.Println(T{}.create())
}
