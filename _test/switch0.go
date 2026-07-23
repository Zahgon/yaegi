package main

import "fmt"

func f(i int) bool { _ = "STUB: not implemented"; return false }

func main() {
	r0 := f(0)
	fmt.Printf("%T %v", r0, r0)
	fmt.Println()
	r1 := f(1)
	fmt.Printf("%T %v", r1, r1)
	fmt.Println()
}
