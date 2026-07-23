package main

import (
	"fmt"
)

func SomeFunc(defaultValue interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	fmt.Println(SomeFunc(1234))
	fmt.Println(SomeFunc("test"))
}
