package main

import "fmt"

func f(a []int, b int) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a := []int{1, 2}
	r := f(a, 3)
	fmt.Println(r.([]int))
}
