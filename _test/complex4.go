package main

import "fmt"

func f(a, b float64) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a := f(3, 2)
	fmt.Println(a.(complex128))
}
