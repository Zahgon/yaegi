package main

import "fmt"

func main() {
	a := 64
	a += 64
	fmt.Printf("a: %v %T", a, a)
	fmt.Println()

	b := 64
	b -= 64
	fmt.Printf("b: %v %T", b, b)
	fmt.Println()

	c := 64
	c *= 64
	fmt.Printf("c: %v %T", c, c)
	fmt.Println()

	d := 64
	d /= 64
	fmt.Printf("d: %v %T", d, d)
	fmt.Println()

	e := 64
	e %= 64
	fmt.Printf("e: %v %T", e, e)
	fmt.Println()
}
