package main

import "fmt"

func main() {
	a := 'r'
	a += 'g'
	fmt.Printf("a: %v %T", a, a)
	fmt.Println()

	b := 'r'
	b -= 'g'
	fmt.Printf("b: %v %T", b, b)
	fmt.Println()

	c := 'r'
	c *= 'g'
	fmt.Printf("c: %v %T", c, c)
	fmt.Println()

	d := 'r'
	d /= 'g'
	fmt.Printf("d: %v %T", d, d)
	fmt.Println()

	e := 'r'
	e %= 'g'
	fmt.Printf("e: %v %T", e, e)
	fmt.Println()

	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(b == d)
}
