package main

import "fmt"

func f(c complex128) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	c := complex(3, 2)
	a := f(c)
	fmt.Println(a.(float64))
}
