package main

import "fmt"

func f(a, b []int) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a := []int{10, 20, 30}
	b := [4]int{}
	c := b[:]
	r := f(c, a)
	fmt.Println(r.(int))
}
