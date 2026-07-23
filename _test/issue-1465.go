package main

import (
	"fmt"
)

func SomeFunc[T int | string](defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func main() {
	fmt.Println(SomeFunc("test"))
	fmt.Println(SomeFunc(1234))
}
