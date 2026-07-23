package main

import "fmt"

func main() {
	a := [...]string{"hello", "world"}
	fmt.Printf("%v %T\n", a, a)
}
