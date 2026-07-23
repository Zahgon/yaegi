package main

import "fmt"

func f() interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a, ok := f().(map[int]int)
	fmt.Println(a, ok)
}
