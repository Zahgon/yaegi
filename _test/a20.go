package main

import "fmt"

type IntArray []int

func (h *IntArray) Add(x int) { _ = "STUB: not implemented"; return }

func main() {
	a := IntArray{}
	a.Add(4)

	fmt.Println(a)
}
