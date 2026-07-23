package main

import (
	"fmt"
)

func foo() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func main() {
	a, b := foo()
	fmt.Println(a, b)
}
