package main

import "fmt"

type T int

func (t T) Error() string { _ = "STUB: not implemented"; return "" }

func f(t T) error { _ = "STUB: not implemented"; return nil }

func main() {
	x := T(1)
	fmt.Println(f(x))
}
