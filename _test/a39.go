package main

import "fmt"

func main() {
	a := [...]byte{}
	b := a
	fmt.Printf("%T %T\n", a, b)
}
